package controllers

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"
)

//
//	Fonctions pour les templates
//

// Déclaration des fonctions utilisées dans les templates
func init() {
	beego.AddFuncMap("CrudContains", CrudContains)
	beego.AddFuncMap("CrudFormat", CrudFormat)
	beego.AddFuncMap("CrudItem", CrudItem)
	beego.AddFuncMap("CrudIndex", CrudIndex)
	beego.AddFuncMap("CrudIndexAnchor", CrudIndexAnchor)
	beego.AddFuncMap("CrudIndexSQL", CrudIndexSQL)
	beego.AddFuncMap("CrudIsInGroup", CrudIsInGroup)
	beego.AddFuncMap("CrudMacro", CrudMacro)
	beego.AddFuncMap("CrudMacroSQL", CrudMacroSQL)
	beego.AddFuncMap("CrudClassSQL", CrudClassSQL)
	beego.AddFuncMap("CrudSplit", CrudSplit)
	beego.AddFuncMap("CrudSQL", CrudSQL)
	beego.AddFuncMap("CrudIncrement", CrudIncrement)
	beego.AddFuncMap("CrudDecrement", CrudDecrement)
	beego.AddFuncMap("CrudDebug", CrudDebug)
	beego.AddFuncMap("HugoIncrement", HugoIncrement)
	beego.AddFuncMap("HugoDecrement", HugoDecrement)
}

// CrudDebug as
func CrudDebug(msg string) (out string) {
	beego.Debug(msg)
	return
}

// HugoIncrement as
func HugoIncrement(in int) (out int) {
	in++
	out = in
	return
}

// HugoDecrement as
func HugoDecrement(in int) (out int) {
	in--
	out = in
	return
}

// CrudIncrement as
func CrudIncrement(snum string) (out string) {
	in, _ := strconv.Atoi(snum)
	in++
	out = strconv.Itoa(in)
	return
}

// CrudDecrement as
func CrudDecrement(snum string) (out string) {
	in, _ := strconv.Atoi(snum)
	in--
	out = strconv.Itoa(in)
	return
}

