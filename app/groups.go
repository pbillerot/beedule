package app

import (
	"github.com/pbillerot/beedule/types"
)

// Groups table des groupes d'utilisateurs
var Groups = types.Table{
	AliasDB:    "admin",
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
		Title:     "Groupes",
		IconName:  "users",
		OrderBy:   "group_id",
		Group:     "admin",
		Mask: types.MaskList{
			Header: []string{
				"group_id",
			},
			Meta: []string{},
			Description: []string{
				"group_note",
			},
			Extra: []string{},
		},
		Elements: types.Elements{
			"group_id":   {},
			"group_note": {},
		},
	},
}

var groupsForms = types.Forms{
	"fadd": {
		Title: "Fiche Groupe",
		Group: "admin",
		Elements: types.Elements{
			"group_id":   {},
			"group_note": {},
		},
	},
	"fedit": {
		Title: "Fiche Groupe",
		Group: "admin",
		Elements: types.Elements{
			"group_id":   {},
			"group_note": {},
		},
	},
}
