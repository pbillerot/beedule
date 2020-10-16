package app

import (
	"github.com/pbillerot/beedule/types"
)

// Jobs table des groupes d'utilisateurs
var Jobs = types.Table{
	AliasDB:    "pendule",
	Key:        "job_id",
	ColDisplay: "label",
	IconName:   "cog",
	Elements:   jobsElements,
	Views:      jobsViews,
	Forms:      jobsForms,
}

var jobsViews = types.Views{
	"vall": {
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		FormView:  "fview",
		Deletable: true,
		Title:     "Jobs",
		IconName:  "cog",
		OrderBy:   "chain_label, sequence",
		Group:     "admin",
		ClassSQL:  "select case when '{etat}' = 'RUN' then 'blue' when '{etat}' = 'OK' then 'green' when '{etat}' = 'KO' then 'red' else '' end",
		Mask: types.MaskList{
			Header: []string{
				"chain_id",
				"job_id",
				"label",
			},
			Meta: []string{
				"active",
				"sequence",
				"sierreur",
				"type",
			},
			Description: []string{
				"result",
			},
			Extra: []string{
				"etat",
				"heuredebut",
				"heurefin",
				"dureemn",
			},
		},
		Elements: types.Elements{
			"job_id":      {},
			"label":       {},
			"chain_label": {},
			"chain_id":    {},
			"sequence":    {},
			"active":      {},
			"sierreur":    {},
			"type":        {},
			"commandes":   {},
			"options":     {},
			"doc":         {},
			"etat":        {},
			"result":      {},
			"heuredebut":  {},
			"heurefin":    {},
			"dureemn":     {},
		},
	},
}

var jobsForms = types.Forms{
	"fview": {
		Title: "Pendule Job",
		Group: "admin",
		Elements: types.Elements{
			"job_id": {Order: 01, Hide: true},
			"_SECTION_JOB": {
				Order:     300,
				Type:      "section",
				LabelLong: "JOB",
				Params: types.Params{
					Form:     "fjob",
					IconName: "cog",
				},
			},
			"active":    {Order: 310},
			"label":     {Order: 320},
			"sequence":  {Order: 322},
			"sierreur":  {Order: 324},
			"type":      {Order: 330},
			"commandes": {Order: 340},
			"_SECTION_RESULTAT": {
				Order:     500,
				Type:      "section",
				LabelLong: "RESULTATS",
				Params: types.Params{
					IconName: "file alternate outline",
				},
			},
			"etat":       {Order: 510},
			"result":     {Order: 515},
			"heuredebut": {Order: 520},
			"heurefin":   {Order: 530},
			"dureemn":    {Order: 540},
			"_action_job": {
				Order:     550,
				Type:      "action",
				LabelLong: "Démarrer ce job",
				Action: types.Action{
					Plugin: "StartJob({job_id})",
				},
			},
			"_SECTION_CHAINE": {
				Order:     700,
				Type:      "section",
				LabelLong: "CHAÎNE",
				Params: types.Params{
					URL:      "/crud/edit/pendule/chains/vall/fedit/{chain_id}",
					IconName: "calendar alternate outline",
				},
			},
			"chain_id":         {Order: 710},
			"chain_active":     {Order: 720},
			"chain_planif":     {Order: 730},
			"chain_etat":       {Order: 740},
			"chain_heuredebut": {Order: 750},
			"chain_heurefin":   {Order: 760},
			"chain_dureemn":    {Order: 770},
		},
	},
	"fadd": {
		Title: "Pendule ajout d'un job",
		Group: "admin",
		Elements: types.Elements{
			"job_id":    {Order: 01, Hide: true},
			"active":    {Order: 10},
			"chain_id":  {Order: 20},
			"label":     {Order: 30},
			"sequence":  {Order: 40},
			"sierreur":  {Order: 50},
			"type":      {Order: 60},
			"commandes": {Order: 70},
		},
	},
	"fjob": {
		Title:    "Job",
		Group:    "admin",
		IconName: "cog",
		Elements: types.Elements{
			"job_id":    {Order: 01, Hide: true},
			"active":    {Order: 10},
			"chain_id":  {Order: 20},
			"label":     {Order: 30},
			"sequence":  {Order: 40},
			"sierreur":  {Order: 50},
			"type":      {Order: 60},
			"commandes": {Order: 70},
		},
	},
}

