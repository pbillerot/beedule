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
	Group    string
	Path     string    // Path ou URL de l'application externe
	DataPath string    // répertoire des données valorisé dans custom.conf [app] datapath lu par {$datapath}
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
	Action       Action // bouton d'action
	Args         Args
	Class        string // Class du texte dans la cellule https://fomantic-ui.com/collections/table.html
	ClassSQL     string // SQL pour alimenter Class error warning info green blue
	ColAlign     string //
	ColWith      int    // TODO
	Default      string // Valeur par défaut (macro possible)
	DefaultSQL   string // Ordre SQL qui retournera la colonne pour alimenter Default
	Error        string // contiendra "error" si le champ est en erreur de saisie
	Format       string //
	ComputeSQL   string // formule de calcul de Value en SQL
	Height       int    // TODO
	Help         string // TODO
	HelpSQL      string // TODO
	Hide         bool
	HideSQL      string   // TODO cachée si condition
	HideOnMobile bool     // La colonne dans une vue sera cachée sur Mobile
	Items        []Item   // slice d'item
	ItemsSQL     string   // Ordre SQL qui retournera la colonne pour alimenter Items
	Jointure     Jointure //
	LabelLong    string
	LabelShort   string
	Max          int // TODO
	MaxLength    int // TODO
	Min          int // TODO
	MinLength    int // TODO
	Order        int // Ordre de l'élément dans une vue ou formulaire
	Params       Params
	PlaceHolder  string
	Pattern      string // Pattern de l'input HTML
	Protected    bool
	ReadOnly     bool //
	Refresh      bool // TODO
	Required     bool
	SQLout       string // Valeur à enregistrer dans la base de données
	Type         string // action amount checkbox counter date datetime email float image jointure list month number percent plugin section tag tel text time radio url week
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
	Actions   Actions // Action sur la vue (ordres sql)
	ClassSQL  string  // couleur theme de la ligne
	Deletable bool    // Suppression fiche autorisée
	FormAdd   string
	FormEdit  string
	FormView  string
	FooterSQL string // requête sur la table courante
	Hide      bool   // TODO
	IconName  string
	Info      string
	Limit     int    // TODO
	Group     string // groupe qui peut accéder à la vue
	OrderBy   string
	Where     string
	Type      string // type de vue normal,image
	Title     string // Titre de la vue
	Elements  Elements
	Mask      MaskList // Masque html d'une ligne dans la vue
}

// MaskList as
type MaskList struct {
	Header      []string
	Meta        []string
	Description []string
	Extra       []string
}

// Form formulaire
type Form struct {
	Actions  Actions // Action appel d'un formulaire ou exécution d'une requête SQL
	Title    string
	Group    string // groupe qui peut accéder au formulair
	Elements Elements
	CheckSQL []string // retourne le libellé des erreurs lors du contrôle des rubriques
	PostSQL  []string // Ordre exécutée après la validation si contrôle OK
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

// Params paramètres de l'élément
type Params struct {
	Action   string
	Form     string
	IconName string
	Legend   string
	URL      string
	Path     string
	SQL      []string
}

// Actions as
type Actions []Action

// Action as Formulaire ou Requête SQL
type Action struct {
	Group       string
	Label       string
	URL         string
	SQL         []string // les ordres SQL seront exécutées avant l'appel du formulaire
	WithConfirm bool
	Hide        bool
	HideSQL     string
}

// Args paramètres à transmettre lors de l'appel
type Args map[string]string

// Item entre tables
type Item struct {
	Key   string
	Label string
}

// HashPassword hashage de Value
func (element *Element) HashPassword() string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(element.Value), 14)
	if err != nil {
		beego.Error(err)
	}
	return string(bytes)
}

// Config Paramètres de config dans le contexte
type Config struct {
	Appname string
	Appnote string
	Icone   string
	Site    string
	Email   string
	Author  string
	Version string
	Theme   string
}

// Session Données de session dans le contexte
type Session struct {
	LoggedIn bool
	Username string
	IsAdmin  bool
	Groups   string
}
