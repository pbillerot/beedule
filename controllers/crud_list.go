package controllers

import (
	"fmt"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// CrudListController as
type CrudListController struct {
	loggedRouter
}

// Get CrudListController
func (c *CrudListController) Get() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := app.Applications[appid]; !ok {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			beego.Error("View not found", c.GetSession("Username").(string), viewid)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		beego.Error("Table not found", c.GetSession("Username").(string), tableid)
		ReturnFrom(c.Controller)
		return
	}

	// Contrôle d'accès à la vue
	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, "") {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), viewid, view.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	// Ctrl d'accès FormAdd FormView FormEdit
	if !IsInGroup(c.Controller, table.Forms[view.FormView].Group, "") {
		view.FormView = ""
	}
	if !IsInGroup(c.Controller, table.Forms[view.FormAdd].Group, "") {
		view.FormAdd = ""
	}
	if !IsInGroup(c.Controller, table.Forms[view.FormEdit].Group, "") {
		view.FormEdit = ""
	}

	// Fusion des attributs des éléments de la table dans les éléments de la vue
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Views[viewid].Elements, "")
	// Calcul des champs SQL
	if view.OrderBy != "" {
		view.OrderBy = macro(c.Controller, view.OrderBy, orm.Params{})
	}
	if view.FooterSQL != "" {
		view.FooterSQL = requeteSQL(c.Controller, view.OrderBy, orm.Params{}, app.Tables[tableid].AliasDB)
	}
	if view.Where != "" {
		view.Where = macro(c.Controller, view.Where, orm.Params{})
	}

	// lecture des records
	records, err := models.CrudList(tableid, viewid, &view, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		// c.Ctx.Redirect(302, "/crud")
		// return
	}
	if len(records) > 0 {
		// Calcul des éléments hors values
		elements = computeElements(c.Controller, false, elements, records[0])
	}

	// Remplissage du contexte pour le template
	c.SetSession("from", fmt.Sprintf("/crud/list/%s/%s/%s", appid, tableid, viewid))
	setContext(c.Controller)
	c.Data["Title"] = view.Title
	c.Data["AppId"] = appid
	c.Data["Application"] = app.Applications[appid]
	c.Data["TableId"] = tableid
	c.Data["ViewId"] = viewid
	c.Data["Table"] = &table
	c.Data["View"] = &view
	c.Data["Elements"] = elements
	c.Data["Records"] = records
	c.Data["Qrecords"] = len(records)
	c.Data["Cols"] = cols

	if view.Type == "image" {
		c.TplName = "crud_list_image.html"
	} else if view.Type == "table" {
		c.TplName = "crud_table.html"
	} else {
		c.TplName = "crud_list.html"
	}
}
