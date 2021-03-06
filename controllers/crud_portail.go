package controllers

import (
	beego "github.com/beego/beego/v2/adapter"
)

// CrudPortailController as
type CrudPortailController struct {
	loggedRouter
}

// Get CrudPortailController
func (c *CrudPortailController) Get() {
	setContext(c.Controller, "admin", "users")
	flash := beego.ReadFromRequest(&c.Controller)
	flash.Store(&c.Controller)
	navigateInit(c.Controller)
	c.TplName = "crud_portail.html"
}
