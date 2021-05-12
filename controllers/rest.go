package controllers

import (
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/core/logs"
)

// RestController implements global settings for all other routers.
type RestController struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *RestController) Prepare() {
}

type dataRest struct {
	User string `json:"userid"`
}

// RestIsc IsConnected
func (c *RestController) RestIsc() {
	var rest dataRest
	if c.GetSession("Username") != nil {
		rest.User = c.GetSession("Username").(string)
	}
	c.Data["json"] = &rest
	c.ServeJSON()
}

// RestPutLog Log
func (c *RestController) RestPutLog() {
	url := c.Ctx.Input.Param(":url")
	logs.Info(url)
	c.Data["json"] = map[string]interface{}{"url": url}
	c.ServeJSON()
}
