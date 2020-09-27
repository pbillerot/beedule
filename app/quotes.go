package app

import (
	"github.com/pbillerot/beedule/types"
)

// Quotes table des cotations
var Quotes = types.Table{
	AliasDB:    "picsou",
	Key:        "id",
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
	"id": {
		Type:       "text",
		LabelLong:  "Valeur",
		LabelShort: "Valeur",
	},
	"name": {
		Type:       "text",
		LabelLong:  "Valeur",
		LabelShort: "Valeur",
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
		LabelLong:  "Close",
		LabelShort: "Close",
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
}

var quotesViews = types.Views{
	"vlast": {
		Info:     "Cotations du jour",
		Title:    "Cotations du jour",
		IconName: "atom",
		OrderBy:  "id,date",
		Group:    "picsou",
		Limit:    50,
		Where:    "date = (select max(date) from quotes)",
		Type:     "table",
		Elements: types.Elements{
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
				Order: 40,
			},
			"adjclose": {
				Order:        50,
				HideOnMobile: true,
			},
			"open": {
				Order:        70,
				HideOnMobile: true,
			},
			"high": {
				Order:        80,
				HideOnMobile: true,
			},
			"low": {
				Order:        90,
				HideOnMobile: true,
			},
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

var quotesForms = types.Forms{}
