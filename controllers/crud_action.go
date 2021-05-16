package controllers

import (
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
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

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		logs.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}
	if val, ok := dico.Ctx.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			logs.Error("View not found", c.GetSession("Username").(string), viewid)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		logs.Error("Table not found", c.GetSession("Username").(string), tableid)
		ReturnFrom(c.Controller)
		return
	}

	// Contrôle d'accès
	table := dico.Ctx.Tables[tableid]
	view := dico.Ctx.Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, actionid) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	setContext(c.Controller, tableid)

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
				err = models.CrudExec(sql, table.Setting.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c.Controller)
				}
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
	}

	c.Ctx.Redirect(302, "/bee/list/"+appid+"/"+tableid+"/"+viewid)
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
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, actionid) {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	setContext(c.Controller, tableid)
	var withPlugin bool

	iactionid, err := strconv.Atoi(actionid)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/bee/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}
	if iactionid <= len(form.Actions) {
		// Exécution des ordres SQL
		for _, action := range form.Actions[iactionid].SQL {
			sql := macro(c.Controller, action, orm.Params{})
			if sql != "" {
				err = models.CrudExec(sql, table.Setting.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c.Controller)
				}
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
	}

	if withPlugin {
		c.Ctx.Redirect(302, "/bee/list/"+appid+"/"+tableid+"/"+viewid)
	} else {
		c.Ctx.Redirect(302, "/bee/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
	}
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
	formid := c.Ctx.Input.Param(":form")
	id := c.Ctx.Input.Param(":id")
	actionid := c.Ctx.Input.Param(":action") // l'id de l'élément

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		logs.Error("App not found", c.GetSession("Username").(string), appid)
		ReturnFrom(c.Controller)
		return
	}
	if val, ok := dico.Ctx.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			logs.Error("View not found", c.GetSession("Username").(string), viewid)
			ReturnFrom(c.Controller)
			return
		}
	} else {
		logs.Error("Table not found", c.GetSession("Username").(string), tableid)
		ReturnFrom(c.Controller)
		return
	}

	// Contrôle d'accès
	table := dico.Ctx.Tables[tableid]
	view := dico.Ctx.Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, actionid) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	setContext(c.Controller, tableid)
	var withPlugin bool

	// Si un formView est défini on utilisera son modèle pour les éléments
	form := dico.Ctx.Tables[tableid].Forms[formid]
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
	var elementsVF map[string]dico.Element
	elementsVF = form.Elements
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
		c.Ctx.Redirect(302, "/bee/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}
	// Calcul des éléments
	elements = computeElements(c.Controller, true, elements, records[0])

	if element, ok := elements[actionid]; ok {
		// Exécution des ordres SQL
		for _, action := range element.Actions {
			if err == nil {
				for _, actionSQL := range action.SQL {
					sql := macro(c.Controller, actionSQL, records[0])
					if sql != "" {
						err = models.CrudExec(sql, table.Setting.AliasDB)
						if err != nil {
							flash.Error(err.Error())
							flash.Store(&c.Controller)
						}
					}
				}
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/bee/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}

	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	if withPlugin {
		c.Ctx.Redirect(302, "/bee/list/"+appid+"/"+tableid+"/"+viewid)
	} else {
		c.Ctx.Redirect(302, "/bee/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
	}
}
