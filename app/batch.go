package app

import (
	"github.com/pbillerot/beedule/types"
)

// Batch table des groupes d'utilisateurs
var Batch = types.Table{
	AliasDB:    "pendule",
	Key:        "id",
	ColDisplay: "label",
	IconName:   "cogs",
	Elements:   batchElements,
	Views:      batchViews,
	Forms:      batchForms,
}

var batchViews = types.Views{
	"vall": {
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		FormView:  "fview",
		Deletable: true,
		Title:     "Pendule",
		IconName:  "cogs",
		OrderBy:   "chaine, sequence",
		Group:     "admin",
		Mask: types.MaskList{
			Header: []string{
				"id",
				"chaine",
				"label",
			},
			Meta: []string{
				"sequence",
				"heuredebut",
				"heurefin",
				"dureemn",
			},
			Description: []string{
				"planif",
			},
			Extra: []string{
				"type",
				"etat",
				"sierreur",
			},
		},
		Elements: types.Elements{
			"id":          {},
			"label":       {},
			"chaine":      {},
			"planif":      {},
			"sequence":    {},
			"sierreur":    {},
			"actif":       {},
			"etat":        {},
			"type":        {},
			"commandes":   {},
			"options":     {},
			"result":      {},
			"nexec":       {},
			"heureresult": {},
			"doc":         {},
			"email":       {},
			"heuredebut":  {},
			"heurefin":    {},
			"dureemn":     {},
		},
	},
}

var batchForms = types.Forms{
	"fview": {
		Title: "Pendule Job",
		Group: "admin",
		Elements: types.Elements{
			"id": {Order: 01, Hide: true},
			"_SECTION_CHAINE": {
				Order:     100,
				Type:      "section",
				LabelLong: "CHAÎNE",
				Params: types.Params{
					Form:     "fchaine",
					IconName: "calendar alternate outline",
				},
			},
			"chaine":     {Order: 120},
			"planif":     {Order: 130},
			"sequence":   {Order: 140},
			"sierreur":   {Order: 150},
			"email":      {Order: 155},
			"nexec":      {Order: 160},
			"heuredebut": {Order: 170},
			"heurefin":   {Order: 180},
			"dureemn":    {Order: 190},
			"_SECTION_JOB": {
				Order:     300,
				Type:      "section",
				LabelLong: "JOB",
				Params: types.Params{
					Form:     "fjob",
					IconName: "cog",
				},
			},
			"actif":     {Order: 310},
			"label":     {Order: 320},
			"type":      {Order: 330},
			"commandes": {Order: 340},
			"_SECTION_RESULTAT": {
				Order:     500,
				Type:      "section",
				LabelLong: "RESULTATS",
				Params: types.Params{
					IconName: "file alternate outline",
				},
			},
			"etat":        {Order: 510},
			"heureresult": {Order: 520},
			"result":      {Order: 530},
		},
	},
	"fadd": {
		Title: "Pendule ajout d'un job",
		Group: "admin",
		Elements: types.Elements{
			"id":       {Order: 01},
			"chaine":   {Order: 10},
			"sequence": {Order: 20},
			"label":    {Order: 30},
		},
	},
	"fchaine": {
		Title:    "Chaîne",
		Group:    "admin",
		IconName: "calendar alternate outline",
		Elements: types.Elements{
			"id":     {Order: 01},
			"chaine": {Order: 10},
			"label":  {Order: 20},
			"planif": {Order: 30},
		},
	},
	"fjob": {
		Title:    "Job",
		Group:    "admin",
		IconName: "cog",
		Elements: types.Elements{
			"id":        {Order: 01, Hide: true},
			"actif":     {Order: 10},
			"chaine":    {Order: 20},
			"sequence":  {Order: 30},
			"sierreur":  {Order: 40},
			"label":     {Order: 50},
			"type":      {Order: 60},
			"commandes": {Order: 70},
		},
	},
}

var batchElements = types.Elements{
	"id": { // n° du job
		Type:       "counter",
		LabelLong:  "n°",
		LabelShort: "n°",
	},
	"label": { // Nom du job dans la chaîne
		Type:       "text",
		LabelLong:  "Nom",
		LabelShort: "Nom",
		Class:      "blue",
	},
	"actif": { // Job actif
		Type:       "checkbox",
		LabelLong:  "Actif",
		LabelShort: "Actif",
	},
	"etat": { // État du job
		Type:       "text", // INI, OK, KO, RUN
		LabelLong:  "État",
		LabelShort: "État",
	},
	"chaine": { // Nom de la chaîne
		Type:       "text",
		LabelLong:  "Chaîne",
		LabelShort: "Chaîne",
		Class:      "green",
	},
	"planif": { // planification Cron de la chaîne sur la séquence 0
		Type:       "text",
		LabelLong:  "Planification",
		LabelShort: "Planification",
		Class:      "orange",
	},
	"sequence": { // n° de séquence dans la chaîne (on commence à 0)
		Type:       "number",
		LabelLong:  "Séquence",
		LabelShort: "Séquence",
	},
	"sierreur": { // n° de séquence dans la chaîne (on commence à 0)
		Type:       "combobox",
		LabelLong:  "Si erreur",
		LabelShort: "Si erreur",
		Items: []types.Item{
			{Key: "1", Label: "Job Suivant"},
			{Key: "2", Label: "Étape suivante"},
			{Key: "3", Label: "Arrêt de la chaîne"},
		},
	},
	"type": { // Type de Job
		Type:       "combobox",
		LabelLong:  "Type",
		LabelShort: "Type",
		Items: []types.Item{
			{Key: "debut", Label: "Début"},
			{Key: "boucle", Label: "Boucle"},
			{Key: "conditionfichier", Label: "Condition Fichier"},
			{Key: "conditionsql", Label: "Condition SQL"},
			{Key: "etape", Label: "Étape"},
			{Key: "plugin", Label: "Plugin"},
			{Key: "sql", Label: "SQL"},
			{Key: "shell", Label: "SHELL"},
			{Key: "variables", Label: "Variables"},
		},
	},
	"commandes": { // Commandes à exécuter
		Type:       "textarea",
		LabelLong:  "Commandes",
		LabelShort: "Commandes",
	},
	"options": { // Options du job
		Type:       "textarea",
		LabelLong:  "Options",
		LabelShort: "Options",
	},
	"nexec": { // N° du job en cours (sera enregistré dans l'étape 0)
		Type:       "text",
		LabelLong:  "Job en cours",
		LabelShort: "n°Run",
	},
	"result": { // Résultat du job
		Type:       "text",
		LabelLong:  "Résultat",
		LabelShort: "Résultat",
	},
	"email": { // Email en cas d'erreur
		Type:       "email",
		LabelLong:  "Email si erreur",
		LabelShort: "Email si erreur",
	},
	"heureresult": { // Heure du résulat du job
		Type:       "datetime",
		LabelLong:  "Heure résultat",
		LabelShort: "Heure résultat",
	},
	"heuredebut": { // Heure début de la chaîne
		Type:       "datetime",
		LabelLong:  "Démarrée le",
		LabelShort: "Démarrée le",
	},
	"heurefin": { // Heure de fin de la chaîne
		Type:       "datetime",
		LabelLong:  "Terminée le",
		LabelShort: "Terminée le",
	},
	"dureemn": { // Durée d'exécution de la chaîne
		Type:       "minute",
		LabelLong:  "Durée",
		LabelShort: "Durée",
	},
}
