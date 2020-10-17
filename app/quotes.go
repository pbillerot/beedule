package app

import (
	"github.com/pbillerot/beedule/types"
)

// Quotes table des cotations
var Quotes = types.Table{
	AliasDB:    "picsou",
	Key:        "keyid",
	ColDisplay: "id",
	IconName:   "atom",
	Elements:   quotesElements,
	Views:      quotesViews,
	Forms:      quotesForms,
}

/*
CREATE TABLE IF NOT EXISTS "quotes" (
	"id"	TEXT,
	"name"	TEXT,
	"date"	TEXT,
	"open"	REAL,
	"high"	REAL,
	"low"	REAL,
	"close"	REAL,
	"close1"	REAL,
	"adjclose"	REAL,
	"volume"	INTEGER
);
*/
var quotesElements = types.Elements{
	"keyid": {
		Type:       "text",
		LabelLong:  "Clé",
		LabelShort: "clé",
		Jointure: types.Jointure{
			Join:   "",
			Column: "date || '-' || id",
		},
	},
	"id": {
		Type:       "text",
		LabelLong:  "Valeur",
		LabelShort: "Valeur",
	},
	"name": {
		Type:       "text",
		LabelLong:  "Nom",
		LabelShort: "Nom",
	},
	"date": {
		Type:       "date",
		LabelLong:  "Date",
		LabelShort: "Date",
	},
	"close1": {
		Type:       "amount",
		LabelLong:  "Close j-1",
		LabelShort: "Close j-1",
	},
	"open": {
		Type:       "amount",
		LabelLong:  "Open",
		LabelShort: "Open",
	},
	"high": {
		Type:       "amount",
		LabelLong:  "High",
		LabelShort: "High",
	},
	"low": {
		Type:       "amount",
		LabelLong:  "Low",
		LabelShort: "Low",
	},
	"close": {
		Type:       "amount",
		LabelLong:  "Quote",
		LabelShort: "Quote",
	},
	"adjclose": {
		Type:       "amount",
		LabelLong:  "Close adj",
		LabelShort: "Close adj",
	},
	"volume": {
		Type:       "number",
		LabelLong:  "Volume",
		LabelShort: "Volume",
	},
	"percent": {
		Type:       "percent",
		LabelLong:  "Gain",
		LabelShort: "Gain",
		Jointure: types.Jointure{
			Join:   "",
			Column: "((adjclose-close1)/close1)*100",
		},
		ClassSQL: "select case when {percent} > 0 then 'green' when {percent} < 0 then 'red' end",
	},
	"ptf_top": {
		Type:       "checkbox",
		LabelLong:  "TOP",
		LabelShort: "TOP",
		Jointure: types.Jointure{
			Join:   "left outer join ptf on ptf_id = id",
			Column: "ptf.ptf_top",
		},
	},
	"ptf_rem": {
		Type:       "textarea",
		LabelLong:  "Rem.",
		LabelShort: "Rem.",
		Jointure: types.Jointure{
			// Join:   "left outer join ptf on ptf_id = id",
			Column: "ptf.ptf_rem",
		},
	},
	"ptf_note": {
		Type:       "textarea",
		LabelLong:  "Note",
		LabelShort: "Note",
		Jointure: types.Jointure{
			// Join:   "left outer join ptf on ptf_id = id",
			Column: "ptf.ptf_note",
		},
	},
	"order_buy": {
		Type:       "text",
		LabelLong:  "Achat en cours",
		LabelShort: "Achat",
		Jointure: types.Jointure{
			Join:   "left outer join orders on orders_ptf_id = id and orders_order = 'buy'",
			Column: "orders_order",
		},
	},
	"_image_day": {
		Type:      "image",
		LabelLong: "Graph du jour",
		Params: types.Params{
			Path: "/crud/data/picsou/png/day/{id}.png",
		},
	},
	"_image_histo": {
		Type:      "image",
		LabelLong: "Historique sur 1 mois",
		Params: types.Params{
			Path: "/crud/data/picsou/png/quotes/{id}.png",
		},
	},
	"_image_analyse": {
		Type:      "image",
		LabelLong: "Analyse sur 7 mois",
		Params: types.Params{
			Path: "/crud/data/picsou/png/ana/{id}.gif",
		},
	},
}

var quotesViews = types.Views{
	"vlast": {
		Title:    "Cotations du jour",
		FormView: "fview",
		IconName: "atom",
		OrderBy:  "id,date",
		Group:    "picsou",
		Limit:    50,
		Where:    "date = (select max(date) from quotes)",
		Type:     "table",
		ClassSQL: "select case when '{order_buy}' = 'buy' then 'violet' else '' end",
		Elements: types.Elements{
			"keyid": {
				Order: 1,
				Hide:  true,
			},
			"id": {
				Order: 10,
			},
			"name": {
				Order:        20,
				HideOnMobile: true,
			},
			"date": {
				Order:        30,
				HideOnMobile: true,
			},
			// "close1": {
			// 	Order:        40,
			// 	HideOnMobile: true,
			// },
			// "adjclose": {
			// 	Order:        50,
			// 	HideOnMobile: true,
			// },
			// "open": {
			// 	Order:        70,
			// 	HideOnMobile: true,
			// },
			// "high": {
			// 	Order:        80,
			// 	HideOnMobile: true,
			// },
			// "low": {
			// 	Order:        90,
			// 	HideOnMobile: true,
			// },
			"close": {
				Order: 100,
			},
			"volume": {
				Order:        110,
				HideOnMobile: true,
			},
			"percent": {
				Order: 200,
			},
			"order_buy": {
				Order: 250,
				Hide:  true,
			},
			"ptf_top": {
				Order: 300,
			},
			"ptf_rem": {
				Order: 310,
			},
		},
		Actions: types.Actions{
			{
				Label: "Effacer les remarques...",
				SQL: []string{
					"update ptf set ptf_rem = ''",
				},
				WithConfirm: false,
			},
		},
	},
}

var quotesForms = types.Forms{
	"fview": {
		Title: "Cotation",
		Group: "picsou",
		Elements: types.Elements{
			"keyid":   {Order: 1},
			"name":    {Order: 10},
			"id":      {Order: 20},
			"close1":  {Order: 30},
			"close":   {Order: 40},
			"percent": {Order: 50},
			"volume":  {Order: 60},
			"_ptf": {
				Order:     100,
				Type:      "section",
				LabelLong: "Portefeuille",
				Params: types.Params{
					URL:      "/crud/edit/picsou/ptf/vall/fedit/{id}",
					IconName: "building",
				},
			},
			"ptf_top":        {Order: 110},
			"ptf_rem":        {Order: 120},
			"_image_day":     {Order: 200},
			"_image_histo":   {Order: 300},
			"_image_analyse": {Order: 400},
		},
	},
}
