package controllers

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"

	beego "github.com/beego/beego/v2/adapter"
)

// CrudViewController as
type CrudDashboardController struct {
	loggedRouter
}

// Get CrudViewController appel du formulaire formView
func (c *CrudDashboardController) Get() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	id := ""

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
		logs.Error("Table not found", c.GetSession("Username").(string), tableid)
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
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), viewid, view.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		// ReturnFrom(c.Controller)
		c.Ctx.Redirect(302, "/bee")
		return
	}

	// Si une vue formView est défini on utilisera son modèle pour les éléments
	formview := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid]
	// Ctrl accès à formviewid
	if formview.Group == "" {
		formview.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, formview.Group, appid, id) {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), viewid, formview.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		// ReturnFrom(c.Controller)
		c.Ctx.Redirect(302, "/bee")
		return
	}

	setContext(c.Controller, appid, tableid)

	elementsVF := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid].Elements
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, cols := mergeElements(c.Controller, appid, tableid, elementsVF, id)

	// Calcul des éléments
	record := orm.Params{}
	for k, element := range elements {
		record[k] = element.Default
	}
	var records = []orm.Params{}
	records = append(records, record)
	elements = computeElements(c.Controller, true, elements, records[0])

	// Cas rubrique type card avec params view table
	uiViews := map[string]*UIView{}
	for _, element := range elements {
		if element.Type == "card" {
			if element.Params.View != "" {
				// Appel constructeur de la vue
				if element.Params.Where != "" {
					element.Params.Where = macro(c.Controller, appid, element.Params.Where, element.Record)
				}
				var uiView UIView
				err = uiView.load(c.Controller, appid, element.Params.Table, element.Params.View, element)
				if err != nil {
					backward(c.Controller)
				} else {
					// ATTENTION le nom des vues dans un formulaire doivent être unique
					uiViews[element.Params.View] = &uiView
				}
			}
		}
	}

	c.SetSession(fmt.Sprintf("anch_%s_%s", tableid, viewid), fmt.Sprintf("anch_%s", strings.ReplaceAll(id, ".", "_")))

	// Positionnement du navigateur sur la page qui va s'ouvrir
	forward(c.Controller, fmt.Sprintf("/bee/dashboard/%s/%s/%s", appid, tableid, viewid))

	c.Data["AppId"] = appid
	c.Data["Application"] = dico.Ctx.Applications[appid]
	c.Data["Id"] = id
	c.Data["TableId"] = tableid
	c.Data["Table"] = &table
	c.Data["ViewId"] = viewid
	c.Data["View"] = &view
	c.Data["Elements"] = &elements
	c.Data["Records"] = &records
	c.Data["Cols"] = &cols
	c.Data["UIViews"] = &uiViews
	// c.Data["UIView"] = &uiView

	c.TplName = "crud_dashboard.html"
}
