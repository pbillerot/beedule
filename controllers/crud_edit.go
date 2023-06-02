package controllers

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
)

// CrudEditController as
type CrudEditController struct {
	loggedRouter
}

// Get CrudEditController
func (c *CrudEditController) Get() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	id := c.Ctx.Input.Param(":id")

	// beego.Debug("from", c.Ctx.Input.Cookie("from"))

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		logs.Error("App not found", c.GetSession("Username").(string), appid)
		backward(c.Controller)
		return
	}
	if val, ok := dico.Ctx.Applications[appid].Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				logs.Error("Form not found", c.GetSession("Username").(string), formid)
				backward(c.Controller)
				return
			}
		} else {
			logs.Error("View not found", c.GetSession("Username").(string), viewid)
			backward(c.Controller)
			return
		}
	} else {
		logs.Error("table not found", c.GetSession("Username").(string), tableid)
		backward(c.Controller)
		return
	}

	// Contrôle d'accès
	table := dico.Ctx.Applications[appid].Tables[tableid]
	view := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid]
	form := dico.Ctx.Applications[appid].Tables[tableid].Forms[formid]
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, appid, id) {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}
	setContext(c.Controller, appid, tableid)

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, appid, tableid, dico.Ctx.Applications[appid].Tables[tableid].Forms[formid].Elements, id)

	// Filtrage si élément owner
	filter := ""
	for key, element := range elements {
		// Un seule élément owner par enregistrement
		if element.Group == "owner" && !IsAdmin(c.Controller) {
			filter = key + " = '" + c.GetSession("Username").(string) + "'"
			break
		}
	}

	records, err := models.CrudRead(filter, appid, tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}
	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	// Calcul des éléments (valeur par défaut comprise)
	elements = computeElements(c.Controller, true, elements, records[0])

	// Positionnement du navigateur sur la page qui va s'ouvrir
	forward(c.Controller, fmt.Sprintf("/bee/edit/%s/%s/%s/%s/%s", appid, tableid, viewid, formid, id))

	// Titre du formulaire
	form.Title = macro(c.Controller, appid, form.Title, records[0])

	c.SetSession(fmt.Sprintf("anch_%s_%s", tableid, viewid), fmt.Sprintf("anch_%s", strings.ReplaceAll(id, ".", "_")))
	c.Data["AppId"] = appid
	c.Data["Application"] = dico.Ctx.Applications[appid]
	c.Data["ColDisplay"] = records[0][table.Setting.ColDisplay].(string)
	c.Data["Id"] = id
	c.Data["TableId"] = tableid
	c.Data["ViewId"] = viewid
	c.Data["FormId"] = formid
	c.Data["Table"] = &table
	c.Data["View"] = &view
	c.Data["Form"] = &form
	c.Data["Elements"] = elements
	c.Data["Records"] = records
	c.Data["Cols"] = cols

	c.TplName = "crud_edit.html"
}

