package controllers

import (
	"github.com/astaxie/beego"
)

// CrudPortailController as
type CrudPortailController struct {
	loggedRouter
}

// Get CrudPortailController
func (c *CrudPortailController) Get() {
	setContext(c.Controller, "parameters")
	flash := beego.ReadFromRequest(&c.Controller)
	flash.Store(&c.Controller)
	c.TplName = "crud_portail.html"
}
