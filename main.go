package main

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pbillerot/beedule/app"
	_ "github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/batch"
	_ "github.com/pbillerot/beedule/routers"
	"github.com/pbillerot/beedule/types"
)

func init() {
	// Enregistrement des drivers des base de données
	// l'alias : Tables[alias]string
	// La déclaration dans conf/custom.conf [alias] drivertype= datasource= drivername=
	// default
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./database/beedule.sqlite")

	// boucle sur les applications pour charger les donnecteurs aux base de données
	// et les répertoires statiques du serveur
	for appid := range app.Applications {
		section, err := beego.AppConfig.GetSection(appid)
		if err != nil {
			beego.Error("GetSection", appid, err)
		}
		if datasource, ok := section["datasource"]; ok {
			drivertype, _ := strconv.Atoi(section["drivertype"])
			drivername := section["drivername"]
			orm.RegisterDriver(drivername, orm.DriverType(drivertype))
			orm.RegisterDataBase(appid, drivername, datasource)
			beego.Info("Enregistrement connecteur", appid, drivertype, drivername, datasource)
		}
	}
	ctx := make(map[string]string)
	for appid := range app.Applications {
		section, err := beego.AppConfig.GetSection(appid)
		if err != nil {
			beego.Error("GetSection", appid, err)
		}
		if datadir, ok := section["datadir"]; ok {
			if dataurl, ok := section["dataurl"]; ok {
				beego.SetStaticPath(dataurl, datadir)
				beego.Info("Enregistrement url static", dataurl, datadir)
				ctx["dataurl"] = dataurl
				ctx["datadir"] = datadir
			} else {
				dataurl = "/crud/data/" + appid
				beego.SetStaticPath(dataurl, datadir)
				beego.Info("Enregistrement url static", dataurl, datadir)
				ctx["dataurl"] = dataurl
				ctx["datadir"] = datadir
			}
			// app.Applications[appid].Ctx = &ctx
		}
	}

	if ok, _ := beego.AppConfig.Bool("debug"); ok {
		orm.Debug = true
	} else {
		orm.Debug = false
	}

	// Enregistrement des Modèles de table gérés par le module orm
	orm.RegisterModel(new(types.Parameters), new(batch.Chain), new(batch.Job), new(batch.Hugodoc))

	// Chargement des Parameters dans app.Params (préfixé par __)
	o := orm.NewOrm()
	o.Using(app.Parameters.AliasDB)
	var parameters []types.Parameters
	num, err := o.QueryTable("parameters").All(&parameters)
	if err != nil {
		beego.Error("parameters", err)
		return
	}
	if num > 0 {
		for _, parameter := range parameters {
			app.Params["__"+parameter.ID] = parameter.Value
		}
	}
	beego.Info("Params", app.Params)
	if param, ok := app.Params["__batch_etat"]; ok {
		if param == "1" {
			batch.StartBatch()
		}
	}

}
func main() {
	beego.Run()
}

// Run point d'entrée en tant que module
func Run() {
	beego.Run()
}
