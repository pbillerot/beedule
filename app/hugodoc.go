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
		LabelLong:  "Nom",
		LabelShort: "Nom",
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
		Type:       "markdown",
		LabelLong:  "Contenu",
		LabelShort: "Contenu",
		PostAction: types.Actions{
			{
				Label: "Enregistrement du contenu dans un fichier",
				// aliasDB, tableName, keyid, keyValue, columnName, pathFile
				Plugin: fmt.Sprintf("contentSQLToFile(%s,%s,%s,%s,%s,%s)", "{$aliasdb}", "{$table}", "{$key}", "{id}", "content", "{$datadir}/content{path}"),
			},
			{
				Label: "Recharger le répertoire",
				// path, table, aliasDB
				Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
			},
		},
	},
	"_image": {
		Type:       "image",
		LabelLong:  "Image",
		LabelShort: "Image",
		Params: types.Params{
			URL:  "{$dataurl}/content{path}",
			Path: "{$datadir}/content{path}",
			Form: "fimage",
		},
		HideSQL: "select 'hide' where '{ext}' not in ('.png','.jpg')",
	},
	"_pdf": {
		Type:       "pdf",
		LabelLong:  "Visualiser ou Télécharger le PDF",
		LabelShort: "PDF",
		Params: types.Params{
			URL:  "{$dataurl}/content{path}",
			Path: "/crud/static/img/pdf.png",
		},
		HideSQL: "select 'hide' where '{ext}' not in ('.pdf')",
	},
}

var hugoViews = types.Views{
	"vall": {
		Title:    "La table brute",
		IconName: "list",
		Type:     "table",
		FormView: "fview",
		FormEdit: "fdoc",
		OrderBy:  "dir,isdir desc,id",
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
				Label: "Recharger le répertoire",
				// path,table,aliasDB
				Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
			},
		},
	},
	"vfolder": {
		Title:    "Les répertoires de foirexpo",
		IconName: "folder",
		Type:     "hugo",
		FormView: "fview",
		FormEdit: "fdoc",
		OrderBy:  "dir,isdir desc,id",
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
				Label: "Recharger le répertoire",
				// path,table,aliasDB
				Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
			},
		},
	},
}

