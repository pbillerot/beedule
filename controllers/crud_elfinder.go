package controllers

import (
	"github.com/LeeEirc/elfinder"
	"github.com/pbillerot/beedule/app"

	"github.com/astaxie/beego"
)

// CrudElFinderController as
type CrudElFinderController struct {
	loggedRouter
}

// CrudElFinder Controller
func (c *CrudElFinderController) CrudElFinder() {
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
	view := table.Views[viewid]
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

	volumes := []elfinder.Volume{
		elfinder.NewLocalVolume("/home/billerot/Abri/foirexpo/content"),
		elfinder.NewLocalVolume("/home/billerot/Abri/foirexpo/data"),
	}

	con := elfinder.NewElFinderConnector(volumes)
	con.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)

	return
}
