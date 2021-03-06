package controllers

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
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
	beego.AddFuncMap("CrudArgs", CrudArgs)
	beego.AddFuncMap("CrudDecrement", CrudDecrement)
	beego.AddFuncMap("CrudDebug", CrudDebug)
	beego.AddFuncMap("BeeReplace", BeeReplace)
	beego.AddFuncMap("dict", Dictionary)
	beego.AddFuncMap("CrudComputeDataset", CrudComputeDataset)
}

// CrudArgs as ?key=val&key=val
func CrudArgs(in map[string]string) (out string) {
	out = ""
	for key, val := range in {
		if out == "" {
			out = "?"
		} else {
			out += "&"
		}
		out += fmt.Sprintf("%s=%s", key, val)
	}
	return
}

// BeeReplace as
func BeeReplace(in string, old string, new string) (out string) {
	out = strings.Replace(in, old, new, 1)
	return
}

// CrudDebug as
func CrudDebug(msg string) (out string) {
	beego.Debug(msg)
	return
}

// CrudIncrement as
func CrudIncrement(in int) int {
	in++
	return in
}

// CrudDecrement as
func CrudDecrement(in int) int {
	in--
	return in
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
func CrudIndexSQL(record orm.Params, key string, element dico.Element, session types.Session) (out string) {
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
			logs.Error(err)
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
func CrudItem(items []dico.Item, in string) (out string) {
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
					if p, ok := dico.Ctx.Parameters[key]; ok {
						out = strings.ReplaceAll(out, "{"+key+"}", p)
					} else {
						labelError := fmt.Sprintf("Rubrique [%s] non trouvée", key)
						logs.Error(labelError)
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
					logs.Error(labelError)
					err = errors.New(labelError)
					break
				}
			} else {
				labelError := fmt.Sprintf("Syntaxe incorrecte [%s]", out)
				logs.Error(labelError)
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
			logs.Error(err)
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
			logs.Error(err)
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
func CrudClassSQL(element dico.Element, record orm.Params, session types.Session) (out string) {
	// out = CrudMacro(element.Class, record, session)
	// if out == "" {
	sql := CrudMacro(element.ClassSQL, record, session)
	if sql != "" {
		recs, err := models.CrudSQL(sql, "default")
		if err != nil {
			logs.Error(err)
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

// CrudComputeDataset retourne le dataset valorisé
func CrudComputeDataset(dataset map[string]string, record orm.Params, session types.Session, aliasDB string) map[string]string {
	out := make(map[string]string, len(dataset))
	for key, value := range dataset {
		sql := CrudMacro(value, record, session)
		recs, err := models.CrudSQL(sql, aliasDB)
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
		out[key] = val
	}
	return out
}

// Dictionary creates a map[string]interface{} from the given parameters by
// walking the parameters and treating them as key-value pairs.  The number
// of parameters must be even.
// The keys can be string slices, which will create the needed nested structure.
// Clone dict du projet https://github.com/gohugoio/hugo/blob/master/tpl/collections/collections.go
func Dictionary(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dictionary call")
	}

	root := make(map[string]interface{})

	for i := 0; i < len(values); i += 2 {
		dict := root
		var key string
		switch v := values[i].(type) {
		case string:
			key = v
		case []string:
			for i := 0; i < len(v)-1; i++ {
				key = v[i]
				var m map[string]interface{}
				v, found := dict[key]
				if found {
					m = v.(map[string]interface{})
				} else {
					m = make(map[string]interface{})
					dict[key] = m
				}
				dict = m
			}
			key = v[len(v)-1]
		default:
			return nil, errors.New("invalid dictionary key")
		}
		dict[key] = values[i+1]
	}

	return root, nil
}
