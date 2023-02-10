package controllers

import (
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
)

// RestController implements global settings for all other routers.
type CrudSortController struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *CrudSortController) Prepare() {
}

type sortResponse struct {
	Response string // ok or ko
	Message  string
}

// RestIsc IsConnected
func (c *CrudSortController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")

	// enregistrement du tri dans la session
	sortid := strings.ToLower(c.GetString("sortid"))
	ctxSortID := fmt.Sprintf("%s-%s-%s-sortid", appid, tableid, viewid)
	if sortid != "" {
		c.SetSession(ctxSortID, sortid)
	} else {
		c.DelSession(ctxSortID)
	}
	sortdirection := strings.ToLower(c.GetString("sortdirection"))
	ctxSortDirection := fmt.Sprintf("%s-%s-%s-sortdirection", appid, tableid, viewid)
	if sortdirection != "" {
		c.SetSession(ctxSortDirection, sortdirection)
	} else {
		c.DelSession(ctxSortDirection)
	}

	var rest sortResponse
	rest.Response = "ok"
	rest.Message = sortid + " " + sortdirection

	c.Data["json"] = &rest
	c.ServeJSON()
}
