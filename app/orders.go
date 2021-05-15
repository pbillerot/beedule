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
	"vgraph": {
		Title:     "Achats en graphique",
		Type:      "image",
		IconName:  "shopping cart",
		FormEdit:  "feditbuy",
		FormView:  "fviewbuy",
		Deletable: false,
		Hide:      false,
		Where:     "orders_order = 'buy'",
		ClassSQL:  "select case when {orders_cost_price} + {orders_cost_price} * {__optimum} < {orders_quote} then 'positive' when {orders_cost_price} < {orders_quote} then 'blue' else 'negative' end || ' crud-card-view2'",
		PreUpdateSQL: []string{
			"update orders set orders_quote = (select close from quotes where id = orders_ptf_id and date = (select max(date) from quotes where id = orders_ptf_id))",
			"update orders set orders_gain = orders_quote * orders_quantity - orders_buy * orders_quantity - orders_buy * orders_quantity * {__cost} - orders_quote * orders_quantity * {__cost}",
			"update orders set orders_gainp = (orders_gain / (orders_buy * orders_quantity)) * 100",
			"update orders set orders_debit = orders_buy * orders_quantity + orders_buy * orders_quantity * {__cost}",
		},
		Elements: map[string]types.Element{
			"orders_id":         {Order: 1, Hide: true},
			"orders_ptf_id":     {Order: 10, Hide: true},
			"orders_order":      {Order: 20, Hide: true},
			"orders_cost_price": {Order: 30, Hide: true},
			"orders_optimum":    {Order: 40, Hide: true},
			"orders_quote":      {Order: 40, Hide: true},
			"_chart_quotes":     {Order: 100},
		},
	},
	"vachat": {
		Title:     "Achat",
		IconName:  "shopping cart",
		FormAdd:   "feditbuy",
		FormEdit:  "feditbuy",
		FormView:  "fviewbuy",
		Deletable: true,
		Hide:      false,
		Where:     "orders_order = 'buy'",
		ClassSQL:  "select case when {orders_cost_price} + {orders_cost_price} * {__optimum} < {orders_quote} then 'positive' when {orders_cost_price} < {orders_quote} then 'blue' else 'negative' end",
		PreUpdateSQL: []string{
			"update orders set orders_quote = (select close from quotes where id = orders_ptf_id and date = (select max(date) from quotes where id = orders_ptf_id))",
			"update orders set orders_gain = orders_quote * orders_quantity - orders_buy * orders_quantity - orders_buy * orders_quantity * {__cost} - orders_quote * orders_quantity * {__cost}",
			"update orders set orders_gainp = (orders_gain / (orders_buy * orders_quantity)) * 100",
			"update orders set orders_debit = orders_buy * orders_quantity + orders_buy * orders_quantity * {__cost}",
		},
		Mask: types.MaskList{
			Header: []string{
				"orders_ptf_id",
			},
			Meta: []string{
				"orders_time",
			},
			Description: []string{
				"orders_rem",
			},
			Extra: []string{
				"ptf_quote",
				"ptf_gain",
				"orders_gain",
				"orders_gainp",
			},
		},
		Elements: map[string]types.Element{
			"orders_id":         {Order: 1, HideOnMobile: true},
			"orders_ptf_id":     {Order: 10},
			"orders_time":       {Order: 30, HideOnMobile: true},
			"orders_buy":        {Order: 40, HideOnMobile: true},
			"orders_cost_price": {Order: 80, HideOnMobile: true},
			"orders_optimum":    {Order: 100, HideOnMobile: true},
			"orders_quote":      {Order: 110, HideOnMobile: true},
			"orders_gain":       {Order: 120},
			"orders_gainp":      {Order: 130},
			"ptf_quote":         {Order: 140},
			"ptf_gain":          {Order: 150},
			"orders_rem":        {Order: 160, HideOnMobile: true},
		},
	},
	"vvente": {
		Title:     "Vente",
		IconName:  "trophy",
		FormEdit:  "feditbuy",
		FormView:  "fviewsell",
		Deletable: true,
		Where:     "orders_order = 'sell'",
		Mask: types.MaskList{
			Header: []string{
				"orders_ptf_id",
			},
			Meta: []string{
				"orders_time",
				"orders_sell_time",
			},
			Description: []string{
				"orders_rem",
			},
			Extra: []string{
				"orders_sell_gain",
				"orders_sell_gainp",
			},
		},
		Elements: map[string]types.Element{
			"orders_id":         {Order: 1, HideOnMobile: true},
			"orders_ptf_id":     {Order: 10},
			"orders_time":       {Order: 30, HideOnMobile: true},
			"orders_sell_time":  {Order: 40, HideOnMobile: true},
			"orders_buy":        {Order: 60, HideOnMobile: true},
			"orders_sell":       {Order: 70, HideOnMobile: true},
			"orders_quantity":   {Order: 80, HideOnMobile: true},
			"orders_sell_cost":  {Order: 90, HideOnMobile: true},
			"orders_credit":     {Order: 100, HideOnMobile: true},
			"orders_sell_gain":  {Order: 110},
			"orders_sell_gainp": {Order: 120, HideOnMobile: true},
		},
	},
}
var ordersForms = types.Forms{
	"fviewbuy": {
		Title: "Ordre d'achat",
		Elements: map[string]types.Element{
			// Achat
			"orders_id":       {Order: 1},
			"orders_ptf_id":   {Order: 10},
			"orders_order":    {Order: 20},
			"orders_time":     {Order: 30},
			"orders_buy":      {Order: 50},
			"orders_cost":     {Order: 60},
			"orders_quantity": {Order: 40},
			"orders_debit":    {Order: 80},
			// Evolution
			"_section_achat_situation": {
				Order:     100,
				Type:      "section",
				LabelLong: "Évolution",
				Params: types.Params{
					Form:     "frem",
					IconName: "balance scale left",
				},
			},
			"orders_cost_price": {Order: 110},
			"orders_optimum":    {Order: 120},
			"orders_quote":      {Order: 130},
			"_refresh_buy":      {Order: 135},
			"orders_gain":       {Order: 140},
			"orders_gainp":      {Order: 150},
			"orders_rem":        {Order: 170},
			"_action_sell":      {Order: 270},
			// Images
			"_chart_quotes":  {Order: 310},
			"_image_analyse": {Order: 320},
		},
	},
	"fviewsell": {
		Title: "Ordre de vente",
		Elements: map[string]types.Element{
			// Achat
			"_section_achat": {
				Order:     10,
				Type:      "section",
				LabelLong: "Achat",
				Params: types.Params{
					Form:     "feditbuy",
					IconName: "money check",
				},
			},
			"orders_id":         {Order: 15},
			"orders_ptf_id":     {Order: 10},
			"orders_order":      {Order: 20},
			"orders_time":       {Order: 30},
			"orders_buy":        {Order: 50},
			"orders_cost":       {Order: 60},
			"orders_cost_price": {Order: 62},
			"orders_optimum":    {Order: 64},
			"orders_quantity":   {Order: 40},
			"orders_debit":      {Order: 80},
			// Vente
			"_section_vente": {
				Order:      200,
				Type:       "section",
				LabelLong:  "Vente",
				LabelShort: "Vente",
				Params: types.Params{
					Form:     "feditsell",
					IconName: "money check",
				},
			},
			"orders_sell_time":  {Order: 210},
			"orders_sell":       {Order: 220},
			"orders_sell_cost":  {Order: 230},
			"orders_credit":     {Order: 240},
			"orders_sell_gain":  {Order: 250},
			"orders_sell_gainp": {Order: 260},
			// Images
			"_image_analyse": {Order: 320},
			"_chart_quotes":  {Order: 330},
		},
	},
	"feditbuy": {
		Title: "Ordre d'achat",
		Group: "trader",
		Elements: map[string]types.Element{
			"orders_id":       {Order: 1},
			"orders_ptf_id":   {Order: 10, Required: true},
			"orders_order":    {Order: 20, Default: "buy"},
			"orders_time":     {Order: 30, Required: true},
			"orders_quote":    {Order: 40, ReadOnly: true},
			"orders_buy":      {Order: 50, Required: true},
			"orders_quantity": {Order: 60, Required: true},
			"orders_debit":    {Order: 70},
		},
		PostSQL: []string{
			"update orders set orders_quote = (select close from quotes where id = orders_ptf_id and date = (select max(date) from quotes where id = orders_ptf_id))",
			"update orders set orders_gain = orders_quote * orders_quantity - orders_buy * orders_quantity - orders_buy * orders_quantity * {__cost} - orders_quote * orders_quantity * {__cost}",
			"update orders set orders_gainp = (orders_gain / (orders_buy * orders_quantity)) * 100",
			"update orders set orders_debit = orders_buy * orders_quantity + orders_buy * orders_quantity * {__cost}",
			"update orders set orders_cost = orders_buy * orders_quantity * {__cost}",
			"update orders set orders_cost_price = (orders_buy * orders_quantity + orders_buy * orders_quantity * ({__cost}*2))/orders_quantity",
		},
	},
	"feditsell": {
		Title: "Ordre de vente",
		Group: "trader",
		Elements: map[string]types.Element{
			"orders_id":         {Order: 1},
			"orders_ptf_id":     {Order: 10},
			"orders_order":      {Order: 20},
			"orders_sell_time":  {Order: 30},
			"orders_quote":      {Order: 40},
			"orders_sell":       {Order: 50},
			"orders_quantity":   {Order: 60},
			"orders_credit":     {Order: 70},
			"orders_sell_gain":  {Order: 80},
			"orders_sell_gainp": {Order: 90},
		},
		PostSQL: []string{
			"update orders set orders_sell_cost = orders_buy * orders_quantity * {__cost} + orders_sell * orders_quantity * {__cost}",
			"update orders set orders_sell_gain = orders_sell * orders_quantity - orders_buy * orders_quantity - orders_buy * orders_quantity * {__cost} - orders_sell * orders_quantity * {__cost}",
			"update orders set orders_sell_gainp = (orders_sell_gain / (orders_buy * orders_quantity)) * 100",
			"update orders set orders_credit = orders_sell * orders_quantity + orders_sell * orders_quantity * {__cost}",
		},
	},
	"frem": {
		Title: "Remarques",
		Group: "trader",
		Elements: map[string]types.Element{
			"orders_id":     {Order: 1},
			"orders_ptf_id": {Order: 10},
			"orders_rem":    {Order: 20},
		},
	},
}

