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
		Info:     "Cotations du jour",
		Title:    "Cotations du jour",
		FormView: "fview",
		IconName: "atom",
		OrderBy:  "id,date",
		Group:    "picsou",
		Limit:    50,
		Where:    "date = (select max(date) from quotes)",
		Type:     "table",
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
				Order: 30,
			},
			"close1": {
				Order:        40,
				HideOnMobile: true,
			},
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
		},
	},
}

var quotesForms = types.Forms{
	"fview": {
		Title: "Cotation",
		Group: "picsou",
		Elements: types.Elements{
			"keyid":          {Order: 1},
			"name":           {Order: 10},
			"id":             {Order: 20},
			"close1":         {Order: 30},
			"adjclose":       {Order: 40},
			"percent":        {Order: 50},
			"open":           {Order: 60},
			"high":           {Order: 70},
			"low":            {Order: 80},
			"close":          {Order: 90},
			"volume":         {Order: 100},
			"_image_day":     {Order: 200},
			"_image_histo":   {Order: 210},
			"_image_analyse": {Order: 220},
		},
	},
}
