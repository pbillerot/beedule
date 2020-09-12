package app

import (
	"github.com/pbillerot/beedule/types"
)

// Ptf Portefeuille des valeurs
var Ptf = types.Table{
	AliasDB:    "picsou",
	Key:        "ptf_id",
	ColDisplay: "ptf_name",
	IconName:   "building",
	Elements:   ptfElements,
	Views:      ptfViews,
	Forms:      ptfForms,
}
var ptfViews = types.Views{
	"vactiv": {
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		FormView:  "fview",
		Deletable: true,
		Title:     "Valeurs actives",
		IconName:  "building",
		Elements: types.Elements{
			"ptf_id":    {Order: 10},
			"ptf_name":  {Order: 20},
			"ptf_note":  {Order: 30},
			"ptf_top":   {Order: 40},
			"ptf_rem":   {Order: 50},
			"ptf_quote": {Order: 60},
			"ptf_gain":  {Order: 70},
		},
		Where:   "ptf_enabled = '1'",
		OrderBy: "ptf_name",
	},
	"vall": {
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		FormView:  "fview",
		Deletable: true,
		Title:     "Toutes les Valeurs",
		IconName:  "building outline",
		Elements: types.Elements{
			"ptf_id":      {Order: 10},
			"ptf_name":    {Order: 20},
			"ptf_enabled": {Order: 30},
			"ptf_note":    {Order: 40},
			"ptf_top":     {Order: 50},
			"ptf_rem":     {Order: 60},
			"ptf_quote":   {Order: 70},
			"ptf_gain":    {Order: 80},
		},
		OrderBy: "ptf_name",
	},
}

var ptfForms = types.Forms{
	"fview": {
		Title: "Fiche Valeur",
		Elements: types.Elements{
			"ptf_id":         {Order: 10},
			"ptf_name":       {Order: 20},
			"ptf_enabled":    {Order: 30},
			"ptf_note":       {Order: 40},
			"ptf_top":        {Order: 50},
			"ptf_rem":        {Order: 60},
			"ptf_quote":      {Order: 70},
			"ptf_gain":       {Order: 80},
			"_image_day":     {Order: 100},
			"_image_histo":   {Order: 110},
			"_image_analyse": {Order: 120},
		},
	},
	"fadd": {
		Title: "Ajout d'une valeur",
		Elements: types.Elements{
			"ptf_id":   {Order: 10},
			"ptf_name": {Order: 20},
			"ptf_isin": {Order: 30},
		},
	},
	"fedit": {
		Title: "Fiche Valeur",
		Elements: types.Elements{
			"ptf_id":   {Order: 10},
			"ptf_name": {Order: 20},
			"ptf_isin": {Order: 30},
		},
	},
}

var ptfElements = types.Elements{
	"ptf_id": {
		Type:       "text",
		LabelLong:  "Valeur",
		LabelShort: "Valeur",
	},
	"ptf_name": {
		Type:       "text",
		LabelLong:  "Nom",
		LabelShort: "Nom",
	},
	"ptf_isin": {
		Type:       "text",
		LabelLong:  "Code ISIN",
		LabelShort: "ISIN",
	},
	"ptf_note": {
		Type:       "text",
		LabelLong:  "Note",
		LabelShort: "Note",
	},
	"ptf_rem": {
		Type:       "text",
		LabelLong:  "Remarque",
		LabelShort: "Remarque",
	},
	"ptf_enabled": {
		Type:       "checkbox",
		LabelLong:  "Valeur Active",
		LabelShort: "Active",
	},
	"ptf_top": {
		Type:       "checkbox",
		LabelLong:  "TOP",
		LabelShort: "TOP",
	},
	"ptf_quote": {
		Type:       "amount",
		LabelLong:  "Quote du jour",
		LabelShort: "Quote",
	},
	"ptf_gain": {
		Type:       "amount",
		LabelLong:  "Gain du jour",
		LabelShort: "Gain",
		ClassSQL:   "select case when {ptf_gain} > 0 then 'green' when {ptf_gain} < 0 then 'red' end",
	},
	"_image_day": {
		Type:       "image",
		LabelLong:  "Graph du jour",
		LabelShort: "Graph J",
		Params: types.Params{
			Path:     "/crud/data/picsou/png/day/{ptf_id}.png",
			IconName: "emblem-photos",
		},
	},
	"_image_histo": {
		Type:       "image",
		LabelLong:  "Historique sur 1 mois",
		LabelShort: "Histo",
		Params: types.Params{
			Path:     "/crud/data/picsou/png/quotes/{ptf_id}.png",
			IconName: "emblem-photos",
		},
	},
	"_image_analyse": {
		Type:       "image",
		LabelLong:  "Analyse sur 7 mois",
		LabelShort: "Analyse",
		Params: types.Params{
			Path:     "/crud/data/picsou/png/ana/{ptf_id}.gif",
			IconName: "emblem-photos",
		},
	},
}
