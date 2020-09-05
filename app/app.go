package app

import "github.com/pbillerot/beedule/types"

// Applications présentées sur le portail
var Applications = map[string]types.Application{
	"users": {
		Title: "Gestion des utilisateurs et des groupes",
		Image: "/crud/static/img/groups.svg",
		AppViews: []types.AppView{
			{Tableid: "users", Viewid: "vall"},
			{Tableid: "groups", Viewid: "vall"},
		},
	},
	"tasks": {
		Title: "Mes Tâches",
		Image: "/crud/static/img/tasks.svg",
		AppViews: []types.AppView{
			{Tableid: "tasks", Viewid: "vall"},
		},
	},
	"picsou": {
		Title: "Picsou, pour bricoler sur la bourse",
		Image: "/crud/static/img/picsou.jpg",
		AppViews: []types.AppView{
			{Tableid: "orders", Viewid: "vachat"},
			{Tableid: "orders", Viewid: "vvente"},
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

// Portail as
var Portail = types.Portail{
	Title:        "Beedule",
	Info:         "Le prototype d'utilisation du framework",
	IconFile:     "/crud/static/img/beedule.png",
	Applications: Applications,
	Tables:       Tables,
}

// Tables Liens vers les tables
// "nom de la table": Structure
var Tables = types.Tables{
	"tasks":  Tasks,
	"users":  Users,
	"groups": Groups,
	"orders": Orders,
}

// Params paramètres globaux aux applications
var Params = map[string]string{
	"__cout": "0.0047",
}