// Post CrudEditController
func (c *CrudEditController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	id := c.Ctx.Input.Param(":id")

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl tableid et viewid
	if val, ok := dico.Ctx.Applications[appid].Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				flash.Error("Formulaire non trouvé : %v", formid)
				flash.Store(&c.Controller)
				backward(c.Controller)
				return
			}
		} else {
			flash.Error("Vue non trouvée : %v", viewid)
			flash.Store(&c.Controller)
			backward(c.Controller)
			return
		}
	} else {
		flash.Error("Application non trouvée : %v", appid)
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}
	table := dico.Ctx.Applications[appid].Tables[tableid]
	view := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid]
	form := dico.Ctx.Applications[appid].Tables[tableid].Forms[formid]

	setContext(c.Controller, appid, tableid)
	var withPlugin bool

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, appid, tableid, dico.Ctx.Applications[appid].Tables[tableid].Forms[formid].Elements, id)

	// Filtrage si élément owner
	filter := ""
	for key, element := range elements {
		// Un seule élément owner par enregistrement
		if element.Group == "owner" && !IsAdmin(c.Controller) {
			filter = key + " = '" + c.GetSession("Username").(string) + "'"
			break
		}
	}

	// lecture du record
	records, err := models.CrudRead(filter, appid, tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé: %v", id)
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	// Lecture, contrôle des champs saisis
	// et remplissage de SQLOut pour l'enregistrement
	berr := false
	for key, element := range elements {
		err = checkElement(&c.Controller, key, &element, records[0])
		if err != nil {
			berr = true
			element.Error = "error"
			flash.Error(err.Error())
		}
		elements[key] = element
	}
	// CheckSqlite
	for _, CheckSqlite := range form.CheckSqlite {
		sql := macroSQL(c.Controller, appid, CheckSqlite, records[0])
		if sql != "" {
			berr = true
			flash.Error(sql)
		}
	}
	// Calcul des éléments (valeur par défaut comprise)
	elements = computeElements(c.Controller, true, elements, records[0])

	if berr { // ERREUR: on va reproposer le formulaire pour rectification
		// Titre du formulaire
		form.Title = macro(c.Controller, appid, form.Title, records[0])

		flash.Store(&c.Controller)
		c.Data["AppId"] = appid
		c.Data["Application"] = dico.Ctx.Applications[appid]
		c.Data["ColDisplay"] = records[0][table.Setting.ColDisplay].(string)
		c.Data["Id"] = id
		c.Data["App"] = dico.Ctx
		c.Data["TableId"] = tableid
		c.Data["ViewId"] = viewid
		c.Data["FormId"] = formid
		c.Data["Table"] = &table
		c.Data["View"] = &view
		c.Data["Form"] = &form
		c.Data["Elements"] = elements
		c.Data["Records"] = records
		c.Data["Cols"] = cols

		c.TplName = "crud_edit.html"
		return
	}

	// C'est OK, les données sont correctes et placées dans SQLout
	err = models.CrudUpdate(appid, tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Data["error"] = "error"
		backward(c.Controller)
		return
	}
	// Remplissage d'un record avec les elements.SQLout
	record := orm.Params{}
	for key, element := range elements {
		record[key] = element.SQLout
	}

	// Traitements particuliers pour les images
	// Enregistrement de l'image modifiée
	for _, element := range elements {
		if element.Type == "image" {
			b64data := element.SQLout[strings.IndexByte(element.SQLout, ',')+1:]
			unbased, err := base64.StdEncoding.DecodeString(b64data)
			// img, _, err := image.Decode(bytes.NewReader([]byte(element.SQLout)))
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
				berr = true
				break
			}

			outputFile, err := os.Create(element.Params.Src)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
				berr = true
				break
			}
			defer outputFile.Close()

			outputFile.Write(unbased)

			// r := bytes.NewReader(unbased)
			// ext := filepath.Ext(element.params.Src)
			// if ext == ".png" {
			// 	im, err := png.Decode(r)
			// 	if err != nil {
			// 		flash.Error(err.Error())
			// 		flash.Store(&c.Controller)
			// 		berr = true
			// 		break
			// 	}
			// 	err = png.Encode(outputFile, im)
			// 	if err != nil {
			// 		flash.Error(err.Error())
			// 		flash.Store(&c.Controller)
			// 		berr = true
			// 	}
			// }
			// if ext == ".jpg" {
			// 	var opts jpeg.Options
			// 	opts.Quality = 1
			// 	err = jpeg.Encode(outputFile, img, &opts)
			// 	if err != nil {
			// 		flash.Error(err.Error())
			// 		flash.Store(&c.Controller)
			// 		berr = true
			// 	}
			// }

		}
	}

	// PostSQL du formulaire
	for _, postsql := range form.PostSQL {
		// Remplissage d'un record avec les elements.SQLout
		record := orm.Params{}
		for key, element := range elements {
			record[key] = element.SQLout
		}
		sql := macro(c.Controller, appid, postsql, record)
		if sql != "" {
			err = models.CrudExec(sql, table.Setting.AliasDB)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
				berr = true
			}
		}
	}

	if !berr {
		flash.Notice("Mise à jour effectuée avec succès")
		flash.Store(&c.Controller)
	}

	if withPlugin {
		c.Ctx.Redirect(302, "/bee/list/"+appid+"/"+tableid+"/"+viewid)
	} else {
		backward(c.Controller)
	}
}