var ordersElements = map[string]types.Element{
	"ptf_quote": {
		Type:       "amount",
		LabelLong:  "Quote",
		LabelShort: "Quote",
		Jointure: types.Jointure{
			Join:   "left outer join ptf on ptf_id = orders_ptf_id",
			Column: "ptf.ptf_quote",
		},
	},
	"ptf_gain": {
		Type:       "percent",
		LabelLong:  "Gain du jour",
		LabelShort: "Gain du jour",
		Jointure: types.Jointure{
			// Join:   "left outer join ptf on ptf_id = orders_ptf_id",
			Column: "ptf.ptf_gain",
		},
		ClassSQL: "select case when {ptf_gain} > 0 then 'green' when {ptf_gain} < 0 then 'red' end",
	},
	"_refresh_buy": {
		Type:      "action",
		LabelLong: "Mettre à jour avec le cours du jour",
		Actions: []types.Action{
			{
				SQL: []string{
					"update orders set orders_quote = (select close from quotes where id = orders_ptf_id and date = (select max(date) from quotes where id = orders_ptf_id))",
					"update orders set orders_gain = orders_quote * orders_quantity - orders_buy * orders_quantity - orders_buy * orders_quantity * {__cost} - orders_quote * orders_quantity * {__cost}",
					"update orders set orders_gainp = (orders_gain / (orders_buy * orders_quantity)) * 100",
				},
			},
		},
	},
	"_action_sell": {
		Type:      "button",
		Group:     "trader",
		LabelLong: "Vendre cette valeur...",
		Params: types.Params{
			URL: "/bee/edit/picsou/orders/vachat/feditsell/{orders_id}?orders_order=sell&orders_sell={orders_quote}",
		},
	},
	"orders_id": {
		Type:       "counter",
		LabelLong:  "N°",
		LabelShort: "N°",
		ColAlign:   "center",
	},
	"orders_ptf_id": {
		Type:       "combobox",
		LabelLong:  "Valeur",
		LabelShort: "Valeur",
		ItemsSQL:   "SELECT ptf_id as 'key', ptf_name as 'label' From ptf order by ptf_name",
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
		Type:       "combobox",
		LabelLong:  "Order",
		LabelShort: "Order",
		ColAlign:   "center",
		Required:   true,
		Items: []types.Item{
			{Key: "buy", Label: "Achat"},
			{Key: "sell", Label: "Vente"},
		},
	},
	"orders_rem": {
		Type:       "textarea",
		LabelLong:  "Remarque",
		LabelShort: "Rem.",
	},
	"orders_time": {
		Type:       "datetime",
		LabelLong:  "Jour Heure d'achat",
		LabelShort: "JH d'achat",
		DefaultSQL: "select datetime('now', 'localtime')",
	},
	"orders_sell_time": {
		Type:       "datetime",
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
		Required:   true,
		Refresh:    true,
		DefaultSQL: "select '{__amount_min}' / '{orders_buy}'",
	},
	"orders_buy": {
		Type:       "amount",
		LabelLong:  "Cours d'achat",
		LabelShort: "Achat à",
		Required:   true,
		Refresh:    true,
	},
	"orders_sell": {
		Type:       "amount",
		LabelLong:  "Cours de vente",
		LabelShort: "Vente à",
		DefaultSQL: "select '{orders_quote}'",
		Required:   true,
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
		ComputeSQL: "select '{orders_buy}' * '{orders_quantity}' + '{orders_buy}' * '{orders_quantity}' * '{__cost}'",
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
		ClassSQL:   "select case when {orders_gain} > 0 then 'green' when {orders_gain} < 0 then 'red' end",
		// ComputeSQL: "select {orders_quote} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_quote} * {orders_quantity} * {__cost}",
	},
	"orders_gainp": {
		Type:       "percent",
		LabelLong:  "Gain en %",
		LabelShort: "en %",
		ColWith:    80,
		Protected:  true,
		ClassSQL:   "select case when {orders_gainp} > 0 then 'green' when {orders_gainp} < 0 then 'red' end",
		// ComputeSQL: "select ( ({orders_quote} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_quote} * {orders_quantity} * {__cost}) / ({orders_buy} * {orders_quantity}) )*100",
	},
	"orders_sell_cost": {
		Type:       "amount",
		LabelLong:  "Frais",
		LabelShort: "Frais",
		ColWith:    60,
		Protected:  true,
		// ComputeSQL: "select {orders_buy} * {orders_quantity} * {__cost} + {orders_sell} * {orders_quantity} * {__cost}",
	},
	"orders_sell_gain": {
		Type:       "amount",
		LabelLong:  "Gain",
		LabelShort: "Gain",
		ColWith:    80,
		Protected:  true,
		ClassSQL:   "select case when {orders_sell_gain} > 0 then 'green' when {orders_sell_gain} < 0 then 'red' end",
		// ComputeSQL: "select {orders_sell} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_sell} * {orders_quantity} * {__cost}",
	},
	"orders_sell_gainp": {
		Type:       "percent",
		LabelLong:  "Gain en %",
		LabelShort: "en %",
		ColWith:    80,
		Protected:  true,
		ClassSQL:   "select case when {orders_sell_gainp} > 0 then 'green' when {orders_sell_gainp} < 0 then 'red' end",
		// ComputeSQL: "select ( ({orders_sell} * {orders_quantity} - {orders_buy} * {orders_quantity} - {orders_buy} * {orders_quantity} * {__cost} - {orders_sell} * {orders_quantity} * {__cost}) / ({orders_buy} * {orders_quantity}) )*100",
	},
	"_image_day": {
		Type:       "image",
		LabelLong:  "Graph du jour",
		LabelShort: "Graph J",
		Params: types.Params{
			Path:     "/bee/data/picsou/png/day/{orders_ptf_id}.png",
			URL:      "/bee/data/picsou/png/day/{orders_ptf_id}.png",
			IconName: "emblem-photos",
		},
	},
	"_image_histo": {
		Type:       "image",
		LabelLong:  "Historique sur 1 mois",
		LabelShort: "Histo",
		Params: types.Params{
			Path:     "/bee/data/picsou/png/quotes/{orders_ptf_id}.png",
			URL:      "/bee/data/picsou/png/quotes/{orders_ptf_id}.png",
			IconName: "emblem-photos",
		},
	},
	"_image_analyse": {
		Type:       "image",
		LabelLong:  "Analyse sur 7 mois",
		LabelShort: "Analyse",
		Params: types.Params{
			Path:     "/bee/data/picsou/png/ana/{orders_ptf_id}.gif",
			URL:      "/bee/data/picsou/png/ana/{orders_ptf_id}.gif",
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
	"_chart_quotes": {
		Type:       "image",
		LabelLong:  "Cotation sur 1 mois",
		LabelShort: "Cotation",
		Dataset: map[string]string{
			"ClassJquery": "select 'bee-chart-quotes'",
			"Title":       "select 'Cours de {orders_ptf_id}'",
			"Quotes":      "select open as matin, close as soir from quotes where id = '{orders_ptf_id}' order by date",
			"Quotep":      "select (open-close1)*100/close1 as matin, (close-close1)*100/close1 as soir from quotes where id = '{orders_ptf_id}' order by date",
			"Labels":      "select printf('%s-%s',substr(date,9,2),substr(date,6,2)) as 'matin', '-' as 'soir' from quotes where id = '{orders_ptf_id}' order by date",
			"Minp":        "select (low-close1)*100/close1 as matin, (low-close1)*100/close1 as soir from quotes where id = '{orders_ptf_id}' order by date",
			"Maxp":        "select (high-close1)*100/close1 as matin, (high-close1)*100/close1 as soir from quotes where id = '{orders_ptf_id}' order by date",
			"Min":         "select min(low) as min from quotes where id = '{orders_ptf_id}'",
			"Max":         "select max(high) as max from quotes where id = '{orders_ptf_id}'",
			"SeuilV":      "select '{orders_optimum}' where '{orders_order}' = 'buy'",
			"SeuilR":      "select '{orders_cost_price}' where '{orders_order}' = 'buy'",
		},
	},
}
