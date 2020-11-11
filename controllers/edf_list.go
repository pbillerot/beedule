package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/pbillerot/beedule/app"
)

// EdfListController as
type EdfListController struct {
	loggedRouter
}

// EdfList CrudListController
func (c *EdfListController) EdfList() {
	appid := c.Ctx.Input.Param(":app")
	dirid := c.Ctx.Input.Param(":dir")
	baseid := c.Ctx.Input.Param(":base")

	// flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := app.Applications[appid]; !ok {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}

	setContextEdf(c.Controller, appid)

	// Remplissage du contexte pour le template
	c.Data["AppId"] = appid
	c.Data["DirId"] = dirid
	c.Data["BaseId"] = baseid
	c.Data["Search"] = ""

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/edf/list/%s", appid))

	c.TplName = "edf_list.html"
}
