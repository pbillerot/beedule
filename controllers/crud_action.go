package controllers

import (
	"strconv"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/batch"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// CrudActionViewController as
type CrudActionViewController struct {
	loggedRouter
}

// Post CrudActionViewController
func (c *CrudActionViewController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	actionid := c.Ctx.Input.Param(":action")

	flash := beego.NewFlash()

	// Ctrl appid tableid viewid formid
	if _, ok := app.Applications[appid]; !ok {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			beego.Error("View not found", c.GetSession("Username").(string), viewid)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		beego.Error("Table not found", c.GetSession("Username").(string), tableid)
		ReturnFrom(c.Controller)
		return
	}

	// Contrôle d'accès
	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, actionid) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	iactionid, err := strconv.Atoi(actionid)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	if iactionid <= len(view.Actions) {
		// Exécution des ordres SQL
		for _, action := range view.Actions[iactionid].SQL {
			sql := macro(c.Controller, action, orm.Params{})
			if sql != "" {
				err = models.CrudExec(sql, table.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c.Controller)
				}
			}
		}
		// Appel duPlugin
		if err == nil {
			if view.Actions[iactionid].Plugin != "" {
				batch.RunPlugin(view.Actions[iactionid].Plugin)
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
	}

	c.Ctx.Redirect(302, "/crud/list/"+appid+"/"+tableid+"/"+viewid)
}

// CrudActionFormController as
type CrudActionFormController struct {
	loggedRouter
}

// Post CrudActionFormController
func (c *CrudActionFormController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	id := c.Ctx.Input.Param(":id")
	actionid := c.Ctx.Input.Param(":action")

	flash := beego.NewFlash()

	// Ctrl appid tableid viewid formid
	if _, ok := app.Applications[appid]; !ok {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				beego.Error("Form not found", c.GetSession("Username").(string), formid)
				ReturnFrom(c.Controller)
				return
			}
		} else {
			beego.Error("View not found", c.GetSession("Username").(string), viewid)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		beego.Error("table not found", c.GetSession("Username").(string), tableid)
		ReturnFrom(c.Controller)
		return
	}

	// Contrôle d'accès
	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	form := app.Tables[tableid].Forms[formid]
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, actionid) {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	iactionid, err := strconv.Atoi(actionid)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}
	if iactionid <= len(form.Actions) {
		// Exécution des ordres SQL
		for _, action := range form.Actions[iactionid].SQL {
			sql := macro(c.Controller, action, orm.Params{})
			if sql != "" {
				err = models.CrudExec(sql, table.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c.Controller)
				}
			}
		}
		// Appel du Plugin
		if err == nil {
			if form.Actions[iactionid].Plugin != "" {
				batch.RunPlugin(form.Actions[iactionid].Plugin)
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
	}

	c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
}

// CrudActionElementController as
type CrudActionElementController struct {
	loggedRouter
}

// Post CrudActionElementController
func (c *CrudActionElementController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	id := c.Ctx.Input.Param(":id")
	actionid := c.Ctx.Input.Param(":action") // l'id de l'élément

	flash := beego.NewFlash()

	// Ctrl appid tableid viewid formid
	if _, ok := app.Applications[appid]; !ok {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			beego.Error("View not found", c.GetSession("Username").(string), viewid)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		beego.Error("Table not found", c.GetSession("Username").(string), tableid)
		ReturnFrom(c.Controller)
		return
	}

	// Contrôle d'accès
	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, actionid) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	// Si un formView est défini on utilisera son modèle pour les éléments
	formviewid := app.Tables[tableid].Views[viewid].FormView
	var elementsVF types.Elements
	if formviewid == "" {
		flash.Error("Enregistrement non trouvé")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}
	elementsVF = app.Tables[tableid].Forms[formviewid].Elements
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, _ := mergeElements(c.Controller, tableid, elementsVF, id)

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
	}
	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}
	// Calcul des éléments
	elements = computeElements(c.Controller, true, elements, records[0])

	if element, ok := elements[actionid]; ok {
		// Exécution des ordres SQL
		for _, action := range element.Action.SQL {
			sql := macro(c.Controller, action, orm.Params{})
			if sql != "" {
				err = models.CrudExec(sql, table.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c.Controller)
				}
			}
		}
		// Appel du Plugin
		if err == nil {
			if elements[actionid].Action.Plugin != "" {
				batch.RunPlugin(elements[actionid].Action.Plugin)
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}

	c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
}
