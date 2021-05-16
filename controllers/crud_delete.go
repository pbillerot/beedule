package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"

	beego "github.com/beego/beego/v2/adapter"
)

// CrudDeleteController as
type CrudDeleteController struct {
	loggedRouter
}

// Post CrudDeleteController
func (c *CrudDeleteController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	id := c.Ctx.Input.Param(":id")

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		logs.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}
	if val, ok := dico.Ctx.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				logs.Error("Form not found", c.GetSession("Username").(string), formid)
				ReturnFrom(c.Controller)
				return
			}
		} else {
			logs.Error("View not found", c.GetSession("Username").(string), viewid)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		logs.Error("table not found", c.GetSession("Username").(string), tableid)
		ReturnFrom(c.Controller)
		return
	}
	// Contrôle d'accès
	table := dico.Ctx.Tables[tableid]
	view := dico.Ctx.Tables[tableid].Views[viewid]
	form := dico.Ctx.Tables[tableid].Forms[formid]
	if view.Group == "" {
		view.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, id) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, id) {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, _ := mergeElements(c.Controller, tableid, dico.Ctx.Tables[tableid].Forms[formid].Elements, id)

	// Filtrage si élément owner
	filter := ""
	for key, element := range elements {
		// Un seule élément owner par enregistrement
		if element.Group == "owner" && !IsAdmin(c.Controller) {
			filter = key + " = '" + c.GetSession("Username").(string) + "'"
			break
		}
	}

	// lecture du record
	records, err := models.CrudRead(filter, tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé: ", id)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	// Suppression de l'enregistrement
	err = models.CrudDelete(tableid, id)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		// PostSQL
		for _, postsql := range form.PostSQL {
			sql := macro(c.Controller, postsql, records[0])
			if sql != "" {
				err = models.CrudExec(sql, table.Setting.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c.Controller)
				}
			}
		}
	}
	c.DelSession(fmt.Sprintf("anch_%s_%s", tableid, viewid))
	c.Ctx.Redirect(302, "/bee/list/"+appid+"/"+tableid+"/"+viewid)
}
