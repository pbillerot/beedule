package controllers

import (
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
	id := c.Ctx.Input.Param(":id")

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		logs.Error("App not found", c.GetSession("Username").(string), appid)
		backward(c.Controller)
		return
	}
	if val, ok := dico.Ctx.Applications[appid].Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			logs.Error("View not found", c.GetSession("Username").(string), viewid)
			backward(c.Controller)
			return
		}
	} else {
		logs.Error("table not found", c.GetSession("Username").(string), tableid)
		backward(c.Controller)
		return
	}
	// Contrôle d'accès
	table := dico.Ctx.Applications[appid].Tables[tableid]
	view := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, appid, id) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, _ := mergeElements(c.Controller, appid, tableid, dico.Ctx.Applications[appid].Tables[tableid].Views[viewid].Elements, id)

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
	records, err := models.CrudRead(filter, appid, tableid, id, elements)
	if err != nil {
		flash.Error("%s", err.Error())
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé: %v", id)
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	// Suppression de l'enregistrement
	err = models.CrudDelete(appid, tableid, id)
	if err != nil {
		flash.Error("%s", err.Error())
		flash.Store(&c.Controller)
	} else {
		// PostSQL
		for _, postsql := range view.PostSQL {
			sql := macro(c.Controller, appid, postsql, records[0])
			if sql != "" {
				err = models.CrudExec(sql, table.Setting.AliasDB)
				if err != nil {
					flash.Error("%s", err.Error())
					flash.Store(&c.Controller)
				}
			}
		}
	}
	backward(c.Controller)
	// c.DelSession(fmt.Sprintf("anch_%s_%s", tableid, viewid))
	// c.Ctx.Redirect(302, "/bee/list/"+appid+"/"+tableid+"/"+viewid)
}
