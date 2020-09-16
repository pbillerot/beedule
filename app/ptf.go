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
		Mask: types.MaskList{
			Header: []string{
				"ptf_name",
				"ptf_id",
			},
			Meta: []string{
				"ptf_enabled",
				"ptf_top",
			},
			Description: []string{
				"ptf_rem",
			},
			Extra: []string{
				"ptf_quote",
				"ptf_gain",
			},
		},
		Elements: types.Elements{
			"ptf_id":      {},
			"ptf_name":    {},
			"ptf_enabled": {},
			"ptf_top":     {},
			"ptf_rem":     {},
			"ptf_quote":   {},
			"ptf_gain":    {},
		},
		Where:   "ptf_enabled = '1'",
		OrderBy: "ptf_name",
	},
	"vinactiv": {
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		FormView:  "fview",
		Deletable: true,
		Title:     "Les Valeurs inactives",
		IconName:  "building outline",
		Mask: types.MaskList{
			Header: []string{
				"ptf_name",
				"ptf_id",
			},
			Meta: []string{
				"ptf_enabled",
				"ptf_top",
			},
			Description: []string{
				"ptf_rem",
			},
			Extra: []string{
				"ptf_quote",
				"ptf_gain",
			},
		},
		Elements: types.Elements{
			"ptf_id":      {},
			"ptf_name":    {},
			"ptf_enabled": {},
			"ptf_top":     {},
			"ptf_rem":     {},
			"ptf_quote":   {},
			"ptf_gain":    {},
		},
		Where:   "ptf_enabled <> '1'",
		OrderBy: "ptf_name",
	},
	"vtop": {
		FormEdit:  "fedit",
		FormView:  "fview",
		Deletable: true,
		Title:     "Les Valeurs TOP",
		IconName:  "building outline",
		Mask: types.MaskList{
			Header: []string{
				"ptf_name",
				"ptf_id",
			},
			Meta: []string{
				"ptf_enabled",
				"ptf_top",
			},
			Description: []string{
				"ptf_rem",
			},
			Extra: []string{
				"ptf_quote",
				"ptf_gain",
			},
		},
		Elements: types.Elements{
			"ptf_id":      {},
			"ptf_name":    {},
			"ptf_enabled": {},
			"ptf_top":     {},
			"ptf_rem":     {},
			"ptf_quote":   {},
			"ptf_gain":    {},
		},
		Where:   "ptf_top = '1'",
		OrderBy: "ptf_name",
	},
	"vdiapo": {
		Title:    "Graphiques",
		IconName: "photo video",
		Type:     "image",
		Elements: types.Elements{
			"ptf_id":         {Order: 10},
			"ptf_name":       {Order: 20},
			"ptf_rem":        {Order: 30},
			"ptf_note":       {Order: 40},
			"_image_day":     {Order: 100},
			"_image_histo":   {Order: 110},
			"_image_analyse": {Order: 120},
		},
	},
}

var ptfForms = types.Forms{
	"fview": {
		Title: "Fiche Valeur",
		Elements: types.Elements{
			"ptf_id":         {Order: 10},
			"ptf_name":       {Order: 20},
			"ptf_enabled":    {Order: 30},
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
			"ptf_id":      {Order: 10},
			"ptf_name":    {Order: 20},
			"ptf_isin":    {Order: 30},
			"ptf_enabled": {Order: 40},
			"ptf_top":     {Order: 50},
			"ptf_rem":     {Order: 70},
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
		LabelLong:  "Note ",
		LabelShort: "Note",
	},
	"ptf_rem": {
		Type:       "text",
		LabelLong:  "Remarque",
		LabelShort: "Remarque",
		Class:      "orange",
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
		Type:       "percent",
		LabelLong:  "Gain du jour",
		LabelShort: "Gain",
		ClassSQL:   "select case when {ptf_gain} > 0 then 'green' when {ptf_gain} < 0 then 'red' end",
	},
	"_image_day": {
		Type:       "image",
		LabelLong:  "Graph du jour",
		LabelShort: "Graph J",
		Params: types.Params{
			Path:   "/crud/data/picsou/png/day/{ptf_id}.png",
			URL:    "/crud/edit/picsou/ptf/vdiapo/fedit/{ptf_id}",
			Legend: "{ptf_name} {ptf_id} - {ptf_rem}",
		},
	},
	"_image_histo": {
		Type:       "image",
		LabelLong:  "Historique sur 1 mois",
		LabelShort: "Histo",
		Params: types.Params{
			Path:   "/crud/data/picsou/png/quotes/{ptf_id}.png",
			Legend: "{ptf_name} {ptf_id} - {ptf_note}",
		},
	},
	"_image_analyse": {
		Type:       "image",
		LabelLong:  "Analyse sur 7 mois",
		LabelShort: "Analyse",
		Params: types.Params{
			Path:   "/crud/data/picsou/png/ana/{ptf_id}.gif",
			Legend: "{ptf_name} {ptf_id}",
		},
	},
}
