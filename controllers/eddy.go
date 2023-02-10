package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"
)

// Eddydoc table
type eddyFile struct {
	Path    string
	Content string
}

// EddyController controleur de l'éditeur de fichier eddy
type EddyController struct {
	loggedRouter
}

// EddyDocument Visualiser Modifier un document
func (c *EddyController) EddyDocument() {
	flash := beego.ReadFromRequest(&c.Controller)
	// keyid = nom du fichier
	keyid := c.Ctx.Input.Param(":key")
	appid := c.GetSession("AppId").(string)

	pathFile := dico.Ctx.Applications[appid].DicoDir + "/" + keyid
	if keyid == "portail.yaml" {
		pathFile = beego.AppConfig.String("portail")
	}

	if c.Ctx.Input.Method() == "POST" {
		// ENREGISTREMENT DU DOCUMENT
		content := c.GetString("document")
		err = ioutil.WriteFile(pathFile, []byte(content), 0644)
		if err != nil {
			msg := fmt.Sprintf("EddyDocument %s : %s", pathFile, err)
			beego.Error(msg)
			flash.Error(msg)
			flash.Store(&c.Controller)
		}
		// Demande d'actualisation de l'arborescence
		dico.Ctx = dico.Portail{}
		msg, err := dico.Ctx.Load()
		if err != nil {
			flash.Error(strings.Join(msg[:], ","))
			flash.Store(&c.Controller)
		}
		c.Ctx.Output.Cookie("eddy-refresh", "true")
	}

	// Lecture du fichier
	var record eddyFile
	content, err := ioutil.ReadFile(pathFile)
	if err != nil {
		msg := fmt.Sprintf("EddyDocument %s : %s", keyid, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
	}
	record.Content = string(content)
	record.Path = keyid

	// Remplissage du contexte pour le template
	// XSRF protection des formulaires
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf"] = c.XSRFToken()
	c.Data["AppId"] = appid
	c.Data["Portail"] = &dico.Ctx
	c.Data["Files"] = dico.Ctx.Applications[appid].Files
	c.Data["Config"] = &models.Config
	// XSRF protection des formulaires
	c.Data["Record"] = record
	c.Data["KeyID"] = keyid
	c.Data["TabTitle"] = keyid
	if path.Ext(keyid) == ".js" {
		c.Data["ModeMarkdown"] = "javascript"
	} else {
		c.Data["ModeMarkdown"] = strings.Replace(path.Ext(keyid), ".", "", -1)
	}
	// load liste des rubriques
	var rubriques string
	if path.Ext(keyid) == ".yaml" && keyid != "portail.yaml" && keyid != "application.yaml" {
		tableid := strings.TrimSuffix(keyid, path.Ext(keyid))
		for k := range dico.Ctx.Applications[appid].Tables[tableid].Elements {
			if len(rubriques) != 0 {
				rubriques += ","
			}
			rubriques += k
		}
	}
	c.Data["Rubriques"] = rubriques
	c.TplName = "eddy.html"
}

// EddyLog Visualiser Modifier un document
func (c *EddyController) EddyLog() {
	flash := beego.ReadFromRequest(&c.Controller)
	// keyid = nom du fichier
	keyid := models.Config.LogPath
	appid := c.GetSession("AppId").(string)

	// Lecture du fichier
	var record eddyFile
	content, err := ioutil.ReadFile(keyid)
	if err != nil {
		msg := fmt.Sprintf("EddyLog %s : %s", keyid, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
	}
	record.Content = string(content)
	record.Path = keyid

	// Remplissage du contexte pour le template
	// XSRF protection des formulaires
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf"] = c.XSRFToken()
	c.Data["AppId"] = appid
	c.Data["Portail"] = &dico.Ctx
	c.Data["Files"] = dico.Ctx.Applications[appid].Files
	c.Data["Config"] = &models.Config
	c.Data["Record"] = record
	c.Data["KeyID"] = keyid
	c.Data["TabTitle"] = keyid
	c.Data["ModeMarkdown"] = ""
	c.Data["Rubriques"] = ""
	c.TplName = "log.html"
}