// CrudIsInGroup as
func CrudIsInGroup(group string, session types.Session) (out bool) {
	out = false
	if group == "" {
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

// CrudIndexAnchor Calcul de l'ancre à partir de la clé
func CrudIndexAnchor(record orm.Params, key string) (out string) {
	out = ""
	if val, ok := record[key]; ok {
		if reflect.ValueOf(val).IsValid() {
			out = "anch_" + strings.ReplaceAll(val.(string), ".", "_")
		}
	}
	return
}

// CrudIndex équivalent de index mais sans valeur nulle
func CrudIndex(record orm.Params, key string) (out string) {
	out = ""
	if val, ok := record[key]; ok {
		if reflect.ValueOf(val).IsValid() {
			out = val.(string)
		}
	}
	return
}

// CrudIndexSQL équivalent de index mais avec computeSQL en +
func CrudIndexSQL(record orm.Params, key string, element types.Element, session types.Session) (out string) {
	out = ""
	// if element.ComputeSQL != "" {
	// 	out = CrudMacroSQL(element.ComputeSQL, record, session)
	// } else {
	if val, ok := record[key]; ok {
		if reflect.ValueOf(val).IsValid() {
			out = val.(string)
		}
	}
	// }
	return
}

// CrudFormat préféré à text/template/printf car les données fournies sont toujours des strings
func CrudFormat(in string, value string) (out string) {
	out = value
	if in != "" && value != "" {
		var recs []orm.Params
		var err error
		if in == "datetime" {
			if strings.Contains(value, "0001-01") {
				return ""
			}
			recs, err = models.CrudSQL("SELECT strftime('%Y-%m-%d %H:%M:%S','"+value+"')", "default")
		} else if in == "date" {
			if strings.Contains(value, "0001-01") {
				return ""
			}
			recs, err = models.CrudSQL("SELECT strftime('%Y-%m-%d','"+value+"')", "default")
		} else if in == "time" {
			if strings.Contains(value, "00-00-00") {
				return ""
			}
			recs, err = models.CrudSQL("SELECT strftime('%H:%M:%S','"+value+"')", "default")
		} else {
			recs, err = models.CrudSQL("SELECT printf('"+in+"','"+value+"')", "default")
		}
		if err != nil {
			beego.Error(err)
		} else {
			for _, rec := range recs {
				for _, val := range rec {
					if reflect.ValueOf(val).IsValid() {
						out = val.(string)
					}
				}
			}
		}
	}
	return
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
func CrudContains(list string, in string) (out bool) {
	if strings.Contains(list, in) {
		out = true
	} else {
		out = false
	}
	return
}

// CrudItem as
func CrudItem(items []types.Item, in string) (out string) {
	for _, item := range items {
		if item.Key == in {
			out = item.Label
		}
	}
	return
}

// CrudMacro remplace les macros
func CrudMacro(in string, record orm.Params, session types.Session) (out string) {
	out = in
	if in != "" {
		if strings.Contains(out, "{$user}") {
			out = strings.ReplaceAll(out, "{$user}", session.Username)
		}
		re := regexp.MustCompile(`.*{(.*)}.*`)
		for strings.Contains(out, "{") {
			match := re.FindStringSubmatch(out)
			if len(match) > 0 {
				key := match[1]
				if strings.Contains(key, "__") {
					// Le champ est un paramètre global
					if p, ok := app.Params[key]; ok {
						out = strings.ReplaceAll(out, "{"+key+"}", p)
					} else {
						labelError := fmt.Sprintf("Rubrique [%s] non trouvée", key)
						beego.Error(labelError)
						err = errors.New(labelError)
						break
					}
					continue
				}
				if val, ok := record[key]; ok {
					if reflect.ValueOf(val).IsValid() {
						out = strings.ReplaceAll(out, "{"+key+"}", val.(string))
					} else {
						out = strings.ReplaceAll(out, "{"+key+"}", "")
					}
				} else {
					labelError := fmt.Sprintf("Rubrique [%s] non trouvée", key)
					beego.Error(labelError)
					err = errors.New(labelError)
					break
				}
			} else {
				labelError := fmt.Sprintf("Syntaxe incorrecte [%s]", out)
				beego.Error(labelError)
				err = errors.New(labelError)
				break
			}
		}
	}
	return
}

// CrudMacroSQL retourne le résulat de la requête avec macro
// in: formule SQLite = select 'grey' where '{task_status}' = 'Terminée'
func CrudMacroSQL(in string, record orm.Params, session types.Session) (out string) {
	out = ""
	sql := CrudMacro(in, record, session)
	if sql != "" {
		recs, err := models.CrudSQL(sql, "default")
		if err != nil {
			beego.Error(err)
		}
		for _, rec := range recs {
			for _, val := range rec {
				if reflect.ValueOf(val).IsValid() {
					out = val.(string)
				}
			}
		}
	}
	return
}

// CrudSQL retourne le résulat de la requête
func CrudSQL(sql string, aliasDB string) (out string) {
	out = ""
	if sql != "" {
		recs, err := models.CrudSQL(sql, aliasDB)
		if err != nil {
			beego.Error(err)
		}
		for _, rec := range recs {
			for _, val := range rec {
				if reflect.ValueOf(val).IsValid() {
					out = val.(string)
				}
			}
		}
	}
	return
}

// CrudClassSQL retourne le résulat de la requête avec macro
func CrudClassSQL(element types.Element, record orm.Params, session types.Session) (out string) {
	// out = CrudMacro(element.Class, record, session)
	// if out == "" {
	sql := CrudMacro(element.ClassSQL, record, session)
	if sql != "" {
		recs, err := models.CrudSQL(sql, "default")
		if err != nil {
			beego.Error(err)
		}
		for _, rec := range recs {
			for _, val := range rec {
				if reflect.ValueOf(val).IsValid() {
					out = val.(string)
				}
			}
		}
	}
	// }
	return
}
