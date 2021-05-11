package controllers

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/batch"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"

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
				action := macro(c.Controller, view.Actions[iactionid].Plugin, orm.Params{})
				batch.RunPlugin(action)
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
				withPlugin = true
				action := macro(c.Controller, form.Actions[iactionid].Plugin, orm.Params{})
				batch.RunPlugin(action)
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

	setContext(c.Controller, tableid)
	var withPlugin bool

	// Si un formView est défini on utilisera son modèle pour les éléments
	form := app.Tables[tableid].Forms[formid]
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, id) {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	var elementsVF types.Elements
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
						err = models.CrudExec(sql, table.AliasDB)
						if err != nil {
							flash.Error(err.Error())
							flash.Store(&c.Controller)
						}
					}
				}
			}
			// Appel Plugin
			if err == nil {
				if action.Plugin != "" {
					withPlugin = true
					if element.Params.WithInputFile {
						file, handler, err := c.Ctx.Request.FormFile(actionid)
						if err != nil {
							beego.Error(err)
						} else {
							fileBytes, err := ioutil.ReadAll(file)
							if err != nil {
								beego.Error(err)
							} else {
								filepath := fmt.Sprintf("%s/%s", element.Params.Path, handler.Filename)
								beego.Info("ajout", filepath)
								err := ioutil.WriteFile(
									filepath,
									fileBytes, 0755)
								if err != nil {
									beego.Error(err)
								}
							}
						}
					}
					command := macro(c.Controller, action.Plugin, records[0])
					_, err = batch.RunPlugin(command)
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
