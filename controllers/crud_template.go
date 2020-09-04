package controllers

import (
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"
)

//
//	Fonctions pour les templates
//

// Déclaration des fonctions utilisées dans les templates
func init() {
	beego.AddFuncMap("CrudSplit", CrudSplit)
	beego.AddFuncMap("CrudContains", CrudContains)
	beego.AddFuncMap("CrudMacroSQL", CrudMacroSQL)
	beego.AddFuncMap("CrudFormat", CrudFormat)
	beego.AddFuncMap("CrudItem", CrudItem)
}

// CrudFormat préféré à text/template/printf car les données fournies sont toujours des strings
func CrudFormat(in string, value string) (out string) {
	out = ""
	if in != "" {
		recs, err := models.CrudSQL("SELECT printf('"+in+"','"+value+"')", "default")
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
			out = item.Value
		}
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
