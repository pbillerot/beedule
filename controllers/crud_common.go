package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"github.com/imdario/mergo"
)

var err error

// ReturnFrom as
func ReturnFrom(c beego.Controller) {
	if c.Data["From"] != "" {
		c.Ctx.Redirect(302, c.Data["From"].(string))
	} else {
		c.Ctx.Redirect(302, "/bee")
	}
	return
}

// IsInGroup as
func IsInGroup(c beego.Controller, group string, id string) (out bool) {
	out = false
	if group == "" {
		out = true
		return
	}
	if IsAdmin(c) {
		out = true
		return
	}
	if group == "owner" && IsOwner(c, id) {
		out = true
		return
	}
	sgroups := c.GetSession("Groups").(string)
	groups := strings.Split(sgroups, ",")
	for _, g := range groups {
		if g == group {
			out = true
			return
		}
	}
	return
}

// IsAdmin as
func IsAdmin(c beego.Controller) (out bool) {
	out = c.GetSession("IsAdmin").(bool)
	return
}

// IsOwner as
func IsOwner(c beego.Controller, key string) (out bool) {
	if c.GetSession("Username").(string) == key {
		out = true
	} else {
		out = false
	}
	return
}

// mergeElements fusionne les attributs des éléments de la table avec ceux de la vue ou formulaire
// cols contiendra les keys ordonnés comme présentées dans le dictionnaire
func mergeElements(c beego.Controller, tableid string, viewOrFormElements map[string]dico.Element, id string) (map[string]dico.Element, map[int]string) {
	table := dico.Ctx.Tables[tableid]

	// Gestion du TRI dans la session
	sortID := ""
	sortDirection := ""
	if v, ok := c.Data["SortID"]; ok {
		sortID = v.(string)
	}
	if v, ok := c.Data["SortDirection"]; ok {
		sortDirection = v.(string)
	}

	elements := make(map[string]dico.Element, len(viewOrFormElements))

	cols := make(map[int]string, len(elements))

	order := 1 // pour ordrer les colonnes
	for key, element := range viewOrFormElements {
		err := mergo.Merge(&element, dico.Ctx.Tables[tableid].Elements[key])
		if err != nil {
			logs.Error(err)
		} else {
			if element.Order == 0 {
				element.Order = order
			}
			order = element.Order
			cols[order] = key
			if dico.Ctx.Tables[tableid].Setting.Key == key && id != "" {
				element.ReadOnly = true
			}
			if element.ComputeSQL != "" {
				element.Protected = true
			}
			if element.PlaceHolder == "" {
				element.PlaceHolder = element.LabelLong
			}
			if element.Params.Column != "" {
				element.Params.Column = macro(c, element.Params.Column, orm.Params{})
			}
			if element.Params.Join != "" {
				element.Params.Join = macro(c, element.Params.Join, orm.Params{})
			}
			if sortID == key {
				element.SortDirection = sortDirection
			}
			// Attributs par défaut en fonction du type
			switch element.Type {
			case "amount":
				if element.Format == "" {
					element.Format = "%3.2f €"
				}
				if element.ColAlign == "" {
					element.ColAlign = "right"
				}
				if element.Class == "" {
					element.Class = "crud-cell-nowrap"
				}
			case "count":
				if element.ColAlign == "" {
					element.ColAlign = "center"
				}
			case "editor":
				if element.Class == "" {
					element.Class = "warning"
				}
			case "date":
				if element.Format == "" {
					element.Format = "date"
				}
			case "datetime":
				if element.Format == "" {
					element.Format = "datetime"
				}
			case "float":
				if element.ColAlign == "" {
					element.ColAlign = "right"
				}
			case "month":
				if element.ColAlign == "" {
					element.ColAlign = "center"
				}
			case "number":
				if element.ColAlign == "" {
					element.ColAlign = "right"
				}
			case "percent":
				if element.Format == "" {
					element.Format = "%3.2f %%"
				}
				if element.ColAlign == "" {
					element.ColAlign = "right"
				}
				if element.Class == "" {
					element.Class = "crud-cell-nowrap"
				}
			case "section":
				if !IsInGroup(c, table.Forms[element.Params.Form].Group, id) {
					element.Params.Form = ""
				}
			case "textarea":
				if element.Class == "" {
					element.Class = "warning"
				}
			case "time":
				if element.Format == "" {
					element.Format = "time"
				}
			case "week":
				if element.ColAlign == "" {
					element.ColAlign = "center"
				}
			}
			elements[key] = element
		}
		order++
	}
	return elements, cols
}

