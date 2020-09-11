package controllers

import (
	"strconv"

	"github.com/pbillerot/beedule/app"
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

	// Ctrl tableid et viewid
	if table, ok := app.Tables[tableid]; ok {
		if _, ok := table.Views[viewid]; ok {
		} else {
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		c.Ctx.Redirect(302, "/crud")
		return
	}
	table := app.Tables[tableid]
	view := table.Views[viewid]

	iactionid, err := strconv.Atoi(actionid)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/list/"+appid+"/"+tableid+"/"+viewid)
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

	// Ctrl tableid et viewid
	if table, ok := app.Tables[tableid]; ok {
		if _, ok := table.Forms[formid]; ok {
		} else {
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		c.Ctx.Redirect(302, "/crud")
		return
	}
	table := app.Tables[tableid]
	form := table.Views[formid]

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

	// Ctrl tableid et viewid
	if val, ok := app.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			c.Ctx.Redirect(302, "/crud")
			return
		}
	} else {
		c.Ctx.Redirect(302, "/crud")
		return
	}

	// Si un formView est défini on utilisera son modèle pour les éléments
	formviewid := app.Tables[tableid].Views[viewid].FormView
	var elementsVF types.Elements
	if formviewid == "" {
		flash.Error("Article non trouvé")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}
	elementsVF = app.Tables[tableid].Forms[formviewid].Elements
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, _ := mergeElements(c.Controller, tableid, elementsVF, id)

	// lecture du record
	records, err := models.CrudRead(tableid, id, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	if len(records) == 0 {
		flash.Error("Article non trouvé")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}
	// Calcul des éléments
	elements = computeElements(c.Controller, false, elements, records[0])

	table := app.Tables[tableid]
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
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
		return
	}

	c.Ctx.Redirect(302, "/crud/view/"+appid+"/"+tableid+"/"+viewid+"/"+id)
}
