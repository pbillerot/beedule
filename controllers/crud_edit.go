package controllers

import (
	"fmt"
	"strings"

	"github.com/pbillerot/beedule/app"
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

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Forms[formid].Elements, id)

	records, err := models.CrudRead(tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	if len(records) == 0 {
		flash.Error("Article non trouvé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	// Calcul des éléments (valeur par défaut comprise)
	elements = computeElements(c.Controller, true, elements, records[0])

	// c.SetSession("from", c.Ctx.Request.Referer())
	c.SetSession(fmt.Sprintf("anch_%s_%s", tableid, viewid), fmt.Sprintf("anch_%s", strings.ReplaceAll(id, ".", "_")))
	setContext(c.Controller)
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

	flash := beego.NewFlash()

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

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Forms[formid].Elements, id)

	// lecture du record
	records, err := models.CrudRead(tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	if len(records) == 0 {
		flash.Error("Article non trouvé: ", id)
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
		setContext(c.Controller)
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
	// PostSQL
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
			}
		}
	}

	flash.Notice("Mise à jour effectuée avec succès")
	flash.Store(&c.Controller)

	ReturnFrom(c.Controller)
	return
}
