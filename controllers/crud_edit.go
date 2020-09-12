package controllers

import (
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

	flash := beego.ReadFromRequest(&c.Controller)

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Forms[formid].Elements, id)

	records, err := models.CrudRead(tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud")
	}
	if len(records) == 0 {
		flash.Error("Article non trouvé")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud")
	}
	// // Calcul des éléments (valeur par défaut comprise)
	elements = computeElements(c.Controller, false, elements, records[0])

	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	form := app.Tables[tableid].Forms[formid]

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
				c.Ctx.Redirect(302, "/crud")
				return
			}
		} else {
			flash.Error("Vue non trouvée :", viewid)
			flash.Store(&c.Controller)
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		flash.Error("Application non trouvée :", appid)
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud")
		return
	}

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Forms[formid].Elements, id)

	// lecture du record
	records, err := models.CrudRead(tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud")
		return
	}

	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	form := app.Tables[tableid].Forms[formid]

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
		c.Ctx.Redirect(302, "/crud")
		return
	}
	// PostSQL
	for _, postsql := range form.PostSQL {
		sql := macro(c.Controller, postsql, orm.Params{})
		if sql != "" {
			err = models.CrudExec(sql, table.AliasDB)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
			}
		} else {
			beego.Error("Ordre sql incorrect ", postsql)
			flash.Error("Ordre sql incorrect ", postsql)
			flash.Store(&c.Controller)
			c.Ctx.Redirect(302, "/crud")
			return
		}
	}

	flash.Notice("Mise à jour effectuée avec succès")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
}
