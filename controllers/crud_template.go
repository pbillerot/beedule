package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/leekchan/accounting"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"
	"github.com/pbillerot/beedule/types"
)

//
//	Fonctions pour les templates
// https://gist.github.com/mholt/2a874dd67e830d07ec11
//

// Déclaration des fonctions utilisées dans les templates
func init() {
	beego.AddFuncMap("contains", CrudContains)
	beego.AddFuncMap("CrudContains", CrudContains)
	beego.AddFuncMap("CrudFormat", CrudFormat)
	beego.AddFuncMap("CrudItem", CrudItem)
	beego.AddFuncMap("CrudIndex", CrudIndex)
	beego.AddFuncMap("CrudIndexAnchor", CrudIndexAnchor)
	beego.AddFuncMap("CrudIndexSQL", CrudIndexSQL)
	beego.AddFuncMap("CrudIsInGroup", CrudIsInGroup)
	beego.AddFuncMap("CrudIsAnonymous", CrudIsAnonymous)
	beego.AddFuncMap("CrudNumberToEnglish", CrudNumberToEnglish)
	beego.AddFuncMap("CrudMacro", CrudMacro)
	beego.AddFuncMap("CrudMacroSQL", CrudMacroSQL)
	beego.AddFuncMap("CrudClassSqlite", CrudClassSqlite)
	beego.AddFuncMap("CrudStyleSqlite", CrudStyleSqlite)
	beego.AddFuncMap("CrudSplit", CrudSplit)
	beego.AddFuncMap("CrudSQL", CrudSQL)
	beego.AddFuncMap("CrudIncrement", CrudIncrement)
	beego.AddFuncMap("CrudArgs", CrudArgs)
	beego.AddFuncMap("CrudComputeArgs", CrudComputeArgs)
	beego.AddFuncMap("CrudDecrement", CrudDecrement)
	beego.AddFuncMap("CrudDebug", CrudDebug)
	beego.AddFuncMap("BeeReplace", BeeReplace)
	beego.AddFuncMap("dict", Dictionary)
	beego.AddFuncMap("CrudComputeDataset", CrudComputeDataset)
	// Fonction sur dictionnaire
	beego.AddFuncMap("DictCreate", dict)
	beego.AddFuncMap("DictGet", get)
	beego.AddFuncMap("DictSet", set)
	beego.AddFuncMap("DictUnSet", unset)
	beego.AddFuncMap("DictAsKey", hasKey)
	beego.AddFuncMap("DictKeys", keys)
	beego.AddFuncMap("DictValues", values)
	// autres
	beego.AddFuncMap("markdown", markDowner)
	beego.AddFuncMap("style", style)
}

// style as
func style(in interface{}) template.HTMLAttr {
	if in.(string) != "" {
		return template.HTMLAttr("style=\"" + in.(string) + "\"")
	}
	return ""
}

