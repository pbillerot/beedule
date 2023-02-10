package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/types"
)

// CrudShareController implements global settings for all other routers.
type CrudShareController struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *CrudShareController) Prepare() {
}

// CrudShareApp appel d'une application partagée par un anonyme
func (c *CrudShareController) CrudShareApp() {
	appid := c.Ctx.Input.Param(":appid")
	sessionid := c.Ctx.Input.Param(":shareid")

	// recherche de shareid dans la table share
	if dico.Ctx.IsShared(appid, sessionid) {
		// C'est OK enregistrement du compte dans la session
		user, _ := GetUser("anonymous")
		flash := beego.ReadFromRequest(&c.Controller)
		if err != nil {
			flash.Error("Compte anonymous inconnu")
			flash.Store(&c.Controller)
			setContext(c.Controller, "admin", "users")
			c.TplName = "crud_login.html"
			return
		}

		group := dico.Ctx.Applications[appid].Group
		c.SetSession("SessionId", sessionid)
		c.SetSession("LoggedIn", true)
		c.SetSession("Username", user.Username)
		c.SetSession("Groups", group)
		c.SetSession("AppId", appid) // pour empêcher l'anonymous d'aller sur une autre application
		c.SetSession("IsAdmin", false)
		session := types.Session{}
		session.SessionID = sessionid
		session.LoggedIn = true
		session.Username = user.Username
		session.Groups = group
		session.AppID = appid
		session.IsAdmin = false
		c.Data["Session"] = &session

		navigateInit(c.Controller)
		logs.Info(fmt.Sprintf("CONNEXION de [%s] groupe:[%s]", user.Username, group))

		// recherche de la 1ère vue à lancer
		c.Ctx.Redirect(302, fmt.Sprintf("/bee/list/%s/%s/%s",
			appid,
			dico.Ctx.Applications[appid].Menu[0].TableID,
			dico.Ctx.Applications[appid].Menu[0].ViewID))

	} else {
		c.Ctx.Redirect(302, "/bee/login")
	}

	c.Ctx.Redirect(302, "/bee/login")
}
