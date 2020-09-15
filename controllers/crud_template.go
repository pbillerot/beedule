package controllers

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
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
	beego.AddFuncMap("CrudIndexSQL", CrudIndexSQL)
	beego.AddFuncMap("CrudMacro", CrudMacro)
	beego.AddFuncMap("CrudMacroSQL", CrudMacroSQL)
	beego.AddFuncMap("CrudSplit", CrudSplit)
}

// CrudIndex équivalent de index mais avec computeSQL en +
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
	if element.ComputeSQL != "" {
		out = CrudMacroSQL(element.ComputeSQL, record, session)
	} else {
		if val, ok := record[key]; ok {
			if reflect.ValueOf(val).IsValid() {
				out = val.(string)
			}
		}
	}
	return
}

// CrudFormat préféré à text/template/printf car les données fournies sont toujours des strings
func CrudFormat(in string, value string) (out string) {
	out = value
	if in != "" {
		recs, err := models.CrudSQL("SELECT printf('"+in+"','"+value+"')", "default")
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
				out = val.(string)
			}
		}
	}
	return
}
