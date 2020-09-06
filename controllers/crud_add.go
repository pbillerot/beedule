package controllers

import (
	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// CrudAddController as
type CrudAddController struct {
	loggedRouter
}

// Get CrudAddController
func (c *CrudAddController) Get() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	var id = ""

	// Ctrl tableid et viewid
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				c.Ctx.Redirect(302, "/crud")
				return
			}
		} else {
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		c.Ctx.Redirect(302, "/crud")
		return
	}

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Forms[formid].Elements, id)

	// Création d'un record fictif vide ""
	record := orm.Params{}
	for _, colname := range cols {
		record[colname] = ""
	}
	var records []orm.Params
	records = append(records, record)

	// Remplissage records avec valeur par défaut
	for ir, record := range records {
		for key, val := range record {
			if val == "" {
				record[key] = elements[key].Value
				records[ir] = record
			}
		}
	}

	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	form := app.Tables[tableid].Forms[formid]

	setContext(c.Controller)
	c.Data["AppId"] = appid
	c.Data["Application"] = app.Applications[appid]
	c.Data["ColDisplay"] = records[0][table.ColDisplay]
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

// Post CrudAddController
func (c *CrudAddController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	var id = ""

	// Ctrl tableid et viewid
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				c.Ctx.Redirect(302, "/crud")
				return
			}
		} else {
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		c.Ctx.Redirect(302, "/crud")
		return
	}

	flash := beego.NewFlash()

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Forms[formid].Elements, id)

	// Création d'un record fictif vide ""
	record := orm.Params{}
	for _, colname := range cols {
		record[colname] = ""
	}
	var records []orm.Params
	records = append(records, record)

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
		// On écrase les données lues dans la table par la saisie
		// Ceci est nécessaire pour représenter les valeurs saisies suite à uen erreur
		records[0][key] = element.SQLout
	}
	if berr { // ERREUR: on va reproposer le formulaire pour rectification
		flash.Store(&c.Controller)
		table := app.Tables[tableid]
		view := app.Tables[tableid].Views[viewid]
		form := app.Tables[tableid].Forms[formid]

		setContext(c.Controller)
		c.Data["AppId"] = appid
		c.Data["Application"] = app.Applications[appid]
		c.Data["ColDisplay"] = records[0][table.ColDisplay]
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
		return
	}
	// C'est OK, les données sont correctes et placées dans SQLout
	err = models.CrudInsert(tableid, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Data["error"] = "error"
		c.Ctx.Redirect(302, c.Ctx.Request.RequestURI)
		return
	}

	flash.Notice("Création effectuée avec succès")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, "/crud/list/"+appid+"/"+tableid+"/"+viewid)
}