package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
)

// CrudAddController as
type CrudAddController struct {
	loggedRouter
}

// Get CrudAddController
func (c *CrudAddController) Get() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	var id = ""

	flash := beego.ReadFromRequest(&c.Controller)
	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		logs.Error("App not found", c.GetSession("Username").(string), appid)
		backward(c.Controller)
		return
	}
	if val, ok := dico.Ctx.Tables[tableid]; ok {
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
		logs.Error("Table not found", c.GetSession("Username").(string), tableid)
		backward(c.Controller)
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
	if !IsInGroup(c.Controller, form.Group, id) {
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		backward(c.Controller)
		return
	}

	setContext(c.Controller, tableid)

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, dico.Ctx.Tables[tableid].Forms[formid].Elements, id)

	// Création d'un record fictif vide ""
	record := orm.Params{}
	for _, colname := range cols {
		if c.GetStrings(colname) != nil {
			record[colname] = c.GetString(colname)
		} else {
			record[colname] = ""
		}
	}
	var records []orm.Params
	records = append(records, record)

	// Calcul des éléments
	elements = computeElements(c.Controller, true, elements, records[0])

	c.Data["AppId"] = appid
	c.Data["Application"] = dico.Ctx.Applications[appid]
	c.Data["ColDisplay"] = records[0][table.Setting.ColDisplay]
	c.Data["Id"] = id
	c.Data["TableId"] = tableid
	c.Data["ViewId"] = viewid
	c.Data["FormId"] = formid
	c.Data["Table"] = &table
	c.Data["View"] = &view
	c.Data["Form"] = &form
	c.Data["Elements"] = elements
	c.Data["Records"] = records
	c.Data["Cols"] = cols

	c.TplName = "crud_edit.html"
}

// Post CrudAddController
func (c *CrudAddController) Post() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")
	formid := c.Ctx.Input.Param(":form")
	var id = ""

	// Ctrl tableid et viewid
	if val, ok := dico.Ctx.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
			if _, ok := val.Forms[formid]; ok {
			} else {
				backward(c.Controller)
				return
			}
		} else {
			backward(c.Controller)
			return
		}
	} else {
		backward(c.Controller)
		return
	}

	flash := beego.ReadFromRequest(&c.Controller)

	table := dico.Ctx.Tables[tableid]
	view := dico.Ctx.Tables[tableid].Views[viewid]
	form := dico.Ctx.Tables[tableid].Forms[formid]
	setContext(c.Controller, tableid)

	// Fusion des attributs des éléments de la table dans les éléments du formulaire
	elements, cols := mergeElements(c.Controller, tableid, dico.Ctx.Tables[tableid].Forms[formid].Elements, id)

	// Création d'un record fictif vide ""
	record := orm.Params{}
	for _, colname := range cols {
		record[colname] = ""
	}
	var records []orm.Params
	records = append(records, record)

	// Lecture, contrôle des champs saisis
	// et remplissage de SQLOut pour l'enregistrement
	berr := false
	for key, element := range elements {
		err = checkElement(&c.Controller, key, &element, records[0])
		if err != nil {
			berr = true
			element.Error = "error"
			flash.Error(err.Error())
		}
		elements[key] = element
		// On écrase les données lues dans la table par la saisie
		// Ceci est nécessaire pour représenter les valeurs saisies suite à uen erreur
		records[0][key] = element.SQLout
	}
	if berr { // ERREUR: on va reproposer le formulaire pour rectification
		flash.Store(&c.Controller)

		// Calcul des éléments (valeur par défaut comprise)
		elements = computeElements(c.Controller, true, elements, records[0])

		c.Data["AppId"] = appid
		c.Data["Application"] = dico.Ctx.Applications[appid]
		c.Data["ColDisplay"] = records[0][table.Setting.ColDisplay]
		c.Data["Id"] = id
		c.Data["TableId"] = tableid
		c.Data["ViewId"] = viewid
		c.Data["FormId"] = formid
		c.Data["Table"] = &table
		c.Data["View"] = &view
		c.Data["Form"] = &form
		c.Data["Elements"] = elements
		c.Data["Records"] = records
		c.Data["Cols"] = cols

		c.TplName = "crud_edit.html"
		return
	}
	// C'est OK, les données sont correctes et placées dans SQLout
	err = models.CrudInsert(tableid, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Data["error"] = "error"
		backward(c.Controller)
		return
	}
	// PostSQL
	for _, postsql := range form.PostSQL {
		sql := macro(c.Controller, postsql, records[0])
		if sql != "" {
			err = models.CrudExec(sql, table.Setting.AliasDB)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
			}
		} else {
			logs.Error("Ordre sql incorrect ", postsql)
			flash.Error("Ordre sql incorrect ", postsql)
			flash.Store(&c.Controller)
			backward(c.Controller)
			return
		}
	}

	flash.Notice("Création effectuée avec succès")
	flash.Store(&c.Controller)

	backward(c.Controller)
}