var hugoForms = types.Forms{
	"fview": {
		Title:    "Document",
		IconName: "file alternate outline",
		Elements: types.Elements{
			"_SECTION_DIRECTORY": {
				Order:     1,
				Type:      "section",
				LabelLong: "REPERTOIRE ou FICHIER",
				Params: types.Params{
					Form:     "ffile",
					IconName: "folder alternate outline",
				},
			},
			"id":     {Order: 5},
			"path":   {Order: 10},
			"isdir":  {Order: 20, Hide: true},
			"dir":    {Order: 40, Hide: true},
			"base":   {Order: 50},
			"ext":    {Order: 60, Hide: true},
			"_image": {Order: 100},
			"_pdf":   {Order: 210},
			"_SECTION_CONTENT": {
				Order:     300,
				Type:      "section",
				LabelLong: "DOCUMENT",
				Class:     "crud-card-view2",
				Params: types.Params{
					Form:     "fdoc",
					IconName: "file code alternate outline",
				},
				HideSQL: "select 'hide' where '{ext}' <> '.md' ",
			},
			"content": {Order: 310},
		},
	},
	"ffile": {
		Title:      "Gestionnaire de Fichier",
		IconName:   "folder alternate outline",
		HideSubmit: true,
		Elements: types.Elements{
			"_section0": {Order: 1, Type: "section"},
			"id":        {Order: 10, Grid: "two wide"},
			"path":      {Order: 14, Grid: "fourteen wide", ReadOnly: true},
			"isdir":     {Order: 20, Hide: true},
			"dir":       {Order: 40, Hide: true},
			"base":      {Order: 50, Hide: true},
			"ext":       {Order: 60, Hide: true},
			"_section1": {Order: 100, Type: "section"},
			"_rename_file": {
				Order:     110,
				Grid:      "ten wide",
				LabelLong: "Renommer ou déplacer le fichier...",
				Type:      "action",
				HideSQL:   "select case when '{isdir}' = '1' then 'hide' else '' end",
				Params: types.Params{
					WithConfirm: true,
					WithInput:   true,
				},
				Default: "{path}",
				Actions: types.Actions{
					{
						Plugin: "renameFile({$datadir}/content{path},{$datadir}/content{_rename_file})",
					},
					{
						Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
					},
				},
			},
			"_section12": {Order: 120, Type: "section"},
			"_copy_file": {
				Order:     121,
				Grid:      "ten wide",
				LabelLong: "Recopier le fichier...",
				Type:      "action",
				HideSQL:   "select case when '{isdir}' = '1' then 'hide' else '' end",
				Params: types.Params{
					WithConfirm: true,
					WithInput:   true,
				},
				Default: "{path}",
				Actions: types.Actions{
					{
						Plugin: "copyFile({$datadir}/content{path},{$datadir}/content{_copy_file})",
					},
					{
						Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
					},
				},
			},
			"_section2": {Order: 200, Type: "section"},
			"_rename_dir": {
				Order:     210,
				Grid:      "ten wide",
				LabelLong: "Renommer ou déplacer le répertoire...",
				Type:      "action",
				HideSQL:   "select case when '{isdir}' = '0' then 'hide' else '' end",
				Params: types.Params{
					WithConfirm: true,
					WithInput:   true,
				},
				Default: "{path}",
				Actions: types.Actions{
					{
						Plugin: "renameDirectory({$datadir}/content{path},{$datadir}/content{_rename_dir})",
					},
					{
						Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
					},
				},
			},
			"_section3": {Order: 300, Type: "section"},
			"_delete_fichier": {
				Order:     310,
				Grid:      "ten wide",
				LabelLong: "Supprimer le fichier...",
				Type:      "action",
				HideSQL:   "select case when '{isdir}' = '1' then 'hide' else '' end",
				Params: types.Params{
					WithConfirm: true,
				},
				Actions: types.Actions{
					{
						Plugin: "deleteFile({$datadir}/content{path})",
					},
					{
						Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
					},
				},
			},
			"_section4": {Order: 400, Type: "section"},
			"_create_dir": {
				Order:     410,
				Grid:      "ten wide",
				LabelLong: "Créer le répertoire...",
				Type:      "action",
				HideSQL:   "select case when '{isdir}' = '0' then 'hide' else '' end",
				Default:   "{path}",
				Params: types.Params{
					WithConfirm: true,
					WithInput:   true,
				},
				Actions: types.Actions{
					{
						Plugin: "createDirectory({$datadir}/content{_create_dir})",
					},
					{
						Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
					},
				},
			},
			"_section5": {Order: 500, Type: "section"},
			"_delete_dir": {
				Order:     510,
				Grid:      "ten wide",
				LabelLong: "Supprimer le répertoire...",
				Type:      "action",
				HideSQL:   "select case when '{isdir}' = '0' then 'hide' else '' end",
				Params: types.Params{
					WithConfirm: true,
				},
				Actions: types.Actions{
					{
						Plugin: "deleteDir({$datadir}/content{path})",
					},
					{
						Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
					},
				},
			},
			"_section6": {Order: 600, Type: "section"},
			"_upload_file": {
				Order:     610,
				Grid:      "ten wide",
				LabelLong: "Charger un fichier...",
				Type:      "action",
				HideSQL:   "select case when '{isdir}' = '0' then 'hide' else '' end",
				Params: types.Params{
					WithConfirm:   true,
					WithInputFile: true,
					Path:          "{$datadir}/content{path}",
				},
				Actions: types.Actions{
					{
						Plugin: fmt.Sprintf("hugoDirectoryToSQL(%s,%s,%s)", "{$datadir}", "{$table}", "{$aliasdb}"),
					},
				},
			},
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
			{
				Plugin: "ContentToFile({$id},hugodoc,content,{$path})",
			},
		},
	},
	"fimage": {
		Title:    "Editeur d'image",
		IconName: "image outline",
		Elements: types.Elements{
			"_section1": {Order: 1, Type: "section"},
			"id":        {Order: 10, Grid: "two wide"},
			"path":      {Order: 40, ReadOnly: true, Grid: "seven wide"},
			"base":      {Order: 50, ReadOnly: true, Grid: "seven wide"},
			"_section2": {Order: 100, Type: "section"},
			"_image":    {Order: 150, Grid: "sixteen wide"},
		},
		Actions: types.Actions{
			{
				Plugin: "ContentToFile({$id},hugodoc,content,{$path})",
			},
		},
	},
}
