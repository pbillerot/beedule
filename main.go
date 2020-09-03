package main

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pbillerot/beedule/app"
	_ "github.com/pbillerot/beedule/app"
	_ "github.com/pbillerot/beedule/routers"
)

func init() {
	// Enregistrement des drivers des base de données
	// l'alias : Tables[alias]string
	// La déclaration dans conf/custom.conf [alias] drivertype= datasource= drivername=

	// default
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./database/beedule.sqlite")

	aliass := map[string]bool{}
	for _, table := range app.Tables {
		alias := table.AliasDB
		section, _ := beego.AppConfig.GetSection(alias)
		drivertype, _ := strconv.Atoi(section["drivertype"])
		drivername := section["drivername"]
		datasource := section["datasource"]
		if _, ok := aliass[alias]; ok == false {
			aliass[alias] = true
			orm.RegisterDriver(drivername, orm.DriverType(drivertype))
			orm.RegisterDataBase(alias, drivername, datasource)
			beego.Info("Enregistrement de", alias, drivertype, drivername, datasource)
		}
	}
	if beego.AppConfig.String("debug") == "true" {
		orm.Debug = true
	} else {
		orm.Debug = false
	}
}
func main() {
	beego.Run()
}

// Run point d'entrée en tant que module
func Run() {
	beego.Run()
}
