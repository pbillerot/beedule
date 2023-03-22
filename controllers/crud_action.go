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
	if !IsInGroup(c.Controller, view.Group, appid, actionid) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	setContext(c.Controller, appid, tableid)

	iactionid, err := strconv.Atoi(actionid)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}
	if iactionid <= len(view.Actions) {
		// Exécution des ordres SQL
		for _, action := range view.Actions[iactionid].SQL {
			sql := macro(c.Controller, appid, action, orm.Params{})
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
		backward(c.Controller)
		return
	}
	if val, ok := dico.Ctx.Applications[appid].Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				logs.Error("Form not found", c.GetSession("Username").(string), formid)
				backward(c.Controller)
				return
			}
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
	form := dico.Ctx.Applications[appid].Tables[tableid].Forms[formid]
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, appid, actionid) {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	setContext(c.Controller, appid, tableid)
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
			sql := macro(c.Controller, appid, action, orm.Params{})
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
	if !IsInGroup(c.Controller, view.Group, appid, actionid) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	setContext(c.Controller, appid, tableid)
	var withPlugin bool

	// Si un formView est défini on utilisera son modèle pour les éléments
	form := dico.Ctx.Applications[appid].Tables[tableid].Forms[formid]
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, appid, id) {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}
	// var elementsVF map[string]dico.Element
	var elementsVF = form.Elements
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, _ := mergeElements(c.Controller, appid, tableid, elementsVF, id)

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
					sql := macro(c.Controller, appid, actionSQL, records[0])
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

// CrudActionAjaxController as
type CrudActionAjaxController struct {
	loggedRouter
}

// CrudActionAjaxController as
type CrudAjaxSqlController struct {
	loggedRouter
}

type ajaxSqlResponse struct {
	Response string // ok or ko
	Message  string
	Dataset  map[string]string // données résulat du sql
}

// Post CrudAjaxSqlController
func (c *CrudAjaxSqlController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	actionid := c.Ctx.Input.Param(":action") // l'id de l'élément

	var rest ajaxSqlResponse
	rest.Response = "Error"
	rest.Message = "Ya un problème"
	rest.Dataset = make(map[string]string)

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		logs.Error("App not found", c.GetSession("Username").(string), appid)
		rest.Message = "App not found"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}
	if val, ok := dico.Ctx.Applications[appid].Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			logs.Error("View not found", c.GetSession("Username").(string), viewid)
			rest.Message = "View not found"
			c.Data["json"] = &rest
			c.ServeJSON()
			return
		}
	} else {
		logs.Error("Table not found", c.GetSession("Username").(string), tableid)
		flash.Error("Table not found")
		flash.Store(&c.Controller)
		rest.Message = "Table not found"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}

	// Contrôle d'accès
	table := dico.Ctx.Applications[appid].Tables[tableid]
	view := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, appid, actionid) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		rest.Message = "Accès non autorisé"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}

	setContext(c.Controller, appid, tableid)

	// Si un formView est défini on utilisera son modèle pour les éléments
	form := dico.Ctx.Applications[appid].Tables[tableid].Forms[formid]
	if form.Group == "" {
		form.Group = view.Group
	}
	if form.Group == "" {
		form.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, form.Group, appid, "") {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), formid, form.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}
	// var elementsVF map[string]dico.Element
	var elementsVF = form.Elements
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, _ := mergeElements(c.Controller, appid, tableid, elementsVF, "")

	// Calcul des éléments
	elements = computeElements(c.Controller, true, elements, orm.Params{})

	if element, ok := elements[actionid]; ok {
		// Exécution du ajaj-sql les paramètres sont dans le POST
		record := orm.Params{}
		for key := range element.Dataset {
			record[key] = c.GetString(key) // donnée de la request
		}
		sql := macro(c.Controller, appid, element.AjaxSQL, record)
		if sql != "" {
			records, err := models.CrudSQL(sql, table.Setting.AliasDB)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
				rest.Response = "Error"
				rest.Message = err.Error()
				c.Data["json"] = &rest
				c.ServeJSON()
				return
			} else {
				// remplissage de la response
				for _, rec := range records {
					for key, val := range rec {
						rest.Dataset[key] = val.(string)
					}
				}
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
		rest.Message = "Action non trouvée"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}

	rest.Response = "ok"
	rest.Message = "ça roule"
	c.Data["json"] = &rest
	c.ServeJSON()

}

type ajaxResponse struct {
	Response string // ok or ko
	Message  string
}

// Post CrudActionElementController
func (c *CrudActionAjaxController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	id := c.Ctx.Input.Param(":id")
	actionid := c.Ctx.Input.Param(":action") // l'id de l'élément

	var rest ajaxResponse
	rest.Response = "Error"
	rest.Message = "Ya un problème"

	flash := beego.ReadFromRequest(&c.Controller)

	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		logs.Error("App not found", c.GetSession("Username").(string), appid)
		rest.Message = "App not found"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}
	if val, ok := dico.Ctx.Applications[appid].Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			logs.Error("View not found", c.GetSession("Username").(string), viewid)
			rest.Message = "View not found"
			c.Data["json"] = &rest
			c.ServeJSON()
			return
		}
	} else {
		logs.Error("Table not found", c.GetSession("Username").(string), tableid)
		flash.Error("Table not found")
		flash.Store(&c.Controller)
		rest.Message = "Table not found"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}

	// Contrôle d'accès
	table := dico.Ctx.Applications[appid].Tables[tableid]
	view := dico.Ctx.Applications[appid].Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, appid, actionid) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		rest.Message = "Accès non autorisé"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}

	setContext(c.Controller, appid, tableid)

	// var elementsVF map[string]dico.Element
	var elementsVF = view.Elements
	// Fusion des attributs des éléments de la table dans les éléments de la vue ou formulaire
	elements, _ := mergeElements(c.Controller, appid, tableid, elementsVF, id)

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
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		rest.Message = err.Error()
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}
	if len(records) == 0 {
		flash.Error("Enregistrement non trouvé")
		flash.Store(&c.Controller)
		rest.Message = "Enregistrement non trouvé"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}
	// Calcul des éléments
	elements = computeElements(c.Controller, true, elements, records[0])

	if element, ok := elements[actionid]; ok {
		// Exécution des ordres SQL
		for _, actionSQL := range element.Params.SQL {
			sql := macro(c.Controller, appid, actionSQL, records[0])
			if sql != "" {
				err = models.CrudExec(sql, table.Setting.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c.Controller)
					rest.Response = "Error"
					rest.Message = err.Error()
					c.Data["json"] = &rest
					c.ServeJSON()
					return
				}
			}
		}
	} else {
		flash.Error("Action non trouvée")
		flash.Store(&c.Controller)
		rest.Message = "Action non trouvée"
		c.Data["json"] = &rest
		c.ServeJSON()
		return
	}

	rest.Response = "ok"
	rest.Message = "ça roule"
	c.Data["json"] = &rest
	c.ServeJSON()

}
