package app

import "github.com/pbillerot/beedule/types"

// Orders table des Ordres d'achat ou de vente d'actions
var Orders = types.Table{
	AliasDB:    "picsou",
	Key:        "orders_id",
	ColDisplay: "orders_ptf_id",
	IconName:   "shopping cart",
	Elements:   ordersElements,
	Views:      ordersViews,
	Forms:      ordersForms,
}

var ordersViews = types.Views{
	"vachat": {
		Title:     "Achat",
		Info:      "Liste des Valeurs placées",
		IconName:  "folder open outline",
		FormAdd:   "faddbuy",
		FormEdit:  "feditbuy",
		Deletable: true,
		Hide:      false,
		Where:     "orders_order = 'buy'",
		Elements: types.Elements{
			"orders_id":     {Order: 1},
			"orders_ptf_id": {Order: 10},
			"orders_order":  {Order: 20},
			"orders_quote":  {Order: 22},
			"orders_time":   {Order: 30},
		},
	},
	"vvente": {
		Title:     "Vente",
		Info:      "Historique des placements",
		IconName:  "folder outline",
		FormAdd:   "faddbuy",
		FormEdit:  "feditbuy",
		Deletable: true,
		Where:     "orders_order = 'sell'",
		Elements: types.Elements{
			"orders_id":     {Order: 1},
			"orders_ptf_id": {Order: 10},
			"orders_order":  {Order: 20},
			"orders_quote":  {Order: 22},
			"orders_time":   {Order: 30},
		},
	},
}
var ordersForms = types.Forms{
	"faddbuy": {
		Title: "Ajout d'un ordre d'achat",
		Elements: types.Elements{
			"orders_id":     {Order: 1},
			"orders_ptf_id": {Order: 10},
			"orders_order":  {Order: 20},
			"orders_quote":  {Order: 22},
			"orders_time":   {Order: 30},
		},
	},
	"feditbuy": {
		Title: "Mise à jour d'un ordre d'achat",
		Elements: types.Elements{
			"orders_id":     {Order: 1},
			"orders_ptf_id": {Order: 10},
			"orders_order":  {Order: 20},
			"orders_quote":  {Order: 22},
			"orders_time":   {Order: 30},
		},
	},
}

