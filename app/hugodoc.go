package app

import (
	"fmt"

	"github.com/pbillerot/beedule/types"
)

// Hugodoc Les documents du site
var Hugodoc = types.Table{
	AliasDB:    "foiredit",
	Key:        "id",
	ColDisplay: "path",
	IconName:   "sitemap",
	Elements:   hugoElements,
	Views:      hugoViews,
	Forms:      hugoForms,
}

var hugoElements = types.Elements{
	"id": {
		Type:       "counter",
		LabelLong:  "id",
		LabelShort: "id",
	},
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
	"level": {
		Type:       "number",
		LabelLong:  "Niveau",
		LabelShort: "Niveau",
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
	"content": {
		Type:       "editor",
		LabelLong:  "Contenu",
		LabelShort: "Contenu",
	},
}

var hugoViews = types.Views{
	"vall": {
		Title:    "La table brute",
		IconName: "list",
		Type:     "table",
		FormView: "fview",
		FormEdit: "fdoc",
		// OrderBy:  "path,isdir,base",
		Elements: types.Elements{
			"id":         {Order: 1, Hide: true},
			"path":       {Order: 10, Hide: true},
			"isdir":      {Order: 20},
			"level":      {Order: 30},
			"dir":        {Order: 40},
			"base":       {Order: 50},
			"ext":        {Order: 60},
			"title":      {Order: 100},
			"draft":      {Order: 110},
			"date":       {Order: 120},
			"tags":       {Order: 130},
			"categories": {Order: 140},
			"content":    {Order: 150, Hide: true},
		},
		Actions: types.Actions{
			{
				// on ne supprime que ses propres tâches
				Label:  "Recharger le répertoire",
				Plugin: fmt.Sprintf("hugoDirectoriesToSQL(%s,%s,%s)", "/home/billerot/Abri/foirexpo", "hugodoc", "foiredit"), // path,table,aliasDB
			},
		},
	},
	"vfolder": {
		Title:    "Les répertoires",
		IconName: "folder",
		Type:     "hugo",
		FormView: "fview",
		FormEdit: "fdoc",
		// OrderBy:  "path,isdir,base",
		Elements: types.Elements{
			"id":         {Order: 1},
			"path":       {Order: 10},
			"isdir":      {Order: 20},
			"level":      {Order: 30},
			"dir":        {Order: 40},
			"base":       {Order: 50},
			"ext":        {Order: 60},
			"title":      {Order: 100},
			"draft":      {Order: 110},
			"date":       {Order: 120},
			"tags":       {Order: 130},
			"categories": {Order: 140},
			"content":    {Order: 150},
		},
		Actions: types.Actions{
			{
				// on ne supprime que ses propres tâches
				Label:  "Recharger le répertoire",
				Plugin: fmt.Sprintf("hugoDirectoriesToSQL(%s,%s,%s)", "/home/billerot/Abri/foirexpo", "hugodoc", "foiredit"), // path,table,aliasDB
			},
		},
	},
}

var hugoForms = types.Forms{
	"fview": {
		Title:    "Document",
		IconName: "file alternate outline",
		Elements: types.Elements{
			"id":    {Order: 1},
			"path":  {Order: 10},
			"isdir": {Order: 20},
			"dir":   {Order: 40},
			"base":  {Order: 50},
			"_SECTION_METADATA": {
				Order:     100,
				Type:      "section",
				LabelLong: "METADATA",
				Params: types.Params{
					IconName: "list alternate outline",
				},
			},
			"title":      {Order: 110},
			"draft":      {Order: 120},
			"date":       {Order: 130},
			"tags":       {Order: 140},
			"categories": {Order: 150},
			"_SECTION_CONTENT": {
				Order:     200,
				Type:      "section",
				LabelLong: "DOCUMENT",
				Params: types.Params{
					Form:     "fdoc",
					IconName: "file code alternate outline",
				},
			},
			"content": {Order: 210},
		},
	},
	"fdoc": {
		Title:    "Document",
		IconName: "file code alternate outline",
		Elements: types.Elements{
			"_section1": {Order: 1, Type: "section"},
			"id":        {Order: 10, Grid: "two wide"},
			"path":      {Order: 40, ReadOnly: true, Grid: "seven wide"},
			"base":      {Order: 50, ReadOnly: true, Grid: "seven wide"},
			"_section2": {Order: 100, Type: "section"},
			"content":   {Order: 150, Grid: "sixteen wide"},
		},
		Actions: types.Actions{
			{Plugin: "ContentToFile({$id},hugodoc,content,{$path})"},
		},
	},
}
