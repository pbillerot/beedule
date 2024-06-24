package main

import (
	"fmt"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/task"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pbillerot/beedule/dico"
	"github.com/pbillerot/beedule/models"
	_ "github.com/pbillerot/beedule/routers"
)

func init() {

	// Init de la structure config
	models.Config.Appname = beego.AppConfig.String("appname")
	models.Config.Appnote = beego.AppConfig.String("appnote")
	models.Config.Date = beego.AppConfig.String("date")
	models.Config.Icone = beego.AppConfig.String("icone")
	models.Config.Site = beego.AppConfig.String("site")
	models.Config.Email = beego.AppConfig.String("email")
	models.Config.Author = beego.AppConfig.String("author")
	models.Config.Version = beego.AppConfig.String("version")
	models.Config.Theme = beego.AppConfig.String("theme")
	models.Config.LogPath = beego.AppConfig.String("logpath")
	models.Config.Portail = beego.AppConfig.String("portail")
	models.Config.TaskCron = beego.AppConfig.String("taskcron")

	// Config du Logger
	beego.BConfig.Log.AccessLogs = true
	logs.Async()
	pathLog := fmt.Sprintf(`{"filename":"%s","maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`, models.Config.LogPath)
	logs.SetLogger(logs.AdapterFile, pathLog)

	validation.SetDefaultMessage(map[string]string{
		"Required":     "est obligatoire",
		"Min":          "Valeur minimum de %d",
		"Max":          "Valeur maximum de %d",
		"Range":        "La valeur doit être comprise entre %d et %d",
		"MinSize":      "Longueur de %d caractères minimum",
		"MaxSize":      "Longueur de %d caractères maximum",
		"Length":       "Le nombre de caractères doit être égale à %d",
		"Alpha":        "Caractères alpha uniquement",
		"Numeric":      "Caractères numérique uniquement",
		"AlphaNumeric": "Caractères alpha-numériques uniquement",
		"Match":        "mot %s interdit",
		"NoMatch":      "mot %s non trouvé",
		"AlphaDash":    "Caractères (-_) interdits",
		"Email":        "Email incorrect",
		"IP":           "N° ip incorrect",
		"Base64":       "Base64 incorrect",
		"Mobile":       "Téléphone mobile incorrect",
		"Tel":          "Téléphone incorrect",
		"Phone":        "Téléphone incorrect",
		"ZipCode":      "Code postal incorrect",
	})

	// default est utilisé pour exécuter les ordres de calcul sql par le moteur
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("dbsqlite"))
	// Enregistrement du driver postgres
	orm.RegisterDriver("postgres", orm.DRPostgres)

	if ok, _ := beego.AppConfig.Bool("debug"); ok {
		orm.Debug = true
	} else {
		orm.Debug = false
	}

	// Chargement du dictionnaire
	dico.Ctx.Load()

	// Initialisation des tâches planifiées (table tasks)
	// https://github.com/beego/beedoc/blob/master/en-US/module/task.md
	// https//pkg.go.dev/github.com/beego/beego/v2@v2.0.7/task
	tk := task.NewTask("Beedule tasks...", models.Config.TaskCron, models.EveryDay)
	task.AddTask("Beedule tasks...", tk)
	task.StartTask()
	// voir crud_modele.EveryDay() pour le traitement des tasks

}
func main() {
	beego.Run()
}

// Run point d'entrée en tant que module
func Run() {
	beego.Run()
}
