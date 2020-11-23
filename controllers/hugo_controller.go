package controllers

import (
	"html/template"
	"time"

	"github.com/astaxie/beego"
	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/types"
)

// HugoController controleur du gestionnaire de fichier EDF
type HugoController struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *HugoController) Prepare() {
	if c.GetSession("LoggedIn") != nil {
		if c.GetSession("LoggedIn").(bool) != true {
			c.Ctx.Redirect(302, "/bee/login")
		}
	} else {
		c.Ctx.Redirect(302, "/bee/login")
	}
	appid := c.Ctx.Input.Param(":app")
	// CTRL APPLICATION
	flash := beego.ReadFromRequest(&c.Controller)
	if _, ok := app.Applications[appid]; !ok {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		c.Ctx.Redirect(302, "/bee/login")
		return
	}
	if !IsInGroup(c.Controller, app.Applications[appid].Group, "") {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), appid)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/bee/login")
		return
	}

	if appid != "" {
		c.Data["TabIcon"] = app.Applications[appid].Image
		c.Data["TabTitle"] = app.Applications[appid].Title
	} else {
		c.Data["TabIcon"] = app.Portail.IconFile
		c.Data["TabTitle"] = app.Portail.Title
	}
	c.Data["AppId"] = appid

	// Contexte de l'application
	c.Data["Application"] = app.Applications[appid]

	// Context lié à custom.conf
	c.Data["DataDir"] = beego.AppConfig.String(appid + "::datadir")
	c.Data["DataUrl"] = "/bee/data/" + appid
	c.Data["UrlTest"] = beego.AppConfig.String(appid + "::urltest")

	// Contexte de la session
	session := types.Session{}
	if c.GetSession("LoggedIn") != nil {
		session.LoggedIn = c.GetSession("LoggedIn").(bool)
	}
	if c.GetSession("Username") != nil {
		session.Username = c.GetSession("Username").(string)
	}
	if c.GetSession("IsAdmin") != nil {
		session.IsAdmin = c.GetSession("IsAdmin").(bool)
	}
	if c.GetSession("Groups") != nil {
		session.Groups = c.GetSession("Groups").(string)
	}
	c.Data["Session"] = &session

	// Identité du framework Beedule
	config := types.Config{}
	config.Appname = beego.AppConfig.String("appname")
	config.Appnote = beego.AppConfig.String("appnote")
	config.Date = beego.AppConfig.String("date")
	config.Icone = beego.AppConfig.String("icone")
	config.Site = beego.AppConfig.String("site")
	config.Email = beego.AppConfig.String("email")
	config.Author = beego.AppConfig.String("author")
	config.Version = beego.AppConfig.String("version")
	config.Theme = beego.AppConfig.String("theme")
	c.Data["Config"] = &config

	// XSRF protection des formulaires
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	// Title
	c.Data["Title"] = config.Appname
	c.Data["Portail"] = &app.Portail
	// Contexte de navigation
	c.Data["From"] = c.Ctx.Input.Cookie("from")
	c.Data["URL"] = c.Ctx.Request.URL.String()

	// Sera ajouté derrière les urls pour ne pas utiliser le cache des images dynamiques
	c.Data["Composter"] = time.Now().Unix()
}
