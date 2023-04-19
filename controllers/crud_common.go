package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"github.com/imdario/mergo"
)

var err error

// IsInGroup as
func IsInGroup(c beego.Controller, group string, appid string, id string) (out bool) {
	out = false
	session := getSession(c)

	// logs.Debug(fmt.Sprintf("%v", session))
	// logs.Debug(fmt.Sprintf("group: %s appid: %s id: %s", group, appid, id))

	if session.Username == "anonymous" {
		if session.AppID != appid {
			return
		}
	}
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
	groups := strings.Split(session.Groups, ",")
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
func mergeElements(c beego.Controller, appid string, tableid string, viewOrFormElements map[string]dico.Element, id string) (map[string]dico.Element, map[int]string) {
	table := dico.Ctx.Applications[appid].Tables[tableid]

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
		ele := dico.Ctx.Applications[appid].Tables[tableid].Elements[key]
		err := mergo.Merge(&element, &ele)
		if err != nil {
			logs.Error(err)
		} else {
			element.ID = key
			if element.Order == 0 {
				element.Order = order
			}
			order = element.Order
			cols[order] = key
			if dico.Ctx.Applications[appid].Tables[tableid].Setting.Key == key && id != "" {
				element.ReadOnly = true
			}
			if element.ComputeSqlite != "" {
				element.Protected = true
			}
			if element.PlaceHolder == "" {
				element.PlaceHolder = element.LabelLong
			}
			if element.Jointure.Column != "" {
				element.Jointure.Column = macro(c, appid, element.Jointure.Column, orm.Params{})
				element.ReadOnly = true
			}
			if element.Jointure.Join != "" {
				element.Jointure.Join = macro(c, appid, element.Jointure.Join, orm.Params{})
				element.ReadOnly = true
			}
			if sortID == key {
				element.SortDirection = sortDirection
			}
			// Attributs par défaut en fonction du type
			// action amount button card chart checkbox counter date email float image
			// list number pdf percent plugin radio tag tel text textarea time url
			switch element.Type {
			case "action":
				break
			case "amount":
				if element.Format == "" {
					element.Format = "amount"
				}
				if element.ColAlign == "" {
					element.ColAlign = "right"
				}
				if element.Class == "" {
					element.Class = "crud-cell-nowrap"
				}
				if element.Width == "" {
					element.Width = "s"
				}
				element.ColNoWrap = true
			case "button":
				break
			case "chart":
				if element.Width == "" {
					element.Width = "m"
				}
			case "checkbox":
				break
			case "counter":
				if element.ColAlign == "" {
					element.ColAlign = "center"
				}
				if element.Width == "" {
					element.Width = "s"
				}
				element.ColNoWrap = true
			case "date":
				if element.Format == "" {
					element.Format = "date"
				}
				if element.Width == "" {
					element.Width = "m"
				}
				element.ColNoWrap = true
			case "editor":
				if element.Class == "" {
					element.Class = "warning"
				}
			case "email":
				element.Width = "m"
			case "float":
				if element.ColAlign == "" {
					element.ColAlign = "right"
				}
				if element.Width == "" {
					element.Width = "s"
				}
				element.ColNoWrap = true
			case "image":
				if element.Width == "" {
					element.Width = "m"
				}
			case "list":
				if element.Width == "" {
					element.Width = "l"
				}
			case "number":
				if element.ColAlign == "" {
					element.ColAlign = "right"
				}
				if element.Width == "" {
					element.Width = "s"
				}
				element.ColNoWrap = true
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
				if element.Width == "" {
					element.Width = "s"
				}
				element.ColNoWrap = true
			case "pdf":
				break
			case "plugin":
				break
			case "radio":
				if element.Width == "" {
					element.Width = "l"
				}
			case "card":
				if !IsInGroup(c, table.Forms[element.Params.Form].Group, appid, id) {
					element.Params.Form = ""
				}
				if element.Width == "" {
					element.Width = "m"
				}
			case "tag":
				if element.Width == "" {
					element.Width = "m"
				}
			case "text":
				if element.Width == "" {
					element.Width = "m"
				}
			case "textarea":
				if element.Width == "" {
					element.Width = "l"
				}
			case "time":
				if element.Format == "" {
					element.Format = "time"
				}
				if element.Width == "" {
					element.Width = "s"
				}
				element.ColNoWrap = true
			case "url":
				if element.Width == "" {
					element.Width = "l"
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
	appid := c.Ctx.Input.Param(":app")
	tableid := c.Ctx.Input.Param(":table")
	table := dico.Ctx.Applications[appid].Tables[tableid]

	fromList := false
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		fromList = true
	}

	elements := make(map[string]dico.Element, len(viewOrFormElements))

	for key, element := range viewOrFormElements {
		element.LabelLong = macro(c, appid, element.LabelLong, record)
		// Help si non readOnly et non protected
		if element.Help != "" {
			if element.Protected || element.ReadOnly {
				element.Help = ""
			}
		}
		// Valorisation de Items ClassSqlite ItemsSQL, DefaultSqlite, HideSqlite
		if element.ClassSqlite != "" {
			element.Class = macroSQL(c, appid, element.ClassSqlite, record)
		}
		if element.HideSqlite != "" {
			if macroSQL(c, appid, element.HideSqlite, record) != "" {
				element.Hide = true
			}
		}
		if element.ItemsSQL != "" {
			sql := macro(c, appid, element.ItemsSQL, record)
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
				action.URL = macro(c, appid, action.URL, record)
			}
			if action.Plugin != "" {
				action.Plugin = macro(c, appid, action.Plugin, record)
			}
		}
		// for key, value := range element.Args {
		// 	element.Args[key] = macro(c, appid, value, record)
		// }
		if element.Params.URL != "" {
			if !fromList {
				element.Params.URL = macro(c, appid, element.Params.URL, record)
			}
		}
		if element.Params.Src != "" {
			if !fromList {
				element.Params.Src = macro(c, appid, element.Params.Src, record)
			}
		}

		if computeValue {
			val := ""
			if col, ok := record[key]; ok {
				if reflect.ValueOf(col).IsValid() {
					val = record[key].(string)
				}
			}
			if element.ComputeSqlite != "" {
				val = macroSQL(c, appid, element.ComputeSqlite, record)
			}
			if val == "" && element.DefaultSqlite != "" {
				val = macroSQL(c, appid, element.DefaultSqlite, record)
			}
			if val == "" && element.Default != "" {
				val = macro(c, appid, element.Default, record)
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
			case "number":
				if val == "" {
					val = "0"
				}
			case "percent":
				if val == "" {
					val = "0"
				}
			}
			// Update record avec valeur calculée
			if col, ok := record[key]; ok {
				if reflect.ValueOf(col).IsValid() {
					record[key] = val
				}
			}
		}
		if !IsInGroup(c, element.Group, appid, "") {
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

	valid := validation.Validation{}
	if !element.Hide && !element.Protected && !element.ReadOnly {
		if element.Required {
			valid.Required(val, element.LabelLong)
		}
		if element.MinLength > 0 {
			valid.MinSize(val, element.MinLength, element.LabelLong)
		}
		if element.MaxLength > 0 && len(val) > element.MaxLength {
			valid.MaxSize(val, element.MaxLength, element.LabelLong)
		}
		if element.Type == "email" {
			valid.Email(val, element.LabelLong)
		}
	}
	if !valid.HasErrors() {
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
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			labelError = fmt.Sprintf(
				"%s [%s]", err.Message, err.Key)
		}
		err = errors.New(labelError)
	}

	return err
}

// getSession
func getSession(c beego.Controller) types.Session {
	session := types.Session{}
	session.SessionID = c.Ctx.GetCookie("beegosessionID")
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
	if c.GetSession("AppId") != nil {
		session.AppID = c.GetSession("AppId").(string)
	}
	return session
}

// setContext remplissage du controller.Data
func setContext(c beego.Controller, appid string, tableid string) (sessionID string) {
	// Contexte de la table
	// aliasDB := dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB
	ctx := dico.Ctx
	aliasDB := ctx.Applications[appid].Tables[tableid].Setting.AliasDB
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
	c.Data["TableID"] = tableid
	c.Data["AliasDB"] = aliasDB
	// c.Data["KeyID"] = dico.Ctx.Applications[appid].Tables[tableid].Setting.Key
	c.Data["KeyID"] = ctx.Applications[appid].Tables[tableid].Setting.Key

	c.SetSession("AppId", appid)
	session := getSession(c)

	c.Data["Session"] = &session

	c.Data["Config"] = &models.Config

	// XSRF protection des formulaires
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf"] = c.XSRFToken()
	// Title
	c.Data["Title"] = models.Config.Appname
	c.Data["Portail"] = &dico.Ctx
	// Contexte crud
	// c.Data["From"] = c.Ctx.Input.Cookie("from")
	c.Data["Composter"] = time.Now().Unix()

	return session.SessionID

}

// macro qui remplace les {$user} et les {key}
func macro(c beego.Controller, appid string, in string, record orm.Params) (out string) {
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
					out = strings.ReplaceAll(out, "{"+key+"}", dico.Ctx.Applications[appid].Parameters[key])
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
func macroSQL(c beego.Controller, appid string, in string, record orm.Params) (out string) {
	out = ""
	sql := macro(c, appid, in, record)
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
func requeteSQL(c beego.Controller, appid string, in string, record orm.Params, aliasDB string) (out string) {
	out = ""
	sql := macro(c, appid, in, record)
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

// ctxNavigation enregistré dans la session de l'utilisteur
type ctxNavigation struct {
	Current int
	URL     map[int]string
}

// Set dans page portail /bee
func (nav *ctxNavigation) init() string {
	nav.Current = 0
	nav.URL = map[int]string{}
	nav.URL[nav.Current] = "/bee"
	return nav.URL[nav.Current]
}

// Appel nouvelle page
// from portail -> list
// from list -> view ou edit
// from view.list -> view ou edit
func (nav *ctxNavigation) forward(url string) string {
	if nav.getBackURL() == url {
		// forward suite à un backward
		return nav.backward()
	}
	if nav.URL[nav.Current] == url {
		// refresh
		return nav.URL[nav.Current]
	}
	nav.Current = nav.Current + 1
	nav.URL[nav.Current] = url
	return nav.URL[nav.Current]
}

// Bouton retour
func (nav *ctxNavigation) backward() string {
	if nav.Current > 0 {
		nav.Current = nav.Current - 1
	}
	return nav.URL[nav.Current]
}

// Donne l'URL du bouton retour arrière
func (nav *ctxNavigation) getBackURL() string {
	if nav.Current != 0 {
		return nav.URL[nav.Current-1]
	}
	return nav.URL[nav.Current]
}

func navigateInit(c beego.Controller) {
	if c.GetSession("navigateur") != nil {
		navigateur := c.GetSession("navigateur").(ctxNavigation)
		navigateur.init()
		c.SetSession("navigateur", navigateur)
		c.Data["From"] = navigateur.getBackURL()
	} else {
		navigateur := ctxNavigation{}
		navigateur.init()
		c.SetSession("navigateur", navigateur)
		c.Data["From"] = navigateur.getBackURL()
	}
	// logs.Info(fmt.Printf("\nNAVIGATEUR init %v\n", c.GetSession("navigateur").(ctxNavigation)))
}
func forward(c beego.Controller, url string) {
	if c.GetSession("navigateur") != nil {
		navigateur := c.GetSession("navigateur").(ctxNavigation)
		navigateur.forward(url)
		c.SetSession("navigateur", navigateur)
		c.Data["From"] = navigateur.getBackURL()
	} else {
		navigateur := ctxNavigation{}
		navigateur.init()
		c.SetSession("navigateur", navigateur)
		c.Data["From"] = navigateur.getBackURL()
	}
	// logs.Info(fmt.Printf("\nNAVIGATEUR forward %v\n", c.GetSession("navigateur").(ctxNavigation)))
}

// backward as
func backward(c beego.Controller) {
	if c.GetSession("navigateur") != nil {
		navigateur := c.GetSession("navigateur").(ctxNavigation)
		c.Ctx.Redirect(302, navigateur.backward())
		c.Data["From"] = navigateur.getBackURL()
	} else {
		c.Ctx.Redirect(302, "/bee")
		c.Data["From"] = "/bee"
	}
	// logs.Info(fmt.Printf("\nNAVIGATEUR backward %v\n", c.GetSession("navigateur").(ctxNavigation)))
}
