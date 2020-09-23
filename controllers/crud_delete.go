package controllers

import (
	"fmt"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"

	"github.com/astaxie/beego"
)

// CrudDeleteController as
type CrudDeleteController struct {
	loggedRouter
}

// Post CrudDeleteController
func (c *CrudDeleteController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	id := c.Ctx.Input.Param(":id")

	flash := beego.NewFlash()

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

	// Contrôle d'accès
	view := app.Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, id) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	// Suppression de l'enregistrement
	err = models.CrudDelete(tableid, id)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	c.DelSession(fmt.Sprintf("anch_%s_%s", tableid, viewid))
	c.Ctx.Redirect(302, "/crud/list/"+appid+"/"+tableid+"/"+viewid)
}