// markDowner as
func markDowner(args ...interface{}) template.HTML {
	// s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(fmt.Sprintf("%s", args...)))

	// create HTML renderer with extensions
	htmlFlags := html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	return template.HTML(markdown.Render(doc, renderer))
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

// CrudNumberToEnglish as
func CrudNumberToEnglish(in int) string {
	out := ""
	switch in {
	case 1:
		out = "one"
	case 2:
		out = "two"
	case 3:
		out = "three"
	case 4:
		out = "four"
	case 5:
		out = "five"
	case 6:
		out = "six"
	case 7:
		out = "seven"
	case 8:
		out = "eight"
	case 9:
		out = "nine"
	}
	return out
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
func CrudIsInGroup(group string, session types.Session, appid string) (out bool) {
	out = false
	if session.Username == "anonymous" && appid != "" {
		if session.AppID != appid && session.Groups != group {
			return
		}
	}
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

// CrudIsAnonymous as
func CrudIsAnonymous(session types.Session) (out bool) {
	out = false
	if session.Username == "anonymous" {
		out = true
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

// CrudIndexSQL équivalent de index mais avec ComputeSqlite en +
func CrudIndexSQL(record orm.Params, key string, element dico.Element, session types.Session) (out string) {
	out = ""
	// if element.ComputeSqlite != "" {
	// 	out = CrudMacroSQL(element.ComputeSqlite, record, session)
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
func CrudFormat(in string, v interface{}) (out string) {
	value := ""
	switch v := v.(type) {
	case string:
		value = v
	case []byte:
		value = string(v)
	case error:
		value = v.Error()
	case fmt.Stringer:
		value = v.String()
	default:
		value = fmt.Sprintf("%v", v)
	}
	out = value
	if in != "" && value != "" {
		var recs []orm.Params
		var err error
		if in == "date" {
			if strings.Contains(value, "0001-01") {
				return ""
			}
			recs, err = models.CrudSQL("SELECT strftime('%Y-%m-%d','"+value+"')", "default")
		} else if in == "datetime" {
			recs, err = models.CrudSQL("SELECT strftime('%Y-%m-%d %H:%M:%S','"+value+"')", "default")
		} else if in == "time" {
			if strings.Contains(value, "00-00-00") {
				return ""
			}
			recs, err = models.CrudSQL("SELECT strftime('%H:%M:%S','"+value+"')", "default")
		} else if in == "amount" {
			ac := accounting.Accounting{Symbol: "€", Precision: 2, Thousand: " ", Decimal: ",", Format: "%v %s"}
			fl, _ := strconv.ParseFloat(value, 32)
			out = ac.FormatMoney(fl)
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
// list : "item 1,item 2,item 3,..."
// in   : "item 2"
// ret  : valeur à retourner si OK
func CrudContains(list string, in string) (out bool) {
	elements := strings.Split(list, ",")
	out = false
	for _, element := range elements {
		if in == element {
			out = true
		}
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
func CrudMacro(in string, appid string, record orm.Params, session types.Session) (out string) {
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
					if p, ok := dico.Ctx.Applications[appid].Parameters[key]; ok {
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
func CrudMacroSQL(in string, appid string, record orm.Params, session types.Session) (out string) {
	out = ""
	sql := CrudMacro(in, appid, record, session)
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

// CrudClassSqlite retourne le résulat de la requête avec macro
func CrudClassSqlite(element dico.Element, appid string, record orm.Params, session types.Session) (out string) {
	// out = CrudMacro(element.Class, record, session)
	// if out == "" {
	sql := CrudMacro(element.ClassSqlite, appid, record, session)
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

// CrudStyleSqlite retourne le résulat de la requête avec macro
func CrudStyleSqlite(element dico.Element, appid string, record orm.Params, session types.Session) (out string) {
	sql := CrudMacro(element.StyleSqlite, appid, record, session)
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

// CrudComputeDataset retourne le dataset valorisé
func CrudComputeDataset(dataset map[string]string, appid string, record orm.Params, session types.Session, aliasDB string) map[string]string {
	out := make(map[string]string, len(dataset))
	for key, value := range dataset {
		sql := CrudMacro(value, appid, record, session)
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

// CrudComputeArgs retourne le dataset valorisé
func CrudComputeArgs(dataset map[string]string, appid string, record orm.Params, session types.Session, aliasDB string) map[string]string {
	out := make(map[string]string, len(dataset))
	for key, value := range dataset {
		val := CrudMacro(value, appid, record, session)
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

// http://masterminds.github.io/sprig/dicts.html

func get(d map[string]interface{}, key string) interface{} {
	if val, ok := d[key]; ok {
		return val
	}
	return ""
}

func set(d map[string]interface{}, key string, value interface{}) map[string]interface{} {
	d[key] = value
	return d
}

func unset(d map[string]interface{}, key string) map[string]interface{} {
	delete(d, key)
	return d
}

func hasKey(d map[string]interface{}, key string) bool {
	_, ok := d[key]
	return ok
}

func keys(dicts ...map[string]interface{}) []string {
	k := []string{}
	for _, dict := range dicts {
		for key := range dict {
			k = append(k, key)
		}
	}
	return k
}
func values(dict map[string]interface{}) []interface{} {
	values := []interface{}{}
	for _, value := range dict {
		values = append(values, value)
	}

	return values
}

func dict(v ...interface{}) map[string]interface{} {
	dict := map[string]interface{}{}
	lenv := len(v)
	for i := 0; i < lenv; i += 2 {
		key := strval(v[i])
		if i+1 >= lenv {
			dict[key] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}
func strval(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}
