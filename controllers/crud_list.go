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

	var uiView UIView
	err = uiView.load(c.Controller, appid, tableid, viewid, dico.Element{})
	if err != nil {
		backward(c.Controller)
		return
	}

	// Remplissage du contexte pour le template
	setContext(c.Controller, uiView.TableID)

	c.Data["AppId"] = uiView.AppID
	c.Data["Application"] = dico.Ctx.Applications[uiView.AppID]
	c.Data["UIView"] = &uiView

	// Positionnement du navigateur sur la page qui va s'ouvrir
	forward(c.Controller, fmt.Sprintf("/bee/list/%s/%s/%s", appid, tableid, viewid))

	if uiView.View.Type == "image" {
		c.TplName = "crud_list_image.html"
	} else if uiView.View.Type == "table" {
		c.TplName = "crud_list_table.html"
	} else {
		c.TplName = "crud_list_card.html"
	}
}

// UIView Vue
type UIView struct {
	Title         string
	AppID         string
	TableID       string
	ViewID        string
	Table         dico.Table
	View          dico.View
	Elements      map[string]dico.Element
	Records       []orm.Params
	Qrecords      int
	Cols          map[int]string
	SortID        string
	SortDirection string
	Search        string
}

func (ui *UIView) load(c beego.Controller, appid string, tableid string, viewid string, parentElement dico.Element) (err error) {
	ui.AppID = appid
	ui.TableID = tableid
	ui.ViewID = viewid

	flash := beego.ReadFromRequest(&c)

	// Ctrl appid tableid viewid formid
	if _, ok := dico.Ctx.Applications[appid]; !ok {
		err = fmt.Errorf("App not found %s", appid)
		logs.Error(err.Error())
		flash.Error(err.Error())
		flash.Store(&c)
		return
	}
	if val, ok := dico.Ctx.Tables[tableid]; ok {
		if _, ok := val.Views[viewid]; ok {
		} else {
			err = fmt.Errorf("View not found %s", viewid)
			logs.Error(err.Error())
			flash.Error(err.Error())
			flash.Store(&c)
			return
		}
	} else {
		err = fmt.Errorf("Table not found %s", tableid)
		logs.Error(err.Error())
		flash.Error(err.Error())
		flash.Store(&c)
		return
	}

	// Contrôle d'accès à la vue
	table := dico.Ctx.Tables[tableid]
	view := dico.Ctx.Tables[tableid].Views[viewid]
	ui.Table = *table
	ui.View = view
	ui.Title = view.Title
	if view.Group == "" {
		view.Group = dico.Ctx.Applications[appid].Group
	}
	if !IsInGroup(c, view.Group, "") {
		err = fmt.Errorf("Accès non autorisé de %s à %s", viewid, view.Group)
		logs.Error(err.Error())
		flash.Error(err.Error())
		flash.Store(&c)
		return
	}
	// Ctrl d'accès FormAdd FormView FormEdit
	if !IsInGroup(c, table.Forms[view.FormView].Group, "") {
		view.FormView = ""
	}
	if !IsInGroup(c, table.Forms[view.FormAdd].Group, "") {
		view.FormAdd = ""
	}
	if !IsInGroup(c, table.Forms[view.FormEdit].Group, "") {
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
	ui.SortID = sortID
	ui.SortDirection = sortDirection
	// Fusion des attributs des éléments de la table dans les éléments de la vue
	// en intégrant les éléments de tri fournis dans c.Data
	c.Data["SortID"] = sortID
	c.Data["SortDirection"] = sortDirection
	elements, cols := mergeElements(c, tableid, dico.Ctx.Tables[tableid].Views[viewid].Elements, "")
	ui.Cols = cols
	// Calcul des champs SQL de la vue
	if view.OrderBy != "" {
		view.OrderBy = macro(c, view.OrderBy, orm.Params{})
	}
	if view.FooterSQL != "" {
		view.FooterSQL = requeteSQL(c, view.FooterSQL, orm.Params{}, dico.Ctx.Tables[tableid].Setting.AliasDB)
	}
	if len(view.PreUpdateSQL) > 0 {
		for _, presql := range view.PreUpdateSQL {
			// Remplissage d'un record avec les elements.SQLout
			record := orm.Params{}
			sql := macro(c, presql, record)
			if sql != "" {
				err = models.CrudExec(sql, table.Setting.AliasDB)
				if err != nil {
					flash.Error(err.Error())
					flash.Store(&c)
				}
			}
		}
	}
	// CAS appel d'une vue dans le formulaire
	if parentElement.Params.View != "" {
		if parentElement.Params.Where != "" {
			view.Where = macro(c, parentElement.Params.Where, parentElement.Record)
		} else {
			if view.Where != "" {
				view.Where = macro(c, view.Where, orm.Params{})
			}
		}
		if parentElement.LabelLong != "" {
			view.Title = parentElement.LabelLong
		}
		if parentElement.Params.IconName != "" {
			view.IconName = parentElement.Params.IconName
		}
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
		if element.Group == "owner" && !IsAdmin(c) {
			if view.Search != "" {
				view.Search = "(" + view.Search + ") AND "
			}
			view.Search += tableid + "." + key + " = '" + c.GetSession("Username").(string) + "'"
			break
		}
	}
	ui.Search = search

	// lecture des records
	records, err := models.CrudList(tableid, viewid, &view, elements)
	ui.Records = records
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c)
	}
	if len(records) > 0 {
		// Calcul des éléments hors values
		elements = computeElements(c, false, elements, records[0])
		ui.Qrecords = len(records)
	}
	ui.Elements = elements

	return nil
}
