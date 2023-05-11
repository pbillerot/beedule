package models

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/types"

	"github.com/beego/beego/v2/client/orm"
)

var err error

// Config de config.yaml
var Config types.BeeConfig

// IfNotEmpty as
func IfNotEmpty(chaine string, valTrue string, valFalse string) string {
	if chaine != "" {
		return valTrue
	}
	return valFalse
}

// https://beego.me/docs/mvc/model/querybuilder.md

// CrudList as
func CrudList(appid string, tableid string, viewid string, view *dico.View, elements map[string]dico.Element) ([]orm.Params, error) {

	// Rédaction de la requête
	var keys []string
	joins := []string{}
	type Sorter struct {
		id       string
		direcion string
	}
	sorter := Sorter{}
	for k, element := range elements {
		if strings.HasPrefix(k, "_") {
			keys = append(keys, "'' as "+k)
			continue
		}
		if element.Jointure.Column != "" {
			keys = append(keys, element.Jointure.Column+" as "+k)
			joins = append(joins, element.Jointure.Join)
		} else {
			keys = append(keys, tableid+"."+k)
		}
		if element.SortDirection != "" {
			if element.SortDirection == "ascending" {
				sorter.id = k
				sorter.direcion = "ASC"
			} else {
				sorter.id = k
				sorter.direcion = "DESC"
			}
		}
	}
	skey := strings.Join(keys, ", ")

	where := ""
	if view.Where != "" {
		where = " WHERE " + view.Where
	}
	if view.Search != "" {
		if view.Where != "" {
			where += " AND (" + view.Search + ")"
		} else {
			where += " WHERE " + view.Search
		}
	}

	limit := ""
	if view.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", view.Limit)
	}
	orderby := ""
	if view.OrderBy != "" {
		orderby = " ORDER BY " + view.OrderBy
	}
	if sorter.id != "" {
		orderby = " ORDER BY " + sorter.id + " " + sorter.direcion
	}

	o := orm.NewOrmUsingDB(dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB)
	var maps []orm.Params
	sql := "SELECT " + skey +
		" FROM " + tableid +
		" " + strings.Join(joins, " ") +
		where +
		orderby +
		limit
	_, err = o.Raw(sql).Values(&maps)
	if err != nil {
		logs.Error(err)
	}
	// logs.Debug(fmt.Sprintf("#%v", maps))
	return maps, err
}

// CrudRead as
func CrudRead(filter string, appid string, tableid string, id string, elements map[string]dico.Element) ([]orm.Params, error) {

	// Rédaction de la requête
	var keys []string
	joins := []string{}
	for k, element := range elements {
		// une jointure ne doit pas être préfixé par un _
		if strings.HasPrefix(k, "_") {
			keys = append(keys, "'' as "+k)
			continue
		}
		if element.Jointure.Column != "" {
			keys = append(keys, element.Jointure.Column+" as "+k)
			joins = append(joins, element.Jointure.Join)
		} else {
			keys = append(keys, tableid+"."+k)
		}

	}
	skey := strings.Join(keys, ", ")

	o := orm.NewOrmUsingDB(dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB)
	var maps []orm.Params
	keyid := tableid + "." + dico.Ctx.Applications[appid].Tables[tableid].Setting.Key
	if elements[dico.Ctx.Applications[appid].Tables[tableid].Setting.Key].Jointure.Column != "" {
		keyid = elements[dico.Ctx.Applications[appid].Tables[tableid].Setting.Key].Jointure.Column
	}
	sql := "SELECT " + skey +
		" FROM " + tableid +
		" " + strings.Join(joins, " ") +
		" WHERE " + keyid + " = ?"
	if filter != "" {
		sql += " AND " + filter
	}
	num, err := o.Raw(sql, id).Values(&maps)
	if err != nil {
		logs.Error(err)
	} else if num == 0 {
		logs.Error(dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB, sql)
		err = errors.New("enregistrement non trouvé")
	}
	return maps, err
}

// CrudUpdate avec element.SQLout
func CrudUpdate(appid string, tableid string, id string, elements map[string]dico.Element) error {

	// Remplissage de la map des valeurs
	sql := ""
	args := []string{}
	for k, element := range elements {
		if strings.HasPrefix(k, "_") {
			continue
		}
		if k == dico.Ctx.Applications[appid].Tables[tableid].Setting.Key {
			continue
		}
		if element.Type == "counter" {
			continue
		}
		if element.Jointure.Join != "" {
			continue
		}
		if len(sql) == 0 {
			sql += "UPDATE " + tableid + " SET "
		} else {
			sql += ", "
		}
		sql += k + " = ?"
		args = append(args, element.SQLout)
	}
	if len(sql) == 0 {
		// Pas de champ à mettre à jour
		return nil
	}
	sql += " WHERE " + tableid + "." + dico.Ctx.Applications[appid].Tables[tableid].Setting.Key + " = ?"
	args = append(args, id)

	o := orm.NewOrmUsingDB(dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB)
	_, err := o.Raw(sql, args).Exec()
	if err != nil {
		logs.Error(dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB, sql)
		logs.Error(err)
	}
	return err
}

