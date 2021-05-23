package controllers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
)

// CrudListController as
type CrudListController struct {
	loggedRouter
}

// CrudList CrudListController
func (c *CrudListController) CrudList() {
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	viewid := c.Ctx.Input.Param(":view")

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

	// Contrôle d'accès à la vue
	table := dico.Ctx.Tables[tableid]
	view := dico.Ctx.Tables[tableid].Views[viewid]
	if view.Group == "" {
		view.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c.Controller, view.Group, "") {
		logs.Error("Accès non autorisé", c.GetSession("Username").(string), viewid, view.Group)
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

	setContext(c.Controller, tableid)

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
	elements, cols := mergeElements(c.Controller, tableid, dico.Ctx.Tables[tableid].Views[viewid].Elements, "")

	// Calcul des champs SQL de la vue
	if view.OrderBy != "" {
		view.OrderBy = macro(c.Controller, view.OrderBy, orm.Params{})
	}
	if view.FooterSQL != "" {
		view.FooterSQL = requeteSQL(c.Controller, view.FooterSQL, orm.Params{}, dico.Ctx.Tables[tableid].Setting.AliasDB)
	}
	if len(view.PreUpdateSQL) > 0 {
		for _, presql := range view.PreUpdateSQL {
			// Remplissage d'un record avec les elements.SQLout
			record := orm.Params{}
			sql := macro(c.Controller, presql, record)
			if sql != "" {
				err = models.CrudExec(sql, table.Setting.AliasDB)
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
		var colName string
		var val string
		var ope string
		re := regexp.MustCompile(`^(.*):(.*)`)
		match := re.FindStringSubmatch(search)
		if len(match) > 0 {
			colName = match[1]
			val = match[2]
			ope = "LIKE"
		}
		re = regexp.MustCompile(`^(.*)=(.*)`)
		match = re.FindStringSubmatch(search)
		if len(match) > 0 {
			colName = match[1]
			val = match[2]
			ope = "="
		}
		for key, element := range elements {
			if strings.HasPrefix(key, "_") {
				continue
			}
			if ope != "" && key == colName {
				// recherche sur une seule colonne
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
						if ope == "=" {
							view.Search += element.Jointure.Column + " = '" + val + "'"
						} else {
							view.Search += element.Jointure.Column + " = '" + val + "'"
						}
					} else {
						if ope == "=" {
							view.Search += tableid + "." + key + " = '" + val + "'"
						} else {
							view.Search += tableid + "." + key + " LIKE '%" + val + "%'"
						}
					}
				}
				break
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
	c.Data["Title"] = view.Title
	c.Data["AppId"] = appid
	c.Data["Application"] = dico.Ctx.Applications[appid]
	c.Data["TableId"] = tableid
	c.Data["ViewId"] = viewid
	c.Data["Table"] = &table
	c.Data["View"] = &view
	c.Data["Elements"] = elements
	c.Data["Records"] = records
	c.Data["Qrecords"] = len(records)
	c.Data["Cols"] = cols

	c.Ctx.Output.Cookie("from", fmt.Sprintf("/bee/list/%s/%s/%s", appid, tableid, viewid))

	if view.Type == "image" {
		c.TplName = "crud_list_image.html"
	} else if view.Type == "table" {
		c.TplName = "crud_list_table.html"
	} else {
		c.TplName = "crud_list_card.html"
	}
}