// computeElements calcule les éléments de l'UI
// si computeValue, valorise à 0 les champ numériques dans record
func computeElements(c beego.Controller, computeValue bool, viewOrFormElements map[string]dico.Element, record orm.Params) map[string]dico.Element {
	tableid := c.Ctx.Input.Param(":table")
	table := dico.Ctx.Tables[tableid]

	fromList := false
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		fromList = true
	}

	elements := make(map[string]dico.Element, len(viewOrFormElements))

	for key, element := range viewOrFormElements {
		element.LabelLong = macro(c, element.LabelLong, record)
		// Valorisation de Items ClassSQL ItemsSQL, DefaultSQL, HideSQL
		if element.ClassSQL != "" {
			element.Class = macroSQL(c, element.ClassSQL, record)
		}
		if element.HideSQL != "" {
			if macroSQL(c, element.HideSQL, record) != "" {
				element.Hide = true
			}
		}
		if element.ItemsSQL != "" {
			sql := macro(c, element.ItemsSQL, record)
			recs, err := models.CrudSQL(sql, table.Setting.AliasDB)
			if err != nil {
				logs.Error(err)
			}
			var list []dico.Item
			for _, rec := range recs {
				item := dico.Item{Key: rec["key"].(string), Label: rec["label"].(string)}
				list = append(list, item)
			}
			element.Items = list
		}
		for _, action := range element.Actions {
			if action.URL != "" {
				action.URL = macro(c, action.URL, record)
			}
			if action.Plugin != "" {
				action.Plugin = macro(c, action.Plugin, record)
			}
		}
		if element.Params.URL != "" {
			if !fromList {
				element.Params.URL = macro(c, element.Params.URL, record)
			}
		}
		if element.Params.Src != "" {
			if !fromList {
				element.Params.Src = macro(c, element.Params.Src, record)
			}
		}

		if computeValue {
			val := ""
			if col, ok := record[key]; ok {
				if reflect.ValueOf(col).IsValid() {
					val = record[key].(string)
				}
			}
			if element.ComputeSQL != "" {
				val = macroSQL(c, element.ComputeSQL, record)
			}
			if val == "" && element.DefaultSQL != "" {
				val = macroSQL(c, element.DefaultSQL, record)
			}
			if val == "" && element.Default != "" {
				val = macro(c, element.Default, record)
			}
			// Valorisation avec les args de l'url
			if c.GetStrings(key) != nil {
				val = c.GetString(key)
			}
			// Valeur par défaut
			switch element.Type {
			case "amount":
				if val == "" {
					val = "0"
				}
			case "float":
				if val == "" {
					val = "0"
				}
			case "month":
				if val == "" {
					val = "0"
				}
			case "number":
				if element.Dataset != nil {
					for key, value := range element.Dataset {
						sql := macro(c, value, record)
						recs, err := models.CrudSQL(sql, table.Setting.AliasDB)
						if err != nil {
							logs.Error(err)
						}
						val := ""
						bstart := true
						for _, cols := range recs {
							// tri des keys
							keys := make([]string, 0, len(cols))
							for k := range cols {
								keys = append(keys, k)
							}
							sort.Strings(keys)
							for _, k := range keys {
								if bstart {
									bstart = false
								} else {
									val += ","
								}
								val += cols[k].(string)

							}
						}
						element.Dataset[key] = val
					}
				}

				if val == "" {
					val = "0"
				}
			case "percent":
				if val == "" {
					val = "0"
				}
			}
			if element.Dataset != nil {
				for key, value := range element.Dataset {
					sql := macro(c, value, record)
					recs, err := models.CrudSQL(sql, table.Setting.AliasDB)
					if err != nil {
						logs.Error(err)
					}
					val := ""
					bstart := true
					for _, cols := range recs {
						// tri des keys
						keys := make([]string, 0, len(cols))
						for k := range cols {
							keys = append(keys, k)
						}
						sort.Strings(keys)
						for _, k := range keys {
							if bstart {
								bstart = false
							} else {
								val += ","
							}
							val += cols[k].(string)

						}
					}
					element.Dataset[key] = val
				}
			}
			// Update record avec valeur calculée
			if col, ok := record[key]; ok {
				if reflect.ValueOf(col).IsValid() {
					record[key] = val
				}
			}
		}
		if key == "_action_sell" {
			elements[key] = element
		}
		if !IsInGroup(c, element.Group, "") {
			element.Hide = true
		}
		element.Record = record
		elements[key] = element
	}

	return elements
}