var jobsElements = types.Elements{
	"job_id": { // n° du job
		Type:       "counter",
		LabelLong:  "n°",
		LabelShort: "n°",
		Class:      "blue",
	},
	"label": { // Nom du job dans la chaîne
		Type:       "text",
		LabelLong:  "Nom",
		LabelShort: "Nom",
		Class:      "blue",
	},
	"active": { // Job actif
		Type:       "checkbox",
		LabelLong:  "Actif",
		LabelShort: "Actif",
	},
	"etat": { // État du job
		Type:       "text", // INI, OK, KO, RUN
		LabelLong:  "État",
		LabelShort: "État",
	},
	"sequence": { // n° de séquence dans la chaîne (on commence à 0)
		Type:       "number",
		LabelLong:  "Séquence",
		LabelShort: "Séquence",
	},
	"sierreur": { // n° de séquence dans la chaîne (on commence à 0)
		Type:       "combobox",
		LabelLong:  "Si erreur",
		LabelShort: "Si erreur",
		Items: []types.Item{
			{Key: "1", Label: "Job Suivant"},
			{Key: "2", Label: "Étape suivante"},
			{Key: "3", Label: "Arrêt de la chaîne"},
		},
		ClassSQL: "select case when '{sierreur}' = '2' then 'green' when '{sierreur}' = '3' then 'red' else '' end",
	},
	"type": { // Type de Job
		Type:       "combobox",
		LabelLong:  "Type",
		LabelShort: "Type",
		Items: []types.Item{
			{Key: "conditionsql", Label: "Condition SQL"},
			{Key: "etape", Label: "Étape"},
			{Key: "plugin", Label: "Plugin"},
			{Key: "sql", Label: "SQL"},
		},
	},
	"commandes": { // Commandes à exécuter
		Type:       "textarea",
		LabelLong:  "Commandes",
		LabelShort: "Commandes",
	},
	"options": { // Options du job
		Type:       "textarea",
		LabelLong:  "Options",
		LabelShort: "Options",
	},
	"result": { // Résultat du job
		Type:       "text",
		LabelLong:  "Résultat",
		LabelShort: "Résultat",
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
	//
	// JOINTURES CHAINE
	//
	"chain_id": { // lien vers la chaîne
		Type:       "combobox",
		LabelLong:  "Chaîne",
		LabelShort: "Chaîne",
		Class:      "violet",
		ItemsSQL:   "select label as 'label', chain_id as 'key' from chains order by label",
	},
	// le Join ne sera défini que sur une seule colonne
	"chain_label": { // Nom de la chaîne
		Type:       "text",
		LabelLong:  "Chaîne",
		LabelShort: "Chaîne",
		Class:      "violet",
		Jointure: types.Jointure{
			Join:   "left outer join chains on chains.chain_id = jobs.chain_id",
			Column: "chains.label",
		},
	},
	"chain_planif": { // planification
		Type:       "text",
		LabelLong:  "Planification",
		LabelShort: "Planification",
		Class:      "orange",
		Jointure: types.Jointure{
			Join:   "left outer join chains on chains.chain_id = jobs.chain_id",
			Column: "chains.planif",
		},
	},
	"chain_active": { // Chaîne active
		Type:       "checkbox",
		LabelLong:  "Active",
		LabelShort: "Active",
		Jointure: types.Jointure{
			Column: "chains.active",
		},
	},
	"chain_etat": { // Etat de la chaîne
		Type:       "text", // INI, OK, KO, RUN
		LabelLong:  "État",
		LabelShort: "État",
		Jointure: types.Jointure{
			Column: "chains.etat",
		},
	},
	"chain_heuredebut": { // Heure début de la chaîne
		Type:       "datetime",
		LabelLong:  "Démarrée le",
		LabelShort: "Démarrée le",
		Jointure: types.Jointure{
			Column: "chains.heuredebut",
		},
	},
	"chain_heurefin": { // Heure de fin de la chaîne
		Type:       "datetime",
		LabelLong:  "Terminée le",
		LabelShort: "Terminée le",
		Jointure: types.Jointure{
			Column: "chains.heurefin",
		},
	},
	"chain_dureemn": { // Durée d'exécution de la chaîne
		Type:       "minute",
		LabelLong:  "Durée",
		LabelShort: "Durée",
		Jointure: types.Jointure{
			Column: "chains.dureemn",
		},
	},
}
