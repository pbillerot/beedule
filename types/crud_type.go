package types

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

// Portail du serveur Beedule
type Portail struct {
	Title        string
	Info         string
	IconFile     string
	Applications map[string]Application
	Tables       Tables
}

// Application présentée sur le portail
type Application struct {
	Name     string
	Title    string
	Image    string
	IconFile string
	IconName string
	Path     string    // Path ou URL de l'application externe
	AppViews []AppView // Vues des tables liées à l'application
}

// AppView Vue liée à l'application
type AppView struct {
	Tableid string
	Viewid  string
}

// Tables map de Table
type Tables map[string]Table

// Views map de View
type Views map[string]View

// Forms map de Form
type Forms map[string]Form

// Elements map de Element
type Elements map[string]Element

// Element is ... Rubrique de l'application
type Element struct {
	Args         Arg
	ColAlign     string // TODO
	Class        string // TODO Class du texte dans la cellule https://fomantic-ui.com/collections/table.html
	ClassSQL     string // TODO SQL pour alimenter Class
	ColWith      int    // TODO
	Default      string // Valeur par défaut (macro possible)
	DefaultSQL   string // Ordre SQL qui retournera la colonne pour alimenter Default
	Error        string // contiendra "error" si le champ est en erreur de saisie
	Format       string // TODO
	ComputeSQL   string // TODO formule de calcul de Value en SQL
	Height       int    // TODO
	Help         string // TODO
	HelpSQL      string // TODO
	Hide         bool
	HideOnMobile bool     // La colonne dans une vue sera cachée sur Mobile
	Items        []Item   // slice d'item
	ItemsSQL     string   // Ordre SQL qui retournera la colonne pour alimenter Items
	Jointure     Jointure // TODO
	LabelLong    string
	LabelShort   string
	Max          int // TODO
	MaxLength    int // TODO
	Min          int // TODO
	MinLength    int // TODO
	Order        int // Ordre de l'élément dans une vue ou formulaire
	Params       Params
	PlaceHolder  string
	Pattern      string // TODO Pattern de l'input HTML
	Protected    bool
	ReadOnly     bool // TODO
	Refresh      bool // TODO
	Required     bool
	SQLout       string // Valeur à enregistrer dans la base de données
	Type         string // amount checkbox counter date datetime email float jointure list month number percent plugin section tag tel text time radio url week
	Value        string
}

// Table Table de l'application
type Table struct {
	AliasDB    string
	Key        string // clé de la table
	ColDisplay string // la colonne qui identifie l'enregistrement
	IconName   string
	Elements   Elements
	Views      Views
	Forms      Forms
}

// View Vue d'une table
type View struct {
	ActionsSQL     Actions // TODO Action sur la vue (ordres sql)
	ClassSQL       string  // couleur theme de la ligne
	Deletable      bool    // Suppression fiche autorisée
	FormAdd        string
	FormEdit       string
	FormView       string
	FooterSQL      string // TODO à exploiter
	Hide           bool   // TODO
	IconName       string
	Info           string
	Limit          int    // TODO
	Groupe         string // groupe qui peut accéder à la vue  // TODO
	OrderBy        string
	Where          string
	Title          string // Titre de la vue
	ElementsSorted string
	Elements       Elements
}

// Form formulaire
type Form struct {
	Title    string
	Groupe   string // groupe qui peut accéder au formulaire  // TODO
	Elements Elements
}

// DBAlias alias des connections aux bases de données
type DBAlias struct {
	DriverName string
	DataSource string
}

// Jointure entre tables
type Jointure struct {
	Join   string
	Column string
}

// Params entre tables
type Params struct {
	Action   string
	Form     string
	IconName string
	URL      string
	Path     string
	SQL1     string
	SQL2     string
	SQL3     string
	SQL4     string
	SQL5     string
}

// Actions as
type Actions []Action

// Action as
type Action struct {
	Label       string
	SQL         string
	WithConfirm bool
}

// Arg entre tables
type Arg map[string]string

// Item entre tables
type Item struct {
	Key   string
	Value string
}

// HashPassword hashage de Value
func (element *Element) HashPassword() string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(element.Value), 14)
	if err != nil {
		beego.Error(err)
	}
	return string(bytes)
}
