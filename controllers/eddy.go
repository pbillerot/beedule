package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/pbillerot/beedule/dico"
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
	// Mémorisation de la position du curseur
	cursorCh := c.GetString("cursor_ch")
	cursorLine := c.GetString("cursor_line")

	if c.Ctx.Input.Method() == "POST" {
		// ENREGISTREMENT DU DOCUMENT
		document := c.GetString("document")
		err = ioutil.WriteFile(beego.AppConfig.String("dicodir")+"/"+keyid, []byte(document), 0644)
		if err != nil {
			msg := fmt.Sprintf("EddyDocument %s : %s", beego.AppConfig.String("dicodir")+"/"+keyid, err)
			beego.Error(msg)
			flash.Error(msg)
			flash.Store(&c.Controller)
		}
		// Demande d'actualisation de l'arborescence
		msg, err := dico.Ctx.Load()
		if err != nil {
			flash.Error(strings.Join(msg[:], ","))
			flash.Store(&c.Controller)
		}
		c.Ctx.Output.Cookie("eddy-refresh", "true")
	}

	// Lecture du fichier
	var record eddyFile
	content, err := ioutil.ReadFile(beego.AppConfig.String("dicodir") + "/" + keyid)
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
	c.Data["Record"] = record
	c.Data["KeyID"] = keyid
	c.Data["TabTitle"] = keyid
	c.Data["CursorCh"] = cursorCh
	c.Data["CursorLine"] = cursorLine
	// load liste des rubriques
	tableid := strings.TrimSuffix(keyid, path.Ext(keyid))
	var rubriques string
	for k := range dico.Ctx.Tables[tableid].Elements {
		if len(rubriques) != 0 {
			rubriques += ","
		}
		rubriques += k
	}
	c.Data["Rubriques"] = rubriques
	c.TplName = "eddy.html"
}
