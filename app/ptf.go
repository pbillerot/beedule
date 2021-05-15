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
	"vall": {
		FormEdit:  "fedit",
		FormView:  "fview",
		Group:     "picsou",
		Deletable: true,
		Title:     "Les Valeurs",
		IconName:  "city",
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
		Elements: map[string]types.Element{
			"ptf_id":      {},
			"ptf_name":    {},
			"ptf_enabled": {},
			"ptf_top":     {},
			"ptf_rem":     {},
			"ptf_quote":   {},
			"ptf_gain":    {},
		},
		// Where:   "ptf_enabled = '1' and ptf_top = '1'",
		OrderBy: "ptf_name",
		Actions: []types.Action{
			{
				Label: "Effacer les remarques...",
				SQL: []string{
					"update ptf set ptf_rem = ''",
				},
				WithConfirm: true,
			},
		},
	},
	"vdiapo": {
		Title:    "Graphiques",
		FormEdit: "fedit",
		FormView: "fview",
		Group:    "picsou",
		IconName: "photo video",
		Type:     "image",
		ClassSQL: "select case when '{ptf_note}' like '%achat%' then 'crud-gondole' else '' end",
		Elements: map[string]types.Element{
			"ptf_id":         {Order: 10},
			"ptf_name":       {Order: 20},
			"ptf_top":        {Order: 25},
			"ptf_rem":        {Order: 30},
			"ptf_note":       {Order: 40},
			"ptf_gain":       {Order: 50},
			"ptf_quote":      {Order: 60},
			"_image_analyse": {Order: 120},
			// "_chart_quotes":  {Order: 130},
		},
		OrderBy: "ptf_name",
		Where:   "ptf_enabled = '1'",
		Actions: []types.Action{
			{
				Label: "Effacer les remarques...",
				SQL: []string{
					"update ptf set ptf_rem = ''",
				},
				WithConfirm: true,
			},
		},
	},
	"vtop": {
		Title:    "Graphiques TOP",
		FormEdit: "fedit",
		FormView: "fview",
		Group:    "picsou",
		IconName: "photo video",
		Type:     "image",
		ClassSQL: "select case when '{ptf_note}' like '%achat%' then 'crud-gondole' else '' end",
		Elements: map[string]types.Element{
			"ptf_id":         {Order: 10},
			"ptf_name":       {Order: 20},
			"ptf_top":        {Order: 25},
			"ptf_rem":        {Order: 30},
			"ptf_note":       {Order: 40},
			"ptf_gain":       {Order: 50},
			"ptf_quote":      {Order: 60},
			"_image_analyse": {Order: 120},
			// "_chart_quotes":  {Order: 130},
		},
		OrderBy: "ptf_name",
		Where:   "ptf_enabled = '1' and ptf_top = '1'",
		Actions: []types.Action{
			{
				Label: "Effacer les remarques...",
				SQL: []string{
					"update ptf set ptf_rem = ''",
				},
				WithConfirm: true,
			},
		},
	},
	"vntop": {
		Title:    "Graphiques NON TOP",
		FormEdit: "fedit",
		FormView: "fview",
		Group:    "picsou",
		IconName: "photo video",
		Type:     "image",
		ClassSQL: "select case when '{ptf_note}' like '%achat%' then 'crud-gondole' else '' end",
		Elements: map[string]types.Element{
			"ptf_id":         {Order: 10},
			"ptf_name":       {Order: 20},
			"ptf_top":        {Order: 25},
			"ptf_rem":        {Order: 30},
			"ptf_note":       {Order: 40},
			"ptf_gain":       {Order: 50},
			"ptf_quote":      {Order: 60},
			"_image_analyse": {Order: 120},
			// "_chart_quotes":  {Order: 130},
		},
		OrderBy: "ptf_name",
		Where:   "ptf_enabled = '1' and ptf_top <> '1'",
		Actions: []types.Action{
			{
				Label: "Effacer les remarques...",
				SQL: []string{
					"update ptf set ptf_rem = ''",
				},
				WithConfirm: true,
			},
		},
	},
}

