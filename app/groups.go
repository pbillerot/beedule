package app

import (
	"github.com/pbillerot/beedule/types"
)

// Groups table des groupes d'utilisateurs
var Groups = types.Table{
	AliasDB:    "users",
	Key:        "group_id",
	ColDisplay: "group_id",
	IconName:   "users",
	Elements:   groupsElements,
	Views:      groupsViews,
	Forms:      groupsForms,
}

var groupsElements = types.Elements{
	"group_id": {
		Type:       "text",
		Order:      1,
		LabelLong:  "Groupes",
		LabelShort: "Groupes",
	},
	"group_note": {
		Type:       "textarea",
		Order:      10,
		LabelLong:  "Note",
		LabelShort: "Note",
	},
}

var groupsViews = types.Views{
	"vall": {
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		Deletable: true,
		Info:      "Gestion des Groupes",
		Title:     "Groupes",
		IconName:  "users",
		OrderBy:   "group_id",
		Group:     "admin",
		Elements: types.Elements{
			"group_id":   {},
			"group_note": {},
		},
	},
}

var groupsForms = types.Forms{
	"fadd": {
		Title:  "Fiche Groupe",
		Groupe: "admin",
		Elements: types.Elements{
			"group_id":   {},
			"group_note": {},
		},
	},
	"fedit": {
		Title:  "Fiche Groupe",
		Groupe: "admin",
		Elements: types.Elements{
			"group_id":   {},
			"group_note": {},
		},
	},
}
