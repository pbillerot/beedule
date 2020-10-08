package controllers

import (
	"fmt"
	"strings"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"

	"github.com/astaxie/beego"
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
	if !IsInGroup(c.Controller, view.Group, id) {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), viewid, view.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	// Si un formView est défini on utilisera son modèle pour les éléments
	formviewid := app.Tables[tableid].Views[viewid].FormView
	formview := app.Tables[tableid].Views[formviewid]
	// Ctrl accès à formviewid
	if formview.Group == "" {
		formview.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, formview.Group, id) {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), formviewid, formview.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	// Ctrl d'accès FormAdd FormView FormEdit
	if !IsInGroup(c.Controller, table.Forms[view.FormView].Group, id) {
		view.FormView = ""
	}
	if !IsInGroup(c.Controller, table.Forms[view.FormAdd].Group, id) {
		view.FormAdd = ""
	}
	if !IsInGroup(c.Controller, table.Forms[view.FormEdit].Group, id) {
		view.FormEdit = ""
	}

	var elementsVF types.Elements
	if formviewid == "" {
		elementsVF = app.Tables[tableid].Views[viewid].Elements
	} else {
		elementsVF = app.Tables[tableid].Forms[formviewid].Elements
	}
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, cols := mergeElements(c.Controller, tableid, elementsVF, id)

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
	// Calcul des éléments
	elements = computeElements(c.Controller, false, elements, records[0])

	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé: ", id)
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}

	c.SetSession(fmt.Sprintf("anch_%s_%s", tableid, viewid), fmt.Sprintf("anch_%s", strings.ReplaceAll(id, ".", "_")))
	c.Ctx.Output.Cookie("from", fmt.Sprintf("/crud/view/%s/%s/%s/%s", appid, tableid, viewid, id))

	setContext(c.Controller)
	if err == nil {
		c.Data["ColDisplay"] = records[0][table.ColDisplay].(string)
	} else {
		c.Data["ColDisplay"] = "Enregistrement non trouvé"
	}
	c.Data["AppId"] = appid
	c.Data["Application"] = app.Applications[appid]
	c.Data["Id"] = id
	c.Data["TableId"] = tableid
	c.Data["ViewId"] = viewid
	c.Data["FormView"] = app.Tables[tableid].Forms[formviewid]
	c.Data["FormId"] = app.Tables[tableid].Views[viewid].FormEdit
	c.Data["Table"] = &table
	c.Data["View"] = &view
	c.Data["Elements"] = &elements
	c.Data["Records"] = &records
	c.Data["Cols"] = &cols

	c.TplName = "crud_view.html"
}
