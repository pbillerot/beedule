package controllers

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/batch"
	"github.com/pbillerot/beedule/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

	beego.Debug("from", c.Ctx.Input.Cookie("from"))

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := app.Applications[appid]; !ok {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				beego.Error("Form not found", c.GetSession("Username").(string), formid)
				ReturnFrom(c.Controller)
				return
			}
		} else {
			beego.Error("View not found", c.GetSession("Username").(string), viewid)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		beego.Error("table not found", c.GetSession("Username").(string), tableid)
		ReturnFrom(c.Controller)
		return
	}

	// Contrôle d'accès
	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	form := app.Tables[tableid].Forms[formid]
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, id) {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	setContext(c.Controller, tableid)

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Forms[formid].Elements, id)

	// Filtrage si élément owner
	filter := ""
	for key, element := range elements {
		// Un seule élément owner par enregistrement
		if element.Group == "owner" && !IsAdmin(c.Controller) {
			filter = key + " = '" + c.GetSession("Username").(string) + "'"
			break
		}
	}

	records, err := models.CrudRead(filter, tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	// Calcul des éléments (valeur par défaut comprise)
	elements = computeElements(c.Controller, true, elements, records[0])

	c.SetSession(fmt.Sprintf("anch_%s_%s", tableid, viewid), fmt.Sprintf("anch_%s", strings.ReplaceAll(id, ".", "_")))
	c.Data["AppId"] = appid
	c.Data["Application"] = app.Applications[appid]
	c.Data["ColDisplay"] = records[0][table.ColDisplay].(string)
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
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				flash.Error("Formulaire non trouvé :", formid)
				flash.Store(&c.Controller)
				ReturnFrom(c.Controller)
				return
			}
		} else {
			flash.Error("Vue non trouvée :", viewid)
			flash.Store(&c.Controller)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		flash.Error("Application non trouvée :", appid)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	form := app.Tables[tableid].Forms[formid]

	setContext(c.Controller, tableid)
	var withPlugin bool

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Forms[formid].Elements, id)

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
	records, err := models.CrudRead(filter, tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé: ", id)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
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
	// CheckSQL
	for _, checksql := range form.CheckSQL {
		sql := macroSQL(c.Controller, checksql, records[0])
		if sql != "" {
			berr = true
			flash.Error(sql)
		}
	}
	// Calcul des éléments (valeur par défaut comprise)
	elements = computeElements(c.Controller, true, elements, records[0])

	if berr { // ERREUR: on va reproposer le formulaire pour rectification
		flash.Store(&c.Controller)
		c.Data["AppId"] = appid
		c.Data["Application"] = app.Applications[appid]
		c.Data["ColDisplay"] = records[0][table.ColDisplay].(string)
		c.Data["Id"] = id
		c.Data["App"] = app.Portail
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
	err = models.CrudUpdate(tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Data["error"] = "error"
		ReturnFrom(c.Controller)
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

			outputFile, err := os.Create(element.Params.Path)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
				berr = true
				break
			}
			defer outputFile.Close()

			outputFile.Write(unbased)

			// r := bytes.NewReader(unbased)
			// ext := filepath.Ext(element.Params.Path)
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

	// PostActions des éléments
	for _, element := range elements {
		if len(element.PostAction) > 0 {
			// Traitement des actions
			for _, action := range element.PostAction {
				if berr {
					break
				}
				// Traitement SQL
				for _, actionsql := range action.SQL {
					sql := macro(c.Controller, actionsql, record)
					if sql != "" {
						err = models.CrudExec(sql, table.AliasDB)
						if err != nil {
							flash.Error(err.Error())
							flash.Store(&c.Controller)
							berr = true
						}
					}
				}
				// Traitement Plugin
				if action.Plugin != "" {
					withPlugin = true
					plugin := macro(c.Controller, action.Plugin, record)
					_, err := batch.RunPlugin(macro(c.Controller, plugin, record))
					if err != nil {
						flash.Error(err.Error())
						flash.Store(&c.Controller)
						berr = true
					}
				}
			}
		}
	}

	// PostSQL du formulaire
	for _, postsql := range form.PostSQL {
		// Remplissage d'un record avec les elements.SQLout
		record := orm.Params{}
		for key, element := range elements {
			record[key] = element.SQLout
		}
		sql := macro(c.Controller, postsql, record)
		if sql != "" {
			err = models.CrudExec(sql, table.AliasDB)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
				berr = true
			}
		}
	}

	if berr == false {
		flash.Notice("Mise à jour effectuée avec succès")
		flash.Store(&c.Controller)
	}

	if withPlugin {
		c.Ctx.Redirect(302, "/bee/list/"+appid+"/"+tableid+"/"+viewid)
	} else {
		ReturnFrom(c.Controller)
	}

	return
}
