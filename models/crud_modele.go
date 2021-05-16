package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/pbillerot/beedule/dico"

	"github.com/beego/beego/v2/client/orm"
)

var err error

// IfNotEmpty as
func IfNotEmpty(chaine string, valTrue string, valFalse string) string {
	if chaine != "" {
		return valTrue
	}
	return valFalse
}

// https://beego.me/docs/mvc/model/querybuilder.md

// CrudList as
func CrudList(tableid string, viewid string, view *dico.View, elements map[string]dico.Element) ([]orm.Params, error) {

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

	o := orm.NewOrmUsingDB(dico.Ctx.Tables[tableid].Setting.AliasDB)
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
func CrudRead(filter string, tableid string, id string, elements map[string]dico.Element) ([]orm.Params, error) {

	// Rédaction de la requête
	var keys []string
	joins := []string{}
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

	}
	skey := strings.Join(keys, ", ")

	o := orm.NewOrmUsingDB(dico.Ctx.Tables[tableid].Setting.AliasDB)
	var maps []orm.Params
	sql := "SELECT " + skey +
		" FROM " + tableid +
		" " + strings.Join(joins, " ") +
		" WHERE " + dico.Ctx.Tables[tableid].Setting.Key + " = ?"
	if filter != "" {
		sql += " AND " + filter
	}
	num, err := o.Raw(sql, id).Values(&maps)
	if err != nil {
		logs.Error(err)
	} else if num == 0 {
		logs.Error(dico.Ctx.Tables[tableid].Setting.AliasDB, sql)
		err = errors.New("Enregistrement non trouvé")
	}
	return maps, err
}

// CrudUpdate avec element.SQLout
func CrudUpdate(tableid string, id string, elements map[string]dico.Element) error {

	// Remplissage de la map des valeurs
	sql := ""
	args := []string{}
	for k, element := range elements {
		if strings.HasPrefix(k, "_") {
			continue
		}
		if k == dico.Ctx.Tables[tableid].Setting.Key {
			continue
		}
		if element.Type == "counter" {
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
	sql += " WHERE " + dico.Ctx.Tables[tableid].Setting.Key + " = ?"
	args = append(args, id)

	o := orm.NewOrmUsingDB(dico.Ctx.Tables[tableid].Setting.AliasDB)
	_, err := o.Raw(sql, args).Exec()
	if err != nil {
		logs.Error(dico.Ctx.Tables[tableid].Setting.AliasDB, sql)
		logs.Error(err)
	}
	return err
}

// CrudInsert avec element.SQLout
func CrudInsert(tableid string, elements map[string]dico.Element) error {

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

	o := orm.NewOrmUsingDB(dico.Ctx.Tables[tableid].Setting.AliasDB)
	_, err := o.Raw(sql, args).Exec()
	if err != nil {
		logs.Error(dico.Ctx.Tables[tableid].Setting.AliasDB, sql)
		logs.Error(err)
	}
	return err
}

// CrudDelete as
func CrudDelete(tableid string, id string) error {

	o := orm.NewOrmUsingDB(dico.Ctx.Tables[tableid].Setting.AliasDB)
	_, err := o.Raw("DELETE FROM "+tableid+
		" WHERE "+dico.Ctx.Tables[tableid].Setting.Key+" = ?", id).Exec()
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
