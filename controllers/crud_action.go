package controllers

import (
	"strconv"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// CrudActionController as
type CrudActionController struct {
	loggedRouter
}

// Post CrudActionController
func (c *CrudActionController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	actionid := c.Ctx.Input.Param(":action")

	flash := beego.NewFlash()

	// Ctrl tableid et viewid
	if table, ok := app.Tables[tableid]; ok {
		if _, ok := table.Views[viewid]; ok {
		} else {
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		c.Ctx.Redirect(302, "/crud")
		return
	}
	table := app.Tables[tableid]
	view := table.Views[viewid]

	iactionid, err := strconv.Atoi(actionid)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/list/"+appid+"/"+tableid+"/"+viewid)
		return
	}
	if iactionid <= len(view.ActionsSQL) {
		sql := macro(c.Controller, view.ActionsSQL[iactionid].SQL, orm.Params{})
		if sql != "" {
			err = models.CrudExec(sql, table.AliasDB)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
	}

	c.Ctx.Redirect(302, "/crud/list/"+appid+"/"+tableid+"/"+viewid)
}
