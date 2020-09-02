package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"net/url"
	"regexp"
	"strings"

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
			if element.PlaceHolder == "" {
				element.PlaceHolder = element.LabelLong
			}
			// Valorisation de Items ClassSQL ItemsSQL, ComputeSQL, DefaultSQL
			if element.ClassSQL != "" {
				element.Class = macroSQL(c, element.ClassSQL, orm.Params{}, app.Tables[tableid].AliasDB)
			}
			if element.ComputeSQL != "" {
				element.ComputeSQL = macroSQL(c, element.ComputeSQL, orm.Params{}, app.Tables[tableid].AliasDB)
			}
			if element.DefaultSQL != "" {
				element.DefaultSQL = macroSQL(c, element.DefaultSQL, orm.Params{}, app.Tables[tableid].AliasDB)
			}
			if element.ItemsSQL != "" {
				sql := macro(c, element.ItemsSQL, orm.Params{})
				records, err := models.CrudSQL(sql, app.Tables[tableid].AliasDB)
				if err != nil {
					beego.Error(err)
				}
				var list string
				for _, record := range records {
					for _, item := range record {
						if list != "" {
							list += ","
						}
						list += item.(string)
					}
				}
				element.Items = list
			}

			if element.Value == "" && element.Default != "" {
				element.Value = macro(c, element.Default, orm.Params{})
			}
			elements[key] = element
		}
		order++
	}
	return elements, cols
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
			out = strings.ReplaceAll(out, "{"+key+"}", record[key].(string))
		}
	}

	return
}

// macroSQL qui remplace les {$user} et les {key} et exécute le résultat en SQL
func macroSQL(c beego.Controller, in string, record orm.Params, aliasDB string) (out string) {
	out = ""
	sql := macro(c, in, record)
	ress, err := models.CrudSQL(sql, aliasDB)
	if err != nil {
		beego.Error(err)
	}
	for _, record := range ress {
		for _, val := range record {
			out = val.(string)
		}
	}

	return
}

// Déclaration des fonctions utilisées dans les templates
func init() {
	beego.AddFuncMap("CrudSplit", CrudSplit)
	beego.AddFuncMap("CrudContains", CrudContains)
	beego.AddFuncMap("CrudMacroSQL", CrudMacroSQL)
}

// CrudSplit CrudSplit strings séparées par une virgule en slice
func CrudSplit(in string, separateur string) (out []string) {
	if in != "" {
		out = strings.Split(in, separateur)
	} else {
		out = []string{}
	}
	return
}

// CrudContains as
// list : "item1,item2,..."
// in   : "item2"
// ret  : valeur à retourner si OK
func CrudContains(list string, in string, ret string) (out string) {
	if strings.Contains(list, in) {
		out = ret
	} else {
		out = ""
	}
	return
}

// CrudMacroSQL retourne le résulat de la requête avec macro
// in: formule SQLite = select 'grey' where '{task_status}' = 'Terminée'
func CrudMacroSQL(in string, record orm.Params, aliasDB string) (out string) {
	out = ""
	if in != "" {
		sql := in
		re := regexp.MustCompile(`.*{(.*)}.*`)
		for strings.Contains(sql, "{") {
			match := re.FindStringSubmatch(sql)
			if len(match) > 0 {
				key := match[1]
				sql = strings.ReplaceAll(sql, "{"+key+"}", record[key].(string))
			}
		}
		if sql != "" {
			recs, err := models.CrudSQL(sql, aliasDB)
			if err != nil {
				beego.Error(err)
			}
			for _, rec := range recs {
				for _, val := range rec {
					out = val.(string)
				}
			}
		}
	}
	return
}

// setContext remplissage du controller.Data
func setContext(c beego.Controller) {
	// Données de session dans le contextx
	type Session struct {
		LoggedIn bool
		Username string
		IsAdmin  bool
	}
	session := Session{}
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

	// Paramètres de config dans le contexte
	type Config struct {
		Appname string
		Appnote string
		Icone   string
		Site    string
		Email   string
		Author  string
		Version string
		Theme   string
	}
	config := Config{}
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
}

func setSession(c beego.Controller, param string, value string) {
	c.SetSession(param, value)
}
