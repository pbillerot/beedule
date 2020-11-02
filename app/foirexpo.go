package app

import (
	"fmt"

	"github.com/pbillerot/beedule/types"
)

// Hugodocument Les documents du site foirexpo
var Hugodocument = types.Table{
	AliasDB:    "foiredit",
	Key:        "path",
	ColDisplay: "path",
	IconName:   "sitemap",
	Elements:   hugoElements,
	Views:      hugoViews,
	Forms:      hugoForms,
}

var hugoElements = types.Elements{
	"path": {
		Type:       "text",
		LabelLong:  "Chemin",
		LabelShort: "Chemin",
	},
	"base": {
		Type:       "text",
		LabelLong:  "Fichier",
		LabelShort: "Nom Fichier",
	},
	"dir": {
		Type:       "text",
		LabelLong:  "Localisation",
		LabelShort: "Localisation",
	},
	"ext": {
		Type:       "text",
		LabelLong:  "Extension",
		LabelShort: "Extension",
	},
	"isdir": {
		Type:       "checkbox",
		LabelLong:  "Est un répertoire",
		LabelShort: "Rép.",
	},
	"title": {
		Type:       "text",
		LabelLong:  "Titre",
		LabelShort: "Titre",
	},
	"date": {
		Type:       "date",
		LabelLong:  "Date",
		LabelShort: "Date",
	},
	"tags": {
		Type:       "tag",
		LabelLong:  "Tag",
		LabelShort: "Tag",
	},
	"categories": {
		Type:       "tag",
		LabelLong:  "Catégorie",
		LabelShort: "Catégorie",
	},
	"draft": {
		Type:       "checkbox",
		LabelLong:  "Brouillon",
		LabelShort: "Brouillon",
	},
}

var hugoViews = types.Views{
	"vall": {
		// FormView:  "fview",
		FormEdit: "fedit",
		Title:    "Tous les documents",
		IconName: "sitemap",
		OrderBy:  "dir,base",
		Type:     "table",
		Elements: types.Elements{
			"path":       {Order: 1, Hide: true},
			"isdir":      {Order: 5},
			"dir":        {Order: 20},
			"base":       {Order: 30},
			"ext":        {Order: 35},
			"draft":      {Order: 40},
			"title":      {Order: 50},
			"date":       {Order: 60},
			"tags":       {Order: 70},
			"categories": {Order: 80},
		},
		Actions: types.Actions{
			{
				// on ne supprime que ses propres tâches
				Label:  "Recharger le répertoire",
				Plugin: fmt.Sprintf("hugoDirectoriesToSQL(%s,%s,%s)", "/home/billerot/Abri/foirexpo", "hugodocument", "foiredit"), // path,table,aliasDB
			},
		},
	},
}

var hugoForms = types.Forms{
	"fview": {
		Title: "Document",
		Elements: types.Elements{
			"path":       {Order: 10},
			"isdir":      {Order: 5},
			"base":       {Order: 20},
			"draft":      {Order: 30},
			"title":      {Order: 40},
			"date":       {Order: 50},
			"tags":       {Order: 60},
			"categories": {Order: 70},
		},
	},
	"fedit": {
		Title: "Document",
		Elements: types.Elements{
			"path":       {Order: 10},
			"isdir":      {Order: 5},
			"base":       {Order: 20},
			"draft":      {Order: 30},
			"title":      {Order: 40},
			"date":       {Order: 50},
			"tags":       {Order: 60},
			"categories": {Order: 70},
		},
	},
}
