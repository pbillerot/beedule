package app

import "github.com/pbillerot/beedule/types"

// Applications présentées sur le portail
var Applications = map[string]types.Application{
	"pendule": {
		Title: "Pendule",
		Image: "/crud/static/img/pendule.svg",
		Group: "admin",
		AppViews: []types.AppView{
			{Tableid: "chains", Viewid: "vall"},
			{Tableid: "jobs", Viewid: "vall"},
		},
	},
	"users": {
		Title: "Gestion des utilisateurs et des groupes",
		Image: "/crud/static/img/groups.svg",
		Group: "admin",
		AppViews: []types.AppView{
			{Tableid: "users", Viewid: "vall"},
			{Tableid: "groups", Viewid: "vall"},
		},
	},
	"tasks": {
		Title: "Mes Tâches",
		Group: "admin",
		Image: "/crud/static/img/tasks.svg",
		AppViews: []types.AppView{
			{Tableid: "tasks", Viewid: "vall"},
		},
	},
	"picsou": {
		Title:    "Picsou, pour bricoler sur la bourse",
		Group:    "picsou",
		Image:    "/crud/static/img/picsou.jpg",
		DataPath: "", // valorisé dans custom.conf
		AppViews: []types.AppView{
			{Tableid: "orders", Viewid: "vachat"},
			{Tableid: "orders", Viewid: "vvente"},
			{Tableid: "ptf", Viewid: "vall"},
			{Tableid: "ptf", Viewid: "vdiapo"},
			{Tableid: "ptf", Viewid: "vtop"},
			{Tableid: "ptf", Viewid: "vntop"},
			{Tableid: "quotes", Viewid: "vlast"},
		},
	},
	"filebrowser": {
		Title: "Gestionnaire de fichiers",
		Image: "/crud/static/img/filebrowser.svg",
		// Path:  "http://localhost:7602/fb",
		Path: "/fb",
	},
	"navidrome": {
		Title: "Serveur de streaming de musique",
		Image: "/crud/static/img/navidrome.png",
		Path:  "/nv",
	},
}

// Tables Liens vers les tables
// "nom de la table": Structure
var Tables = types.Tables{
	"tasks":  Tasks,
	"users":  Users,
	"groups": Groups,
	"orders": Orders,
	"ptf":    Ptf,
	"quotes": Quotes,
	"chains": Chains,
	"jobs":   Jobs,
}

// Portail as
var Portail = types.Portail{
	Title:        "Beedule",
	Info:         "Le prototype d'utilisation du framework",
	IconFile:     "/crud/static/img/beedule.png",
	Applications: Applications,
	Tables:       Tables,
}

// Params paramètres globaux aux applications
var Params = map[string]string{
	"__cost":       "0.0047", // Coût transaction 0.47 %
	"__optimum":    "0.03",   // Seuil rentabilité à 3 %
	"__amount_min": "1200",   // Montant minimum de la transaction
}
