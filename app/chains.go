package app

import (
	"github.com/pbillerot/beedule/types"
)

// Chains table des groupes d'utilisateurs
var Chains = types.Table{
	AliasDB:    "pendule",
	Key:        "chain_id",
	ColDisplay: "label",
	IconName:   "cogs",
	Elements:   chainsElements,
	Views:      chainsViews,
	Forms:      chainsForms,
}

var chainsViews = types.Views{
	"vall": {
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		FormView:  "fview",
		Deletable: true,
		Title:     "Chaînes",
		IconName:  "calendar alternate outline",
		OrderBy:   "label",
		Group:     "admin",
		ClassSQL:  "select case when '{etat}' = 'RUN' then 'blue' when '{etat}' = 'OK' then 'green' when '{etat}' = 'KO' then 'red' else '' end",
		Mask: types.MaskList{
			Header: []string{
				"chain_id",
				"label",
			},
			Meta: []string{
				"active",
				"email",
			},
			Description: []string{
				"planif",
			},
			Extra: []string{
				"etat",
				"heuredebut",
				"heurefin",
				"dureemn",
			},
		},
		Elements: types.Elements{
			"chain_id":   {},
			"label":      {},
			"planif":     {},
			"email":      {},
			"active":     {},
			"etat":       {},
			"doc":        {},
			"heuredebut": {},
			"heurefin":   {},
			"dureemn":    {},
		},
		Actions: types.Actions{
			{
				Label:  "Démarrer / Arrêter le pendule",
				Plugin: "StartStopPendule()",
				Checkbox: types.Setters{
					GetSQL:  "select value from parameters where id = 'batch_etat'",
					AliasDB: "admin",
				},
			},
		},
	},
}

var chainsForms = types.Forms{
	"fview": {
		Title:    "Pendule Job",
		Group:    "admin",
		IconName: "calendar alternate outline",
		Elements: types.Elements{
			"chain_id": {Order: 01},
			"label":    {Order: 10},
			"planif":   {Order: 20},
			"email":    {Order: 30},
			"active":   {Order: 40},
			"_SECTION_RESULTAT": {
				Order:     500,
				Type:      "section",
				LabelLong: "RESULTATS",
				Params: types.Params{
					IconName: "file alternate outline",
				},
			},
			"etat":       {Order: 510},
			"heuredebut": {Order: 530},
			"heurefin":   {Order: 540},
			"dureemn":    {Order: 550},
		},
	},
	"fadd": {
		Title:    "Ajout d'une Chaîne",
		Group:    "admin",
		IconName: "calendar alternate outline",
		Elements: types.Elements{
			"chain_id": {Order: 01},
			"label":    {Order: 10},
			"email":    {Order: 20},
			"planif":   {Order: 30},
			"active":   {Order: 40},
		},
	},
	"fedit": {
		Title:    "Chaîne",
		Group:    "admin",
		IconName: "calendar alternate outline",
		Elements: types.Elements{
			"chain_id": {Order: 01},
			"label":    {Order: 10},
			"email":    {Order: 20},
			"planif":   {Order: 30},
			"active":   {Order: 40},
		},
	},
}

var chainsElements = types.Elements{
	"chain_id": { // n° du job
		Type:       "counter",
		LabelLong:  "n°",
		LabelShort: "n°",
	},
	"label": { // Nom de la chaîne
		Type:       "text",
		LabelLong:  "Nom",
		LabelShort: "Nom",
		Class:      "violet",
	},
	"active": { // Chaîne active
		Type:       "checkbox",
		LabelLong:  "Active",
		LabelShort: "Active",
	},
	"planif": { // planification Cron de la chaîne sur la séquence 0
		Type:       "text",
		LabelLong:  "Planification",
		LabelShort: "Planification",
		Class:      "orange",
	},
	"email": { // Email en cas d'erreur
		Type:       "email",
		LabelLong:  "Email si erreur",
		LabelShort: "Email si erreur",
	},
	"etat": { // État de la chaîne
		Type:       "text", // INI, OK, KO, RUN
		LabelLong:  "État",
		LabelShort: "État",
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
