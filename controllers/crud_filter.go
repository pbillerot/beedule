package controllers

import (
	"fmt"
	"strings"

	"github.com/pbillerot/beedule/dico"

	beego "github.com/beego/beego/v2/adapter"
)

// RestController implements global settings for all other routers.
type CrudFilterController struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *CrudFilterController) Prepare() {
}

// RestIsc IsConnected
func (c *CrudFilterController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	view := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid]

	// enregistrement des filtres dans la session
	reset := c.GetString("resetfilter")
	for _, keyFilter := range view.Filters {
		filter := strings.ToLower(c.GetString(keyFilter))
		ctxFilter := fmt.Sprintf("%s-%s-%s-filter-%s", appid, tableid, viewid, keyFilter)
		if reset == "reset" {
			filter = ""
		}
		if filter != "" {
			c.SetSession(ctxFilter, filter)
		} else {
			c.DelSession(ctxFilter)
		}
	}

	c.Ctx.Redirect(302, "/bee/list/"+appid+"/"+tableid+"/"+viewid)
}