// checkElement:
// - recup de la saisie dans val
// - contrôle de la saisie
// - valorisation de element.SQLout pour l'enregistrement dans la bdd
func checkElement(c *beego.Controller, key string, element *dico.Element, record orm.Params) error {
	labelError := ""
	val := ""

	// Récupération de la saisie
	switch element.Type {
	case "tag":
		val = strings.Join(c.GetStrings(key)[:], ",")
	default:
		val = c.GetString(key)
	}
	record[key] = val

	if element.Required && val == "" {
		labelError = fmt.Sprintf(
			"[%s] est obligatoire", element.LabelLong)
	}
	if element.MinLength > 0 && len(val) < element.MinLength {
		labelError = fmt.Sprintf(
			"%d caractères minimum pour [%s]", element.MinLength, element.LabelLong)
	}
	if element.MaxLength > 0 && len(val) > element.MaxLength {
		labelError = fmt.Sprintf(
			"%d caractères maximum pour [%s]", element.MinLength, element.LabelLong)
	}

	if labelError == "" {
		// Valorisation de SQLout pour l'enregistrement
		switch element.Type {
		case "amount":
			if val == "" {
				element.SQLout = "0"
			} else {
				element.SQLout = val
			}
		case "float":
			if val == "" {
				element.SQLout = "0"
			} else {
				element.SQLout = val
			}
		case "month":
			if val == "" {
				element.SQLout = "0"
			} else {
				element.SQLout = val
			}
		case "number":
			if val == "" {
				element.SQLout = "0"
			} else {
				element.SQLout = val
			}
		case "percent":
			if val == "" {
				element.SQLout = "0"
			} else {
				element.SQLout = val
			}
		case "checkbox":
			if val == "" {
				element.SQLout = "0"
			} else {
				// le mot de passe a été changé
				element.SQLout = "1"
			}
		case "password":
			if val == "***" {
				element.SQLout = record[key].(string)
			} else {
				// le mot de passe a été changé
				element.SQLout = element.HashPassword(val)
			}
		default:
			element.SQLout = val
		}
	}
	var err error
	if labelError != "" {
		err = errors.New(labelError)
	}

	return err
}

// setContext remplissage du controller.Data
func setContext(c beego.Controller, table string) {
	// Contexte de la table
	aliasDB := dico.Ctx.Tables[table].Setting.AliasDB
	section, _ := beego.AppConfig.GetSection(aliasDB)
	dataurl := "/bee/data/" + aliasDB
	if url, ok := section["dataurl"]; ok {
		dataurl = url
	}
	datadir := "./data/" + aliasDB
	if dir, ok := section["datadir"]; ok {
		datadir = dir
	}
	c.Data["DataUrl"] = dataurl
	c.Data["Datadir"] = datadir
	c.Data["TableID"] = table
	c.Data["AliasDB"] = aliasDB
	c.Data["KeyID"] = dico.Ctx.Tables[table].Setting.Key

	// Contexte de la session
	session := types.Session{}
	if c.GetSession("LoggedIn") != nil {
		session.LoggedIn = c.GetSession("LoggedIn").(bool)
	}
	if c.GetSession("Username") != nil {
		session.Username = c.GetSession("Username").(string)
	}
	if c.GetSession("IsAdmin") != nil {
		session.IsAdmin = c.GetSession("IsAdmin").(bool)
	}
	if c.GetSession("Groups") != nil {
		session.Groups = c.GetSession("Groups").(string)
	}
	c.Data["Session"] = &session

	c.Data["Config"] = &models.Config

	// XSRF protection des formulaires
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	// Title
	c.Data["Title"] = models.Config.Appname
	c.Data["Portail"] = &dico.Ctx
	// Contexte crud
	// c.Data["From"] = c.Ctx.Input.Cookie("from")
	c.Data["Composter"] = time.Now().Unix()
}

