package controllers

import (
	"fmt"
	"strings"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// CrudHugoController as
type CrudHugoController struct {
	loggedRouter
}

// CrudHugo CrudHugoController
func (c *CrudHugoController) CrudHugo() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")

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

	// Contrôle d'accès à la vue
	table := app.Tables[tableid]
	view := app.Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = app.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, "") {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), viewid, view.Group)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		ReturnFrom(c.Controller)
		return
	}
	// Ctrl d'accès FormAdd FormView FormEdit
	if !IsInGroup(c.Controller, table.Forms[view.FormView].Group, "") {
		view.FormView = ""
	}
	if !IsInGroup(c.Controller, table.Forms[view.FormAdd].Group, "") {
		view.FormAdd = ""
	}
	if !IsInGroup(c.Controller, table.Forms[view.FormEdit].Group, "") {
		view.FormEdit = ""
	}

	// Gestion du TRI enregistré dans la session et contexte
	sortID := c.GetString("sortid")
	sortDirection := c.GetString("sortdirection")
	ctxSortid := fmt.Sprintf("%s-%s-%s-sortid", appid, tableid, viewid)
	ctxSortdirection := fmt.Sprintf("%s-%s-%s-sortdirection", appid, tableid, viewid)
	if sortID != "" {
		c.SetSession(ctxSortid, sortID)
		c.SetSession(ctxSortdirection, sortDirection)
	} else {
		if c.GetSession(ctxSortid) != nil {
			sortID = c.GetSession(ctxSortid).(string)
		}
		if c.GetSession(ctxSortdirection) != nil {
			sortDirection = c.GetSession(ctxSortdirection).(string)
		}
	}
	// Data récupéré dans mergeElements et dans le template ensuite
	c.Data["SortID"] = sortID
	c.Data["SortDirection"] = sortDirection

	// Fusion des attributs des éléments de la table dans les éléments de la vue
	elements, cols := mergeElements(c.Controller, tableid, app.Tables[tableid].Views[viewid].Elements, "")

	// Calcul des champs SQL de la vue
	if view.OrderBy != "" {
		view.OrderBy = macro(c.Controller, view.OrderBy, orm.Params{})
	}
	if view.FooterSQL != "" {
		view.FooterSQL = requeteSQL(c.Controller, view.FooterSQL, orm.Params{}, app.Tables[tableid].AliasDB)
	}
	if len(view.PreUpdateSQL) > 0 {
		for _, presql := range view.PreUpdateSQL {
			// Remplissage d'un record avec les elements.SQLout
			record := orm.Params{}
			sql := macro(c.Controller, presql, record)
			if sql != "" {
				err = models.CrudExec(sql, table.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c.Controller)
				}
			}
		}
	}
	if view.Where != "" {
		view.Where = macro(c.Controller, view.Where, orm.Params{})
	}

	// RECHERCHE DANS LA VUE
	search := strings.ToLower(c.GetString("search"))
	ctxSearch := fmt.Sprintf("%s-%s-%s-search", appid, tableid, viewid)
	if strings.ToLower(c.GetString("searchstop")) != "" {
		c.DelSession(ctxSearch)
		search = ""
	}
	if search != "" {
		c.SetSession(ctxSearch, search)
	} else {
		if c.GetSession(ctxSearch) != nil {
			search = c.GetSession(ctxSearch).(string)
		}
	}

	if search != "" {
		for key, element := range elements {
			if strings.HasPrefix(key, "_") {
				continue
			}
			switch element.Type {
			case "checkbox":
				if strings.Contains(strings.ToLower(element.LabelShort), search) {
					if view.Search != "" {
						view.Search += " OR "
					}
					if element.Jointure.Column != "" {
						view.Search += element.Jointure.Column + " = '1'"
					} else {
						view.Search += tableid + "." + key + " = '1'"
					}
				}
			case "combobox":
				// TODO recherche dans le label du combobox
			default:
				if view.Search != "" {
					view.Search += " OR "
				}
				if element.Jointure.Column != "" {
					view.Search += element.Jointure.Column + " LIKE '%" + search + "%'"
				} else {
					view.Search += tableid + "." + key + " LIKE '%" + search + "%'"
				}
			}
		}
	}
	// Filtrage si élément owner
	for key, element := range elements {
		// Un seule élément owner par enregistrement
		if element.Group == "owner" && !IsAdmin(c.Controller) {
			if view.Search != "" {
				view.Search = "(" + view.Search + ") AND "
			}
			view.Search += tableid + "." + key + " = '" + c.GetSession("Username").(string) + "'"
			break
		}
	}
	c.Data["Search"] = search

	// lecture des records
	records, err := models.CrudList(tableid, viewid, &view, elements)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	if len(records) > 0 {
		// Calcul des éléments hors values
		elements = computeElements(c.Controller, false, elements, records[0])
	}

	// Remplissage du contexte pour le template
	setContext(c.Controller)
	c.Data["Title"] = view.Title
	c.Data["AppId"] = appid
	c.Data["Application"] = app.Applications[appid]
	c.Data["TableId"] = tableid
	c.Data["ViewId"] = viewid
	c.Data["Table"] = &table
	c.Data["View"] = &view
	c.Data["Elements"] = elements
	c.Data["Records"] = records
	c.Data["Qrecords"] = len(records)
	c.Data["Cols"] = cols

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/crud/list/%s/%s/%s", appid, tableid, viewid))

	if view.Type == "image" {
		c.TplName = "crud_list_image.html"
	} else if view.Type == "table" {
		c.TplName = "crud_table.html"
	} else {
		c.TplName = "crud_list.html"
	}
}
