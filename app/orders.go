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
		FormView:  "fviewbuy",
		Deletable: true,
		Hide:      false,
		Where:     "orders_order = 'buy'",
		ClassSQL:  "select case when {orders_cost_price} + {orders_cost_price} * {__optimum} < {orders_quote} then 'positive' when {orders_cost_price} < {orders_quote} then 'blue' else 'negative' end",
		Elements: types.Elements{
			"orders_id":     {Order: 1},
			"orders_ptf_id": {Order: 10},
			// "orders_order":  {Order: 20},
			// "orders_time":       {Order: 30},
			"orders_buy":        {Order: 40},
			"orders_cost":       {Order: 50, HideOnMobile: true},
			"orders_quantity":   {Order: 60},
			"orders_debit":      {Order: 70, HideOnMobile: true},
			"orders_cost_price": {Order: 80},
			"orders_optimum":    {Order: 100},
			"orders_quote":      {Order: 110},
			"orders_gain":       {Order: 120},
			"orders_gainp":      {Order: 130, HideOnMobile: true},
			"orders_rem":        {Order: 140, HideOnMobile: true},
		},
		ActionsSQL: types.Actions{
			{
				Label: "Mettre à jour avec le cours du jour",
				SQL: []string{
					"update orders set orders_quote = (select close from quotes where id = orders_ptf_id and date = (select max(date) from quotes where id = orders_ptf_id))",
					"update orders set orders_gain = orders_quote * orders_quantity - orders_buy * orders_quantity - orders_buy * orders_quantity * {__cost} - orders_quote * orders_quantity * {__cost}",
					"update orders set orders_gainp = (orders_gain / (orders_buy * orders_quantity)) * 100",
				},
				WithConfirm: false,
			},
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
			"orders_id":         {Order: 1},
			"orders_ptf_id":     {Order: 10},
			"orders_time":       {Order: 30},
			"orders_sell_time":  {Order: 40},
			"orders_buy":        {Order: 60},
			"orders_sell":       {Order: 70},
			"orders_quantity":   {Order: 80},
			"orders_sell_cost":  {Order: 90},
			"orders_credit":     {Order: 100},
			"orders_sell_gain":  {Order: 110},
			"orders_sell_gainp": {Order: 120, HideOnMobile: true},
		},
	},
}
var ordersForms = types.Forms{
	"fviewbuy": {
		Title: "Ordre d'achat",
		Elements: types.Elements{
			"orders_id":                {Order: 1},
			"orders_ptf_id":            {Order: 10},
			"orders_order":             {Order: 20},
			"orders_time":              {Order: 30},
			"orders_buy":               {Order: 50},
			"orders_cost":              {Order: 60},
			"orders_quantity":          {Order: 40},
			"orders_debit":             {Order: 80},
			"_section_achat_situation": {Order: 100},
			"orders_cost_price":        {Order: 110},
			"orders_optimum":           {Order: 120},
			"orders_quote":             {Order: 130},
			"orders_gain":              {Order: 140},
			"orders_gainp":             {Order: 150},
			"orders_rem":               {Order: 160},
			"_image_day":               {Order: 170},
			"_image_histo":             {Order: 180},
			"_image_analyse":           {Order: 190},
		},
	},
	"faddbuy": {
		Title: "Ajout d'un ordre d'achat",
		Elements: types.Elements{
			"orders_id":       {Order: 1},
			"orders_ptf_id":   {Order: 10},
			"orders_order":    {Order: 20, Default: "buy"},
			"orders_time":     {Order: 30},
			"orders_quote":    {Order: 40},
			"orders_buy":      {Order: 50},
			"orders_quantity": {Order: 60},
			"orders_debit":    {Order: 70},
		},
	},
	"feditbuy": {
		Title: "Mise à jour d'un ordre d'achat",
		Elements: types.Elements{
			"orders_id":       {Order: 1},
			"orders_ptf_id":   {Order: 10},
			"orders_order":    {Order: 20, Default: "buy"},
			"orders_time":     {Order: 30},
			"orders_quote":    {Order: 40},
			"orders_buy":      {Order: 50},
			"orders_quantity": {Order: 60},
			"orders_debit":    {Order: 70},
		},
	},
}

