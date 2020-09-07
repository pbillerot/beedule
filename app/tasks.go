package app

import (
	"github.com/pbillerot/beedule/types"
)

// Tasks Mes Tâches
var Tasks = types.Table{
	AliasDB:    "tasks",
	Key:        "task_id",
	ColDisplay: "task_id",
	IconName:   "tasks",
	Elements:   tasksElements,
	Views:      tasksViews,
	Forms:      tasksForms,
}

var tasksElements = types.Elements{
	"task_id": {
		Type:       "counter",
		Order:      1,
		LabelLong:  "N° tâche",
		LabelShort: "N° tâche",
	},
	"task_user": {
		Type:       "text",
		LabelLong:  "Propriétaire",
		LabelShort: "Propriétaire",
		Default:    "{$user}",
	},
	"task_name": {
		Type:       "text",
		LabelLong:  "Nom de la tâche",
		LabelShort: "Nom de la tâche",
		Required:   true,
	},
	"task_status": {
		Type:       "checkbox",
		LabelLong:  "Terminée",
		LabelShort: "Terminée",
		ColAlign:   "center",
	},
	"task_note": {
		Type:       "textarea",
		LabelLong:  "Note",
		LabelShort: "Note",
	},
}

var tasksViews = types.Views{
	"vall": {
		FormView:  "fview",
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		Deletable: false,
		Info:      "Mes Tâches",
		Title:     "Tâches",
		IconName:  "tasks",
		ClassSQL:  "select 'positive' where '{task_status}' = '1'",
		OrderBy:   "task_status",
		// Where:   "task_user = '{$user}'",
		Elements: types.Elements{
			"task_id":     {Order: 10, HideOnMobile: true},
			"task_user":   {Order: 20},
			"task_name":   {Order: 30},
			"task_status": {Order: 40},
			"task_note":   {Order: 50, HideOnMobile: true},
		},
		ActionsSQL: types.Actions{
			{
				// on ne supprime que ses propres tâches
				Label: "Supprimer les tâches terminées...",
				SQL: []string{
					"delete from tasks where task_status = '1' and task_user = '{$user}'",
				},
				WithConfirm: true,
			},
		},
	},
}

var tasksForms = types.Forms{
	"fadd": {
		Title: "Fiche Compte",
		Elements: types.Elements{
			"task_id":     {Order: 10},
			"task_user":   {Order: 20, Protected: true},
			"task_name":   {Order: 30},
			"task_status": {Order: 40},
			"task_note":   {Order: 50},
		},
	},
	"fview": {
		Title: "Fiche Compte",
		Elements: types.Elements{
			"task_id":     {Order: 10},
			"task_name":   {Order: 30},
			"task_status": {Order: 40},
			"task_note":   {Order: 50},
		},
	},
	"fedit": {
		Title: "Fiche Compte",
		Elements: types.Elements{
			"task_id":     {Order: 10},
			"task_user":   {Order: 20, Protected: true},
			"task_name":   {Order: 30},
			"task_status": {Order: 40},
			"task_note":   {Order: 50},
		},
	},
}