var ptfForms = types.Forms{
	"fview": {
		Title: "Fiche Valeur",
		Group: "picsou",
		Elements: map[string]types.Element{
			"ptf_id":         {Order: 10},
			"ptf_name":       {Order: 20},
			"ptf_enabled":    {Order: 30},
			"ptf_top":        {Order: 50},
			"ptf_note":       {Order: 55},
			"ptf_rem":        {Order: 60},
			"ptf_quote":      {Order: 70},
			"ptf_gain":       {Order: 80},
			"_action_buy":    {Order: 90},
			"_chart_quotes":  {Order: 110},
			"_image_analyse": {Order: 120},
		},
	},
	"fadd": {
		Title: "Ajout d'une valeur",
		Group: "picsou",
		Elements: map[string]types.Element{
			"ptf_id":   {Order: 10},
			"ptf_name": {Order: 20},
			"ptf_isin": {Order: 30},
		},
	},
	"fedit": {
		Title: "Fiche Valeur",
		Group: "picsou",
		Elements: map[string]types.Element{
			"ptf_id":      {Order: 10},
			"ptf_name":    {Order: 20},
			"ptf_isin":    {Order: 30},
			"ptf_enabled": {Order: 40},
			"ptf_top":     {Order: 50},
			"ptf_note":    {Order: 60, Protected: true},
			"ptf_rem":     {Order: 70},
		},
		PostSQL: []string{
			"UPDATE PTF set ptf_note = 'TOP' where ptf_note = '' and ptf_top = '1' and ptf_id = '{ptf_id}'",
			"UPDATE PTF set ptf_note = '' where ptf_note = 'TOP' and ptf_top = '0' and ptf_id = '{ptf_id}'",
		},
	},
}

var ptfElements = map[string]types.Element{
	"_action_buy": {
		Type:      "action",
		LabelLong: "Acheter la valeur...",
		Group:     "trader",
		Actions: []types.Action{
			{
				Label: "Acheter cette valeur",
				URL:   "/bee/add/picsou/orders/vachat/feditbuy?orders_order=buy&orders_ptf_id={ptf_id}&orders_quote={ptf_quote}&orders_buy={ptf_quote}",
			},
		},
	},

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
		ClassSQL:   "select 'orange'",
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
			Path: "/bee/data/picsou/png/day/{ptf_id}.png",
			URL:  "/bee/data/picsou/png/day/{ptf_id}.png",
			Header: []string{
				"ptf_name",
				"ptf_id",
			},
			Description: []string{
				"ptf_rem",
			},
			Extra: []string{
				"ptf_top",
				"ptf_quote",
				"ptf_gain",
			},
		},
	},
	"_image_histo": {
		Type:       "image",
		LabelLong:  "Historique sur 1 mois",
		LabelShort: "Histo",
		Params: types.Params{
			Path: "/bee/data/picsou/png/quotes/{ptf_id}.png",
			URL:  "/bee/data/picsou/png/quotes/{ptf_id}.png",
			Header: []string{
				"ptf_name",
				"ptf_id",
			},
			Extra: []string{
				"ptf_top",
				"ptf_quote",
				"ptf_gain",
			},
		},
	},
	"_image_analyse": {
		Type:       "image",
		LabelLong:  "Analyse sur 7 mois",
		LabelShort: "Analyse",
		Params: types.Params{
			Path: "/bee/data/picsou/png/ana/{ptf_id}.gif",
			URL:  "/bee/data/picsou/png/ana/{ptf_id}.gif",
			Header: []string{
				"ptf_name",
				"ptf_id",
			},
			Extra: []string{
				"ptf_top",
				"ptf_quote",
				"ptf_gain",
			},
		},
	},
	"_chart_quotes": {
		Type:       "image",
		LabelLong:  "Cotation sur 1 mois",
		LabelShort: "Cotation",
		Dataset: map[string]string{
			"ClassJquery": "select 'bee-chart-quotes'",
			"Title":       "select 'Cours de {ptf_id}'",
			"Quotes":      "select open as matin, close as soir from quotes where id = '{ptf_id}' order by date",
			"Quotep":      "select (open-close1)*100/close1 as matin, (close-close1)*100/close1 as soir from quotes where id = '{ptf_id}' order by date",
			"Labels":      "select printf('%s-%s',substr(date,9,2),substr(date,6,2)) as 'matin', '-' as 'soir' from quotes where id = '{ptf_id}' order by date",
			"Minp":        "select (low-close1)*100/close1 as matin, (low-close1)*100/close1 as soir from quotes where id = '{ptf_id}' order by date",
			"Maxp":        "select (high-close1)*100/close1 as matin, (high-close1)*100/close1 as soir from quotes where id = '{ptf_id}' order by date",
			"Min":         "select min(low) as min from quotes where id = '{ptf_id}'",
			"Max":         "select max(high) as max from quotes where id = '{ptf_id}'",
			"SeuilV":      "select orders_cost_price + orders_cost_price * {__optimum} from orders where orders_ptf_id = '{ptf_id}' and orders_order = 'buy'",
			"SeuilR":      "select orders_cost_price from orders where orders_ptf_id = '{ptf_id}' and orders_order = 'buy'",
		},
	},
}
