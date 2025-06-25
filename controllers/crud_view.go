package controllers

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"

	beego "github.com/beego/beego/v2/adapter"
)

// CrudViewController as
type CrudViewController struct {
	loggedRouter
}

// Get CrudViewController appel du formulaire formView
func (c *CrudViewController) Get() {
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

	// Si un formView est défini on utilisera son modèle pour les éléments
	formviewid := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid].FormView
	formview := dico.Ctx.Applications[appid].Tables[tableid].Forms[formviewid]
	// Ctrl accès à formviewid
	if formview.Group == "" {
		formview.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, formview.Group, appid, id) {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), formviewid, formview.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		// ReturnFrom(c.Controller)
		c.Ctx.Redirect(302, "/bee")
		return
	}
	// Ctrl d'accès FormAdd FormView FormEdit
	if !IsInGroup(c.Controller, table.Forms[view.FormView].Group, appid, id) {
		view.FormView = ""
	}
	if !IsInGroup(c.Controller, table.Forms[view.FormAdd].Group, appid, id) {
		view.FormAdd = ""
	}
	if !IsInGroup(c.Controller, table.Forms[view.FormEdit].Group, appid, id) {
		view.FormEdit = ""
	}

	setContext(c.Controller, appid, tableid)

	var elementsVF map[string]dico.Element
	if formviewid == "" {
		elementsVF = dico.Ctx.Applications[appid].Tables[tableid].Views[viewid].Elements
	} else {
		elementsVF = dico.Ctx.Applications[appid].Tables[tableid].Forms[formviewid].Elements
	}
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, cols := mergeElements(c.Controller, appid, tableid, elementsVF, id)

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
	}
	// Calcul des éléments
	elements = computeElements(c.Controller, true, elements, records[0])
	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé: %v", id)
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	// Cas rubrique type card avec params view table
	uiViews := map[string]*UIView{}
	var uiView UIView
	for _, element := range elements {
		if element.Type == "card" {
			if element.Params.View != "" {
				// Appel constructeur de la vue
				if element.Params.Where != "" {
					element.Params.Where = macro(c.Controller, appid, element.Params.Where, element.Record)
				}
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
	forward(c.Controller, fmt.Sprintf("/bee/view/%s/%s/%s/%s", appid, tableid, viewid, id))

	// Titre du formulaire
	form := dico.Ctx.Applications[appid].Tables[tableid].Forms[formviewid]
	form.Title = macro(c.Controller, appid, form.Title, records[0])

	if err == nil {
		c.Data["ColDisplay"] = records[0][table.Setting.ColDisplay].(string)
	} else {
		c.Data["ColDisplay"] = "Enregistrement non trouvé"
	}
	c.Data["AppId"] = appid
	c.Data["Application"] = dico.Ctx.Applications[appid]
	c.Data["Id"] = id
	c.Data["TableId"] = tableid
	c.Data["ViewId"] = viewid
	c.Data["FormView"] = &form
	c.Data["FormViewId"] = formviewid
	c.Data["FormId"] = dico.Ctx.Applications[appid].Tables[tableid].Views[viewid].FormEdit
	c.Data["Table"] = &table
	c.Data["View"] = &view
	c.Data["Elements"] = &elements
	c.Data["Records"] = &records
	c.Data["Cols"] = &cols
	c.Data["UIViews"] = &uiViews
	c.Data["UIView"] = &uiView

	c.TplName = "crud_view.html"
}
