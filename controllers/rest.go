package controllers

import (
	"github.com/astaxie/beego"
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
	beego.Info(url)
	c.Data["json"] = map[string]interface{}{"url": url}
	c.ServeJSON()
}