var ordersElements = types.Elements{
	"orders_id": {
		Type:       "number",
		LabelLong:  "N° enregistrement",
		LabelShort: "N°",
	},
	"orders_ptf_id": {
		Type:       "combo",
		LabelLong:  "Valeur",
		LabelShort: "Valeur",
		ItemsSQL:   "SELECT ptf_id, ptf_name From ptf order by ptf_name",
	},
	"_ptf_isin": {
		Type:       "text",
		LabelLong:  "Code ISIN",
		LabelShort: "ISIN",
		ComputeSQL: "SELECT ptf_isin From ptf where ptf_id = '{orders_ptf_id}'",
		Jointure: types.Jointure{
			Join:   "LEFT OUTER JOIN ptf on ptf.ptf_id = orders.orders_ptf_id",
			Column: "ptf_isin",
		},
	},
	"orders_order": {
		Type:       "combo",
		LabelLong:  "Order",
		LabelShort: "Order",
		ColAlign:   "center",
		Items: []types.Item{
			{Key: "buy", Value: "Achat"},
			{Key: "sell", Value: "Vente"},
		},
	},
	"orders_rem": {
		Type:       "text",
		LabelLong:  "Remarque",
		LabelShort: "Rem.",
	},
	"orders_time": {
		Type:       "text",
		LabelLong:  "Jour Heure d'achat",
		LabelShort: "JH d'achat",
		DefaultSQL: "select datetime('now', 'localtime')",
	},
	"orders_sell_time": {
		Type:       "text",
		LabelLong:  "Jour Heure de vente",
		LabelShort: "JH de vente",
		DefaultSQL: "select datetime('now', 'localtime')",
	},
	"orders_quote": {
		Type:       "amount",
		LabelLong:  "Cours du jour",
		LabelShort: "Cours J",
		ComputeSQL: "select close from quotes where id = '{orders_ptf_id}' and date = (select max(date) from quotes where id = '{orders_ptf_id}')",
	},
	"orders_quantity": {
		Type:       "int",
		LabelLong:  "Quantité",
		LabelShort: "Qt",
		Refresh:    true,
	},
	"orders_buy": {
		Type:       "amount",
		LabelLong:  "Cours d'achat",
		LabelShort: "Achat à",
		Refresh:    true,
	},
	"orders_sell": {
		Type:       "amount",
		LabelLong:  "Cours de vente",
		LabelShort: "Vente à",
		Default:    "{orders_quote}",
		Refresh:    true,
	},
	"orders_cost_price": {
		Type:       "amount",
		LabelLong:  "Prix de revient",
		LabelShort: "Revient",
		ColWith:    80,
		ComputeSQL: "select ({orders_buy} * {orders_quantity} + {orders_buy} * {orders_quantity} * ({__cost}*2))/{orders_quantity}",
	},
	"orders_cost": {
		Type:       "amount",
		LabelLong:  "Frais",
		LabelShort: "Frais",
		ColWith:    60,
		ComputeSQL: "select {orders_buy} * {orders_quantity} * {__cost}",
	},
	"orders_debit": {
		Type:       "amount",
		LabelLong:  "Débit",
		LabelShort: "Débit",
		ColWith:    80,
		ComputeSQL: "select {orders_buy} * {orders_quantity} + {orders_buy} * {orders_quantity} * {__cost}",
	},
	"orders_credit": {
		Type:       "amount",
		LabelLong:  "Crédit",
		LabelShort: "Crédit",
		ColWith:    80,
		ComputeSQL: "select {orders_sell} * {orders_quantity} + {orders_sell} * {orders_quantity} * {__cost}",
	},
	"orders_gain": {
		Type:       "amount",
		LabelLong:  "Gain",
		LabelShort: "Gain",
		ColWith:    80,
		ComputeSQL: "select {orders_quote} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_quote} * {orders_quantity} * {__cost}",
		ClassSQL:   "select case when orders_gain < 0 then '#FF0000' else '#006600' end",
	},
	"orders_gainp": {
		Type:       "float",
		LabelLong:  "Gain en %",
		LabelShort: "en %",
		Format:     "%3.2f %",
		ColWith:    80,
		ComputeSQL: "select ( ({orders_quote} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_quote} * {orders_quantity} * {__cost}) / ({orders_buy} * {orders_quantity}) )*100",
		ClassSQL:   "select case when orders_gainp < 0 then '#FF0000' else '#006600' end",
	},
	"orders_sell_cost": {
		Type:       "float",
		LabelLong:  "Frais",
		LabelShort: "Frais",
		Format:     "%3.2f %",
		ColWith:    60,
		ComputeSQL: "select {orders_buy} * {orders_quantity} * {__cost} + {orders_sell} * {orders_quantity} * {__cost}",
	},
	"orders_sell_gain": {
		Type:       "float",
		LabelLong:  "Gain",
		LabelShort: "Gain",
		Format:     "%3.2f %",
		ColWith:    80,
		ComputeSQL: "select {orders_sell} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_sell} * {orders_quantity} * {__cost}",
		ClassSQL:   "select case when orders_sell_gain < 0 then '#FF0000' else '#006600' end",
	},
	"orders_sell_gainp": {
		Type:       "float",
		LabelLong:  "Gain en %",
		LabelShort: "en %",
		Format:     "%3.2f %",
		ColWith:    80,
		ComputeSQL: "select ( ({orders_sell} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_sell} * {orders_quantity} * {__cost}) / ({orders_buy} * {orders_quantity}) )*100",
		ClassSQL:   "select case when orders_sell_gainp < 0 then '#FF0000' else '#006600' end",
	},
	"_optimum": {
		Type:       "amount",
		LabelLong:  "Optimum",
		LabelShort: "Optimum",
		ComputeSQL: "select {orders_cost_price} + {orders_cost_price} * 0.03",
	},
	"_fsell": {
		Type:       "form",
		LabelLong:  "Vendre",
		LabelShort: "Vendre",
		Params: types.Params{
			Form:     "fsell",
			IconName: "media-playback-pause-symbolic",
			Action:   "create",
		},
		Args: types.Arg{
			"orders_ptf_id":     "{orders_ptf_id}",
			"orders_ptf_name":   "{orders_ptf_name}",
			"orders_quantity":   "{orders_quantity}",
			"orders_quote":      "{orders_quote}",
			"orders_buy":        "{orders_buy}",
			"orders_sell":       "{orders_quote}",
			"orders_cost_price": "{orders_cost_price}",
			"orders_order":      "sell",
		},
	},
	"_analyse": {
		Type:       "url",
		LabelLong:  "Analyse graphique",
		LabelShort: "Ana",
		Params: types.Params{
			Path:     "picsou_image.PicsouImage",
			IconName: "x-office-spreadsheet-template",
		},
		Args: types.Arg{
			"url":   "https://investir.lesechos.fr/charts/gif/{_ptf_isin}.gif",
			"title": "Analyse graphique de {orders_ptf_id} - {_ptf_name}",
		},
	},
	"_graph": {
		Type:       "plugin",
		LabelLong:  "Graph",
		LabelShort: "Dyn",
		Params: types.Params{
			Path:     "picsou_graph.PicsouGraphDay",
			IconName: "utilities-system-monitor",
		},
		Args: types.Arg{
			"ptf_id": "{orders_ptf_id}",
			"path":   "png/day/{orders_ptf_id}.png",
		},
	},
	"_graphday": {
		Type:       "plugin",
		LabelLong:  "Graph",
		LabelShort: "Day",
		Params: types.Params{
			Path:     "picsou_image.PicsouImage",
			IconName: "utilities-system-monitor",
		},
		Args: types.Arg{
			"path": "png/day/{orders_ptf_id}.png",
		},
	},
	"_image": {
		Type:       "plugin",
		LabelLong:  "Histo",
		LabelShort: "Histo",
		Params: types.Params{
			Path:     "picsou_image.PicsouImage",
			IconName: "emblem-photos",
		},
		Args: types.Arg{
			"path": "png/quotes/{orders_ptf_id}.png",
		},
	},
	"_yahoo": {
		Type:       "url",
		LabelLong:  "Lien vers Yahoo",
		LabelShort: "Yahoo",
		Params: types.Params{
			URL:      "https://fr.finance.yahoo.com/chart/{orders_ptf_id}",
			IconName: "applications-internet",
		},
	},
	"_batch": {
		Type:       "batch_sql",
		LabelLong:  "Lien vers Yahoo",
		LabelShort: "Yahoo",
		Params: types.Params{
			SQL1:     "update orders set orders_quote = (select close from quotes where id = orders_ptf_id and date = (select max(date) from quotes where id = orders_ptf_id))",
			SQL2:     "update orders set orders_gain = orders_quote * orders_quantity - orders_buy * orders_quantity - orders_buy * orders_quantity * {__cost} - orders_quote * orders_quantity * {__cost}",
			SQL3:     "update orders set orders_gainp = (orders_gain / (orders_buy * orders_quantity)) * 100",
			IconName: "view-refresh",
		},
	},
}
