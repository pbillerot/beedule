package models

import (
	"errors"
	"strings"

	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/types"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
func CrudList(tableid string, viewid string, view *types.View, elements types.Elements) ([]orm.Params, error) {

	// Rédaction de la requête
	var keys []string
	joins := []string{}
	for k, element := range elements {
		if strings.HasPrefix(k, "_") {
			keys = append(keys, "'' as "+k)
			continue
		}
		if element.Type == "section" {
			keys = append(keys, "'' as "+k)
			continue
		}
		if element.Jointure.Join != "" {
			keys = append(keys, element.Jointure.Column+" as "+k)
			joins = append(joins, element.Jointure.Join)
		} else {
			keys = append(keys, tableid+"."+k)
		}
	}
	skey := strings.Join(keys, ", ")

	where := ""
	if view.Where != "" {
		where = " WHERE " + view.Where
	}
	orderby := ""
	if view.OrderBy != "" {
		orderby = " ORDER BY " + view.OrderBy
	}

	o := orm.NewOrm()
	o.Using(app.Tables[tableid].AliasDB)
	var maps []orm.Params
	sql := "SELECT " + skey +
		" FROM " + tableid +
		" " + strings.Join(joins, " ") +
		where +
		orderby
	_, err = o.Raw(sql).Values(&maps)
	if err != nil {
		beego.Error(err)
	}
	return maps, err
}

// CrudRead as
func CrudRead(tableid string, id string, elements types.Elements) ([]orm.Params, error) {

	// Rédaction de la requête
	var keys []string
	joins := []string{}
	for k, element := range elements {
		if strings.HasPrefix(k, "_") {
			keys = append(keys, "'' as "+k)
			continue
		}
		if element.Type == "section" {
			keys = append(keys, "'' as "+k)
			continue
		}
		if element.Jointure.Join != "" {
			keys = append(keys, element.Jointure.Column+" as "+k)
			joins = append(joins, element.Jointure.Join)
		} else {
			keys = append(keys, tableid+"."+k)
		}
	}
	skey := strings.Join(keys, ", ")

	o := orm.NewOrm()
	o.Using(app.Tables[tableid].AliasDB)
	var maps []orm.Params
	sql := "SELECT " + skey +
		" FROM " + tableid +
		" " + strings.Join(joins, " ") +
		" WHERE " + app.Tables[tableid].Key + " = ?"
	num, err := o.Raw(sql, id).
		Values(&maps)
	if err != nil {
		beego.Error(err)
	} else if num == 0 {
		beego.Error(app.Tables[tableid].AliasDB, sql)
		err = errors.New("Article non trouvé")
	}
	return maps, err
}

// CrudUpdate as
func CrudUpdate(tableid string, id string, elements types.Elements) error {

	// Remplissage de la map des valeurs
	sql := ""
	args := []string{}
	for k, element := range elements {
		if strings.HasPrefix(k, "_") {
			continue
		}
		if k == app.Tables[tableid].Key {
			continue
		}
		if element.Type == "counter" {
			continue
		}
		if element.Type == "section" {
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
	sql += " WHERE " + app.Tables[tableid].Key + " = ?"
	args = append(args, id)

	o := orm.NewOrm()
	o.Using(app.Tables[tableid].AliasDB)
	_, err := o.Raw(sql, args).Exec()
	if err != nil {
		beego.Error(app.Tables[tableid].AliasDB, sql)
		beego.Error(err)
	}
	return err
}

// CrudInsert as
func CrudInsert(tableid string, elements types.Elements) error {

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
		if element.Type == "section" {
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

	o := orm.NewOrm()
	o.Using(app.Tables[tableid].AliasDB)
	_, err := o.Raw(sql, args).Exec()
	if err != nil {
		beego.Error(app.Tables[tableid].AliasDB, sql)
		beego.Error(err)
	}
	return err
}

// CrudDelete as
func CrudDelete(tableid string, id string) error {

	o := orm.NewOrm()
	o.Using(app.Tables[tableid].AliasDB)
	_, err := o.Raw("DELETE FROM "+tableid+
		" WHERE "+app.Tables[tableid].Key+" = ?", id).Exec()
	if err != nil {
		beego.Error(err)
	}
	return err
}

// CrudSQL as
func CrudSQL(sql string, aliasDB string) ([]orm.Params, error) {
	o := orm.NewOrm()
	o.Using(aliasDB)
	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		beego.Error(aliasDB, sql)
		beego.Error(err)
	}
	return maps, err
}

// CrudExec as
func CrudExec(sql string, aliasDB string) error {

	o := orm.NewOrm()
	o.Using(aliasDB)
	_, err := o.Raw(sql).Exec()
	if err != nil {
		beego.Error(aliasDB, sql)
		beego.Error(err)
	}
	return err
}
