package controllers

import (
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
)

// RestController implements global settings for all other routers.
type CrudSearchController struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *CrudSearchController) Prepare() {
}

type searchResponse struct {
	Response string // ok or ko
	Message  string
}

// RestIsc IsConnected
func (c *CrudSearchController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")

	// enregistrement de la recherche dans la session
	search := strings.ToLower(c.GetString("search"))
	ctxSearch := fmt.Sprintf("%s-%s-%s-search", appid, tableid, viewid)
	if search != "" {
		c.SetSession(ctxSearch, search)
	} else {
		c.DelSession(ctxSearch)
	}

	var rest searchResponse
	rest.Response = "ok"
	rest.Message = search

	c.Data["json"] = &rest
	c.ServeJSON()
}