// CrudInsert avec element.SQLout
func CrudInsert(appid string, tableid string, elements map[string]dico.Element) error {

	// Construction de l'ordre sql
	sqlcol := ""
	sqlval := ""
	bstart := true
	args := []string{}
	for k, element := range elements {
		if strings.HasPrefix(k, "_") {
			continue
		}
		if element.Type == "counter" {
			continue
		}
		if element.Jointure.Join != "" {
			continue
		}
		if !bstart {
			sqlcol += ", "
			sqlval += ", "
		}
		bstart = false
		sqlcol += k
		sqlval += "?"
		args = append(args, element.SQLout)
	}
	sql := "INSERT INTO " + tableid + " (" + sqlcol + ") VALUES (" + sqlval + ")"

	o := orm.NewOrmUsingDB(dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB)
	_, err := o.Raw(sql, args).Exec()
	if err != nil {
		logs.Error(dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB, sql)
		logs.Error(err)
	}
	return err
}

// CrudDelete as
func CrudDelete(appid, tableid string, id string) error {

	o := orm.NewOrmUsingDB(dico.Ctx.Applications[appid].Tables[tableid].Setting.AliasDB)
	_, err := o.Raw("DELETE FROM "+tableid+
		" WHERE "+dico.Ctx.Applications[appid].Tables[tableid].Setting.Key+" = ?", id).Exec()
	if err != nil {
		logs.Error(err)
	}
	return err
}

// CrudSQL as
func CrudSQL(sql string, aliasDB string) ([]orm.Params, error) {
	o := orm.NewOrmUsingDB(aliasDB)
	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		logs.Error(aliasDB, sql)
		logs.Error(err)
	}
	return maps, err
}

// CrudExec as
func CrudExec(sql string, aliasDB string) error {
	o := orm.NewOrmUsingDB(aliasDB)
	_, err := o.Raw(sql).Exec()
	if err != nil {
		logs.Error(aliasDB, sql)
		logs.Error(err)
	}
	return err
}

// EveryDay as appelé par le planificateur des tâches tous les jours
/**
CREATE TABLE "tasks" (
	"id"	INTEGER,
	"name"	TEXT,
	"day"	INTEGER DEFAULT 0,
	"month"	INTEGER DEFAULT 0,
	"last_day"	INTEGER DEFAULT 0,
	"last_month"	INTEGER DEFAULT 0,
	"disabled"	INTEGER,
	"sql"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
)
*/
func EveryDay(ctx context.Context) error {
	now := time.Now()
	now_day := now.Day()
	now_month := int(now.Month())
	logs.Info("o_o everyDay", now_day, now_month)
	for _, application := range dico.Ctx.Applications {
		if application.TasksTableName != "" {
			sql := fmt.Sprintf("select * from %s where enabled = 1 order by priority", application.TasksTableName)
			recs, err := CrudSQL(sql, application.AliasDB)
			if err == nil {
				for _, rec := range recs {
					id, _ := strconv.Atoi(rec["id"].(string))
					day, _ := strconv.Atoi(rec["day"].(string))
					last_day, _ := strconv.Atoi(rec["last_day"].(string))
					month, _ := strconv.Atoi(rec["month"].(string))
					last_month, _ := strconv.Atoi(rec["last_month"].(string))
					if last_month < now_month || (now_month < 12 && last_month == 12) {
						// maj début de mois -> raz last_day
						last_day = 0
					}
					if month == 0 || (last_month < month && month == now_month) {
						// month ok
						if day == 0 || (last_day < day && day == now_day) {
							// day ok
							result := ""
							// exécution du sql
							if rec["sql"] != nil && rec["sql"].(string) != "" {
								err = CrudExec(rec["sql"].(string), application.AliasDB)
								if err != nil {
								// maj result
								maj := fmt.Sprintf("update %s set result = '%s' where id = %d",
									application.TasksTableName, result, id)
								CrudExec(maj, application.AliasDB)
								Continue
							}
							// exécution du shell
							if rec["shell"] != nil && rec["shell"].(string) != "" {
								result, err = ShellExec(rec["shell"].(string))
								if err != nil {
									// maj result
									maj := fmt.Sprintf("update %s set result = '%s' where id = %d",
										application.TasksTableName, result, id)
									CrudExec(maj, application.AliasDB)
									Continue
								}
							}
							// maj planif
							maj := fmt.Sprintf("update %s set last_day = %d, last_month = %d where id = %d",
								application.TasksTableName, now_day, now_month, id)
							err = CrudExec(maj, application.AliasDB)
							if err != nil {
								continue
							}
						}
					}
				}
			}
		}
		// if application.Batman != "" {
		// 	ShellExec(application.Batman)
		// }
	}
	return nil
}

func ShellExec(commande string) (out string, err error) {
	out = ""
	if commande != "" {
		var stdout, stderr bytes.Buffer
		cmd := exec.Command("/bin/sh", "-c", commande)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		logs.Info("shell", commande)
		err := cmd.Run()
		if err != nil {
			logs.Error("could not run command: ", err)
		}
		logs.Info(strings.TrimSpace(stderr.String()))
		logs.Info(strings.TrimSpace(stdout.String()))
		out = fmt.Sprintf("%s\n%s", strings.TrimSpace(stdout.String()), strings.TrimSpace(stderr.String()))
	}
	return out, err
}
