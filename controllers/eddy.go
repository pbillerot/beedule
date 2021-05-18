package controllers

import (
	"fmt"
	"io/ioutil"
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
	keyid := c.Ctx.Input.Param(":key")

	if c.Ctx.Input.Method() == "POST" {
		// ENREGISTREMENT DU DOCUMENT
		document := c.GetString("document")
		err = ioutil.WriteFile("./config/"+keyid, []byte(document), 0644)
		if err != nil {
			msg := fmt.Sprintf("EddyDocument %s : %s", "./config/"+keyid, err)
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
	content, err := ioutil.ReadFile("./config/" + keyid)
	if err != nil {
		msg := fmt.Sprintf("EddyDocument %s : %s", keyid, err)
		beego.Error(msg)
		flash.Error(msg)
		flash.Store(&c.Controller)
	}
	record.Content = string(content)
	record.Path = keyid

	// Remplissage du contexte pour le template
	c.Data["Record"] = record
	c.Data["KeyID"] = keyid
	c.Data["TabTitle"] = keyid
	c.TplName = "eddy.html"
}