var ordersElements = types.Elements{
	"_section_achat_situation": {
		Type:       "section",
		LabelLong:  "Situation",
		LabelShort: "Situation",
		Params: types.Params{
			Form:     "feditbuy",
			IconName: "folder open outline",
		},
	},
	"orders_id": {
		Type:       "number",
		LabelLong:  "N° enregistrement",
		LabelShort: "N°",
		ColAlign:   "center",
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
		// ComputeSQL: "SELECT ptf_isin From ptf where ptf_id = '{orders_ptf_id}'",
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
		Protected:  true,
	},
	"orders_quantity": {
		Type:       "number",
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
		Protected:  true,
		// ComputeSQL: "({orders_buy} * {orders_quantity} + {orders_buy} * {orders_quantity} * ({__cost}*2))/{orders_quantity}",
	},
	"orders_cost": {
		Type:       "amount",
		LabelLong:  "Frais",
		LabelShort: "Frais",
		ColWith:    60,
		Protected:  true,
		// ComputeSQL: "orders_buy * orders_quantity * {__cost}",
	},
	"orders_debit": {
		Type:       "amount",
		LabelLong:  "Débit",
		LabelShort: "Débit",
		ColWith:    80,
		Protected:  true,
		// ComputeSQL: "select {orders_buy} * {orders_quantity} + {orders_buy} * {orders_quantity} * {__cost}",
	},
	"orders_credit": {
		Type:       "amount",
		LabelLong:  "Crédit",
		LabelShort: "Crédit",
		ColWith:    80,
		Protected:  true,
		// ComputeSQL: "select {orders_sell} * {orders_quantity} + {orders_sell} * {orders_quantity} * {__cost}",
	},
	"orders_optimum": {
		Type:       "amount",
		LabelLong:  "Optimum",
		LabelShort: "Optimum",
		Protected:  true,
		Jointure: types.Jointure{
			Column: "orders_cost_price + orders_cost_price * {__optimum}",
		},
		// ComputeSQL: "select {orders_cost_price} + {orders_cost_price} * {__optimum}",
	},
	"orders_gain": {
		Type:       "amount",
		LabelLong:  "Gain",
		LabelShort: "Gain",
		ColWith:    80,
		Protected:  true,
		// ComputeSQL: "select {orders_quote} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_quote} * {orders_quantity} * {__cost}",
	},
	"orders_gainp": {
		Type:       "percent",
		LabelLong:  "Gain en %",
		LabelShort: "en %",
		ColWith:    80,
		Protected:  true,
		// ComputeSQL: "select ( ({orders_quote} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_quote} * {orders_quantity} * {__cost}) / ({orders_buy} * {orders_quantity}) )*100",
	},
	"orders_sell_cost": {
		Type:       "float",
		LabelLong:  "Frais",
		LabelShort: "Frais",
		Format:     "%3.2f %",
		ColWith:    60,
		Protected:  true,
		// ComputeSQL: "select {orders_buy} * {orders_quantity} * {__cost} + {orders_sell} * {orders_quantity} * {__cost}",
	},
	"orders_sell_gain": {
		Type:       "float",
		LabelLong:  "Gain",
		LabelShort: "Gain",
		Format:     "%3.2f %",
		ColWith:    80,
		Protected:  true,
		// ComputeSQL: "select {orders_sell} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_sell} * {orders_quantity} * {__cost}",
	},
	"orders_sell_gainp": {
		Type:       "percent",
		LabelLong:  "Gain en %",
		LabelShort: "en %",
		ColWith:    80,
		Protected:  true,
		// ComputeSQL: "select ( ({orders_sell} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_sell} * {orders_quantity} * {__cost}) / ({orders_buy} * {orders_quantity}) )*100",
	},
	"_image_day": {
		Type:       "image",
		LabelLong:  "Graph du jour",
		LabelShort: "Graph J",
		Params: types.Params{
			Path:     "/crud/data/picsou/png/day/{orders_ptf_id}.png",
			IconName: "emblem-photos",
		},
	},
	"_image_histo": {
		Type:       "image",
		LabelLong:  "Historique sur 1 mois",
		LabelShort: "Histo",
		Params: types.Params{
			Path:     "/crud/data/picsou/png/quotes/{orders_ptf_id}.png",
			IconName: "emblem-photos",
		},
	},
	"_image_analyse": {
		Type:       "image",
		LabelLong:  "Analyse sur 7 mois",
		LabelShort: "Analyse",
		Params: types.Params{
			Path:     "/crud/data/picsou/png/ana/{orders_ptf_id}.gif",
			IconName: "emblem-photos",
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
}
