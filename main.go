package main

import (
	"strconv"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pbillerot/beedule/dico"
	_ "github.com/pbillerot/beedule/routers"
)

func init() {

	// Répertoire du dictionnaire
	// beego.AppConfig.String("dicodir")

	// Enregistrement des drivers des base de données
	// l'alias : Tables[alias]string
	// La déclaration dans conf/custom.conf [alias] drivertype= datasource= drivername=

	// default est utilisé pour exécuter les ordres de calcul sql par le moteur
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./database/beedule.sqlite")

	if ok, _ := beego.AppConfig.Bool("debug"); ok {
		orm.Debug = true
	} else {
		orm.Debug = false
	}

	// Chargement du dictionnaire
	dico.Ctx.Load()

	// boucle sur les tables pour charger les connecteurs aux bases de données
	// et les répertoires statiques du serveur
	for _, table := range dico.Ctx.Tables {
		// lecture dans les sections de app.conf pour enregistre les connecteurs aux bases de données
		section, err := beego.AppConfig.GetSection(table.Setting.AliasDB)
		if err == nil {
			_, err := orm.GetDB(table.Setting.AliasDB)
			if err != nil {
				// Cet alias n'a pas encore été déclaré
				if datasource, ok := section["datasource"]; ok {
					// la section correspondante a été trouvée
					drivertype, _ := strconv.Atoi(section["drivertype"])
					drivername := section["drivername"]
					logs.Info("...Enregistrement connecteur", table.Setting.AliasDB, drivertype, drivername, datasource)
					orm.RegisterDriver(drivername, orm.DriverType(drivertype))
					orm.RegisterDataBase(table.Setting.AliasDB, drivername, datasource)
					// Déclaration éventuelle du répertoire statique des applications
					if datadir, ok := section["datadir"]; ok {
						logs.Info("...Enregistrement statique", "/bee/data/"+table.Setting.AliasDB, datadir)
						beego.SetStaticPath("/bee/data/"+table.Setting.AliasDB, datadir)
					}
				} else {
					// ERR l'alias n'pas été déclaré dans app.conf
					logs.Error("ERREUR aliasDB non déclaré dans app.conf", table.Setting.AliasDB)
				}
			}
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
