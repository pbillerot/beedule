package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/imdario/mergo"
)

var err error

// mergeElements fusionne les attributs des éléments de la table avec ceux de la vue ou formulaire
// cols contiendra les keys ordonnés comme présentées dans le dictionnaire
func mergeElements(c beego.Controller, tableid string, viewOrFormElements types.Elements, id string) (types.Elements, map[int]string) {

	elements := types.Elements{}

	cols := make(map[int]string, len(elements))

	order := 1 // pour ordrer les colonnes
	for key, element := range viewOrFormElements {
		err := mergo.Merge(&element, app.Tables[tableid].Elements[key])
		if err != nil {
			beego.Error(err)
		} else {
			if element.Order == 0 {
				element.Order = order
			}
			order = element.Order
			cols[order] = key
			if app.Tables[tableid].Key == key && id != "" {
				element.ReadOnly = true
			}
			if element.ComputeSQL != "" {
				element.Protected = true
			}
			if element.PlaceHolder == "" {
				element.PlaceHolder = element.LabelLong
			}
			if element.Jointure.Column != "" {
				element.Jointure.Column = macro(c, element.Jointure.Column, orm.Params{})
			}
			if element.Jointure.Join != "" {
				element.Jointure.Join = macro(c, element.Jointure.Join, orm.Params{})
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
			case "textarea":
				if element.Class == "" {
					element.Class = "warning"
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

// computeElements calcule les éléments
func computeElements(c beego.Controller, computeValue bool, viewOrFormElements types.Elements, record orm.Params) types.Elements {
	tableid := c.Ctx.Input.Param(":table")
	table := app.Tables[tableid]

	elements := types.Elements{}
	for key, element := range viewOrFormElements {
		// Valorisation de Items ClassSQL ItemsSQL, ComputeSQL, DefaultSQL, HideSQL
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
			recs, err := models.CrudSQL(sql, table.AliasDB)
			if err != nil {
				beego.Error(err)
			}
			var list []types.Item
			for _, rec := range recs {
				// item := types.Item{Key: rec["key"].(string), Label: rec["label"].(string)}
				item := types.Item{}
				// la 1ère colonne représente la clé
				// la 2ème le label à afficher
				for _, label := range rec {
					if item.Key == "" {
						item.Key = label.(string)
					} else {
						item.Label = label.(string)
					}
				}
				list = append(list, item)
			}
			element.Items = list
		}
		if element.Action.URL != "" {
			element.Action.URL = macro(c, element.Action.URL, record)
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
			// Update record avec valeur calculée
			if col, ok := record[key]; ok {
				if reflect.ValueOf(col).IsValid() {
					record[key] = val
				}
			}
		}
		elements[key] = element
	}
	return elements
}

// computeElements calcule les éléments
func defaultElements(c beego.Controller, computeValue bool, viewOrFormElements types.Elements, record orm.Params) types.Elements {
	tableid := c.Ctx.Input.Param(":table")
	table := app.Tables[tableid]

	elements := types.Elements{}
	for key, element := range viewOrFormElements {
		// Valorisation de Items ClassSQL ItemsSQL, ComputeSQL, DefaultSQL, HideSQL
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
			recs, err := models.CrudSQL(sql, table.AliasDB)
			if err != nil {
				beego.Error(err)
			}
			var list []types.Item
			for _, rec := range recs {
				// item := types.Item{Key: rec["key"].(string), Label: rec["label"].(string)}
				item := types.Item{}
				// la 1ère colonne représente la clé
				// la 2ème le label à afficher
				for _, label := range rec {
					if item.Key == "" {
						item.Key = label.(string)
					} else {
						item.Label = label.(string)
					}
				}
				list = append(list, item)
			}
			element.Items = list
		}
		if element.Action.URL != "" {
			element.Action.URL = macro(c, element.Action.URL, record)
		}

		if computeValue {
			if element.ComputeSQL != "" {
				element.Value = macroSQL(c, element.ComputeSQL, record)
			}
			if element.Value == "" && element.DefaultSQL != "" {
				element.Value = macroSQL(c, element.DefaultSQL, record)
			}
			if element.Value == "" && element.Default != "" {
				element.Value = macro(c, element.Default, record)
			}
		}
		elements[key] = element
	}
	return elements
}

// checkElement:
// - recup de la saisie dans element.Value
// - contrôle de la saisie
// - valorisation de element.SQLout pour l'enregistrement dans la bdd
func checkElement(c *beego.Controller, key string, element *types.Element, record orm.Params) error {
	labelError := ""

	// Récupération de la saisie
	switch element.Type {
	case "tag":
		element.Value = strings.Join(c.GetStrings(key)[:], ",")
	default:
		element.Value = c.GetString(key)
	}

	if element.Required && element.Value == "" {
		labelError = fmt.Sprintf(
			"[%s] est obligatoire", element.LabelLong)
	}
	if element.MinLength > 0 && len(element.Value) < element.MinLength {
		labelError = fmt.Sprintf(
			"%d caractères minimum pour [%s]", element.MinLength, element.LabelLong)
	}
	if element.MaxLength > 0 && len(element.Value) > element.MaxLength {
		labelError = fmt.Sprintf(
			"%d caractères maximum pour [%s]", element.MinLength, element.LabelLong)
	}

	if labelError == "" {
		// Valorisation de SQLout pour l'enregistrement
		switch element.Type {
		case "checkbox":
			if element.Value == "" {
				element.SQLout = "0"
			} else {
				// le mot de passe a été changé
				element.SQLout = "1"
			}
		case "password":
			if element.Value == "***" {
				element.SQLout = record[key].(string)
			} else {
				// le mot de passe a été changé
				element.SQLout = element.HashPassword()
			}
		default:
			element.SQLout = element.Value
		}
	}
	var err error
	if labelError != "" {
		err = errors.New(labelError)
	}

	return err
}

// setContext remplissage du controller.Data
func setContext(c beego.Controller) {
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
	c.Data["Session"] = &session

	config := types.Config{}
	config.Appname = beego.AppConfig.String("appname")
	config.Appnote = beego.AppConfig.String("appnote")
	config.Icone = beego.AppConfig.String("icone")
	config.Site = beego.AppConfig.String("site")
	config.Email = beego.AppConfig.String("email")
	config.Author = beego.AppConfig.String("author")
	config.Version = beego.AppConfig.String("version")
	config.Theme = beego.AppConfig.String("theme")
	c.Data["Config"] = &config

	// XSRF protection des formulaires
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	// Title
	c.Data["Title"] = config.Appname
	c.Data["Portail"] = &app.Portail
	// Contexte bee
	u, _ := url.Parse(c.Ctx.Request.RequestURI)
	c.Data["beePath"] = u.Path
	c.Data["beeReturn"] = c.Ctx.Request.Referer()
	c.Data["Composter"] = time.Now().Unix()
	if c.GetSession("from") != nil {
		beego.Debug("*** FROM:", c.GetSession("from"))
		c.Data["From"] = c.GetSession("from").(string)
	}
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
	re := regexp.MustCompile(`.*{(.*)}.*`)
	for strings.Contains(out, "{") {
		match := re.FindStringSubmatch(in)
		if len(match) > 0 {
			key := match[1]
			if val, ok := record[key]; ok {
				if reflect.ValueOf(val).IsValid() {
					out = strings.ReplaceAll(out, "{"+key+"}", val.(string))
				} else {
					beego.Error("Colonne NULL", key)
					out = ""
				}
			} else {
				if strings.Contains(key, "__") {
					// Le champ est un paramètre global
					out = strings.ReplaceAll(out, "{"+key+"}", app.Params[key])
				} else {
					beego.Error("Rubrique non trouvée", key)
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
			beego.Error(err, in)
		}
		for _, record := range ress {
			for _, val := range record {
				if reflect.ValueOf(val).IsValid() {
					out = val.(string)
				} else {
					beego.Error(in)
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
		beego.Error(err, in)
	}
	for _, record := range ress {
		for _, val := range record {
			if reflect.ValueOf(val).IsValid() {
				out = val.(string)
			} else {
				beego.Error(in)
				out = ""
			}
		}
	}
	return
}
