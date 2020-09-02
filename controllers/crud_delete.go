package controllers

import (
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

	// Ctrl tableid et viewid
	if table, ok := app.Tables[tableid]; ok {
		if view, ok := table.Views[viewid]; ok {
			if !view.Deletable {
				flash.Error("Suppression non autorisée")
				flash.Store(&c.Controller)
				c.Ctx.Redirect(302, "/crud")
			}
		} else {
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		c.Ctx.Redirect(302, "/crud")
		return
	}
	//
	err = models.CrudDelete(tableid, id)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	c.Ctx.Redirect(302, "/crud/list/"+appid+"/"+tableid+"/"+viewid)
}
