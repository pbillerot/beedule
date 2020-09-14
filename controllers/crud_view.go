package controllers

import (
	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"

	"github.com/astaxie/beego"
)

// CrudViewController as
type CrudViewController struct {
	loggedRouter
}

// Get CrudViewController appel du formulaire formView
func (c *CrudViewController) Get() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	id := c.Ctx.Input.Param(":id")

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl tableid et viewid
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		c.Ctx.Redirect(302, "/crud")
		return
	}

	// Si un formView est défini on utilisera son modèle pour les éléments
	formviewid := app.Tables[tableid].Views[viewid].FormView
	var elementsVF types.Elements
	if formviewid == "" {
		elementsVF = app.Tables[tableid].Views[viewid].Elements
	} else {
		elementsVF = app.Tables[tableid].Forms[formviewid].Elements
	}
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, cols := mergeElements(c.Controller, tableid, elementsVF, id)

	// lecture du record
	records, err := models.CrudRead(tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	// Calcul des éléments
	elements = computeElements(c.Controller, false, elements, records[0])

	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	if len(records) == 0 {
		flash.Error("Article non trouvé: ", id)
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/list/"+appid+"/"+tableid+"/"+viewid)
		return
	}

	setContext(c.Controller)
	if err == nil {
		c.Data["ColDisplay"] = records[0][table.ColDisplay].(string)
	} else {
		c.Data["ColDisplay"] = "Article non trouvé"
	}
	c.Data["AppId"] = appid
	c.Data["Application"] = app.Applications[appid]
	c.Data["Id"] = id
	c.Data["TableId"] = tableid
	c.Data["ViewId"] = viewid
	c.Data["FormView"] = app.Tables[tableid].Forms[formviewid]
	c.Data["FormId"] = app.Tables[tableid].Views[viewid].FormEdit
	c.Data["Table"] = &table
	c.Data["View"] = &view
	c.Data["Elements"] = &elements
	c.Data["Records"] = &records
	c.Data["Cols"] = &cols

	c.TplName = "crud_view.html"
}
