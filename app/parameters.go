package app

import "github.com/pbillerot/beedule/types"

// Parameters table des paramètres du site
var Parameters = types.Table{
	AliasDB:    "admin",
	Key:        "id",
	ColDisplay: "id",
	IconName:   "tools",
	Elements:   parametersElements,
	Views:      parametersViews,
	Forms:      parametersForms,
}

var parametersElements = map[string]types.Element{
	"id": {
		Type:       "text",
		Order:      1,
		LabelLong:  "Paramètre",
		LabelShort: "Paramètre",
	},
	"value": {
		Type:       "text",
		Order:      10,
		LabelLong:  "Valeur",
		LabelShort: "Valeur",
	},
	"label": {
		Type:       "text",
		Order:      20,
		LabelLong:  "Désignation",
		LabelShort: "Désignation",
	},
	"action": {
		Type:       "action",
		Order:      40,
		LabelLong:  "Démarrer/Arrêter le pendule",
		LabelShort: "Démarrer/Arrêter le pendule",
		Actions: []types.Action{
			{
				SQL: []string{
					"update parameters set value = case when value = '1' then '0' else '1' end where id = 'batch_etat'",
				},
				Plugin: "StartStopPendule()",
			},
		},
	},
}

var parametersViews = types.Views{
	"vall": {
		FormAdd:   "fedit",
		FormEdit:  "fedit",
		FormView:  "fedit",
		Deletable: true,
		Title:     "Paramètres",
		IconName:  "tools",
		OrderBy:   "id",
		Group:     "admin",
		Mask: types.MaskList{
			Header: []string{
				"id",
			},
			Meta: []string{
				"value",
				"action",
			},
			Description: []string{
				"label",
			},
			Extra: []string{},
		},
		Elements: map[string]types.Element{
			"id":    {},
			"value": {},
			"label": {},
		},
	},
}

var parametersForms = types.Forms{
	"fedit": {
		Title: "Paramètre",
		Group: "admin",
		Elements: map[string]types.Element{
			"id":     {},
			"value":  {},
			"label":  {},
			"action": {},
		},
	},
}
