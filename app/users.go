package app

import "github.com/pbillerot/beedule/types"

// Users Table des comptes utilisateurs pour gérer les habilitations
var Users = types.Table{
	AliasDB:    "users",
	Key:        "user_name",
	ColDisplay: "user_name",
	IconName:   "user",
	Elements:   usersElements,
	Views:      usersViews,
	Forms:      usersForms,
}

var usersElements = types.Elements{
	"user_name": {
		Type:       "text",
		Order:      1,
		LabelLong:  "Nom ou pseudo",
		LabelShort: "Nom ou pseudo",
		Pattern:    "[a-zA-Z0-9]+",
		Required:   true,
	},
	"user_password": {
		Type:       "password",
		LabelLong:  "Mot de passe",
		LabelShort: "Mot de passe",
		Required:   true,
		MinLength:  3,
	},
	"user_email": {
		Type:       "email",
		LabelLong:  "Email",
		LabelShort: "Email",
		Required:   true,
	},
	"user_is_admin": {
		Type:       "checkbox",
		LabelLong:  "Administrateur",
		LabelShort: "Admin",
		ColAlign:   "center",
	},
	"user_groupes": {
		Type:       "tag",
		LabelLong:  "Groupes",
		LabelShort: "Groupes",
		ItemsSQL:   "select group_id as key, group_id as value from groups order by group_id",
	},
	"_SECTION_MDP": {
		Type:      "section",
		LabelLong: "Sécurité",
		Params: types.Params{
			Form:     "fmdp",
			IconName: "lock",
		},
	},
}

var usersViews = types.Views{
	"vall": {
		FormView:  "fview",
		FormAdd:   "fadd",
		FormEdit:  "fedit",
		Group:     "admin",
		Deletable: true,
		Info:      "Gestion des Comptes",
		Title:     "Comptes",
		IconName:  "user",
		OrderBy:   "user_name",
		Mask: types.MaskList{
			Header: []string{
				"user_name",
			},
			Meta: []string{
				"user_email",
			},
			Description: []string{},
			Extra: []string{
				"user_is_admin",
				"user_groupes",
			},
		},
		Elements: types.Elements{
			"user_name":     {Order: 10},
			"user_email":    {Order: 20, HideOnMobile: true},
			"user_is_admin": {Order: 30, HideOnMobile: true},
			"user_groupes":  {Order: 40},
		},
	},
}

var usersForms = types.Forms{
	"fadd": {
		Title: "Fiche Compte",
		Group: "admin",
		Elements: types.Elements{
			"user_name":  {Order: 10},
			"user_email": {Order: 20},
		},
	},
	"fview": {
		Title: "Fiche Compte",
		Group: "admin",
		Elements: types.Elements{
			"user_name":     {Order: 10},
			"user_email":    {Order: 20},
			"_SECTION_MDP":  {Order: 30},
			"user_password": {Order: 40},
			"user_is_admin": {Order: 50},
			"user_groupes":  {Order: 60},
		},
	},
	"fmdp": {
		Title: "Sécurité",
		Group: "admin",
		Elements: types.Elements{
			"user_name":     {Order: 10, ReadOnly: true},
			"user_password": {Order: 40},
			"user_is_admin": {Order: 50},
			"user_groupes":  {Order: 60},
		},
	},
	"fedit": {
		Title: "Fiche Compte",
		Group: "admin",
		Elements: types.Elements{
			"user_name":  {Order: 10},
			"user_email": {Order: 20},
		},
	},
}