func setSession(c beego.Controller, param string, value string) {
	c.SetSession(param, value)
}

// macro qui remplace les {$user} et les {key}
func macro(c beego.Controller, in string, record orm.Params) (out string) {
	out = in

	if strings.Contains(out, "{$user}") {
		out = strings.ReplaceAll(out, "{$user}", c.GetSession("Username").(string))
	}
	if strings.Contains(out, "{$datadir}") {
		out = strings.ReplaceAll(out, "{$datadir}", c.Data["Datadir"].(string))
	}
	if strings.Contains(out, "{$dataurl}") {
		out = strings.ReplaceAll(out, "{$dataurl}", c.Data["DataUrl"].(string))
	}
	if strings.Contains(out, "{$table}") {
		out = strings.ReplaceAll(out, "{$table}", c.Data["TableID"].(string))
	}
	if strings.Contains(out, "{$aliasdb}") {
		out = strings.ReplaceAll(out, "{$aliasdb}", c.Data["AliasDB"].(string))
	}
	if strings.Contains(out, "{$key}") {
		out = strings.ReplaceAll(out, "{$key}", c.Data["KeyID"].(string))
	}
	re := regexp.MustCompile(`.*{(.*)}.*`)
	for strings.Contains(out, "{") {
		match := re.FindStringSubmatch(out)
		if len(match) > 0 {
			key := match[1]
			if val, ok := record[key]; ok {
				if reflect.ValueOf(val).IsValid() {
					// beego.Debug("avant", key, val.(string), out)
					out = strings.ReplaceAll(out, "{"+key+"}", val.(string))
					// beego.Debug("apres", out)
				} else {
					logs.Error("Colonne NULL", key)
					out = ""
				}
			} else {
				if strings.Contains(key, "__") {
					// Le champ est un paramètre global
					out = strings.ReplaceAll(out, "{"+key+"}", dico.Ctx.Parameters[key])
				} else {
					logs.Error("Rubrique non trouvée", key)
					out = ""
				}
			}
		}
	}

	return
}

// macroSQL qui remplace les {$user} et les {key} et exécute le résultat en SQL sur l'alias default
func macroSQL(c beego.Controller, in string, record orm.Params) (out string) {
	out = ""
	sql := macro(c, in, record)
	if sql != "" {
		ress, err := models.CrudSQL(sql, "default")
		if err != nil {
			logs.Error(err, in)
		}
		for _, record := range ress {
			for _, val := range record {
				if reflect.ValueOf(val).IsValid() {
					out = val.(string)
				} else {
					logs.Error(sql)
					out = ""
				}
			}
		}
	}
	return
}

// requeteSQL qui remplace les {$user} et les {key} et exécute le résultat en SQL
func requeteSQL(c beego.Controller, in string, record orm.Params, aliasDB string) (out string) {
	out = ""
	sql := macro(c, in, record)
	ress, err := models.CrudSQL(sql, aliasDB)
	if err != nil {
		logs.Error(err, in)
	}
	for _, record := range ress {
		for _, val := range record {
			if reflect.ValueOf(val).IsValid() {
				out = val.(string)
			} else {
				logs.Error(in)
				out = ""
			}
		}
	}
	return
}
