package controllers

import (
	beego "github.com/beego/beego/v2/adapter"
	"github.com/pbillerot/beedule/dico"
)

// CrudPortailController as
type CrudPortailController struct {
	loggedRouter
}

// Get CrudPortailController
func (c *CrudPortailController) Get() {
	// c'est l'application admin avec la table users qui gère le contenu du portail
	sessionID := setContext(c.Controller, "admin", "users")
	flash := beego.ReadFromRequest(&c.Controller)
	flash.Store(&c.Controller)
	navigateInit(c.Controller)

	// remplissage de ShareApps
	dico.Ctx.ShareUpdate(sessionID)

	c.TplName = "crud_portail.html"
}
