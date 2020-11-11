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
	Actions       Actions // bouton d'actions - utilise Params
	Args          Args
	Class         string   // Class du texte dans la cellule https://fomantic-ui.com/collections/table.html
	ClassSQL      string   // SQL pour alimenter Class error warning info green blue
	ColAlign      string   //
	ColWith       int      // TODO largeur de la colonne
	Default       string   // Valeur par défaut (macro possible)
	DefaultSQL    string   // Ordre SQL qui retournera la colonne pour alimenter Default
	Error         string   // contiendra "error" si le champ est en erreur de saisie
	Format        string   // "%3.2f %%" "%3.2f €" date datetime time
	ComputeSQL    string   // formule de calcul de Value en SQL dans VIEW EDIR ADD (pas dans LIST)
	Grid          string   // Class pour donner la largeur du champ dans le formulaire "four wide field" 16 colonnes
	Group         string   // Groupe autorisé à accéder à cette rubrique - Si owner c'est l'enregistreement qui sera protégé
	Height        int      // TODO hauteur de la colonne
	Help          string   // TODO aide sur la rubrique
	HelpSQL       string   // TODO aide sql sur la rubrique
	Hide          bool     // élémnt caché dans la vue ou formulaire
	HideSQL       string   // TODO cachée si condition
	HideOnMobile  bool     // La colonne dans une vue sera cachée sur Mobile
	Items         []Item   // slice d'item
	ItemsSQL      string   // Ordre SQL qui retournera la colonne pour alimenter Items
	Jointure      Jointure // élément issu d'une jointure SQL (lecture seule)
	LabelLong     string   // Label dans un formulaire
	LabelShort    string   // Label dans une vue
	Max           int      // TODO valeur max
	MaxLength     int      // TODO longeur max
	Min           int      // TODO valeur min
	MinLength     int      // TODO longueur min
	Order         int      // Ordre de l'élément dans une vue ou formulaire
	Params        Params   // paramètres optionnels
	Pattern       string   // Pattern de l'input HTML
	PlaceHolder   string   // Label dans le champ en saisie si vide
	PostAction    Actions  // actions sql ou plugin à exécuter après la mise à jour
	Protected     bool     // Est en misa à jour mais protégé en saisie
	ReadOnly      bool     // Lecteur seule
	Refresh       bool     // TODO avec un bouton refresh pour actualiser le formulaire en mise à jour
	Required      bool     // obligatoire
	SortDirection string   // "", ascending, ou descending pour demander un tri à la requête sql
	SQLout        string   // Valeur à enregistrer dans la base de données (zone calculée par le beedule)
	Type          string   // Type : action amount button checkbox combobox counter date datetime email float image jointure list markdown month number pdf percent plugin section tag tel text time radio url week
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
	Variables  map[string]string
}

// View Vue d'une table
type View struct {
	Actions      Actions  // Action sur la vue (ordres sql)
	ClassSQL     string   // couleur theme de la ligne
	Deletable    bool     // Suppression fiche autorisée
	FormAdd      string   // Formulaire d'ajout
	FormEdit     string   // Formulaire d'édition
	FormView     string   // Masque de visualisation d'un enregistreement
	FooterSQL    string   // requête sur la table courante
	Hide         bool     // Vue cachée dans le menu
	HideOnMobile bool     // Vue cachée dur mobile
	IconName     string   // nom de l'icone
	Limit        int      // pour limiter le nbre de ligne dans la vue
	Group        string   // groupe qui peut accéder à la vue
	OrderBy      string   // Tri des données SQL
	Where        string   // Condition SQL
	Type         string   // type de vue : card(default),image,table
	Title        string   // Titre de la vue
	Elements     Elements // Eléments à récupérer de la base de données
	Mask         MaskList // Masque html d'une ligne dans la vue
	PreUpdateSQL []string // requêtes SQL avant l'affichage
	Search       string   // Chaîne de recherche dans toutes les colonnes de la vue
}

// MaskList Content d'une card https://fomantic-ui.com/views/card.html
type MaskList struct {
	Header      []string
	Meta        []string
	Description []string
	Extra       []string
}

// Form formulaire
type Form struct {
	Actions    Actions  // Action appel d'un formulaire ou exécution d'une requête SQL
	Title      string   // Titre du formulaire
	Group      string   // groupe qui peut accéder au formulaire
	HideSubmit bool     // pour caher le bouton valider
	IconName   string   // nom de l'icone
	Elements   Elements // Eléments à récupérer de la base de données
	CheckSQL   []string // retourne le libellé des erreurs lors du contrôle des rubriques
	PostSQL    []string // Ordre exécutée après la validation si contrôle OK
}

// DBAlias alias des connections aux bases de données
type DBAlias struct {
	DriverName string // mysql
	DataSource string // <user>:<mdp>@tcp(<host>:3306)/<basename>?charset=utf8
}

// Jointure entre tables
type Jointure struct {
	Join   string // la commande du genre : left outer join on field = field
	Column string // colonne retournée par la jointure
}

// Params paramètres d'un élément
type Params struct {
	Action          string
	Form            string
	IconName        string
	Header          []string
	Description     []string
	Meta            []string
	Extra           []string
	URL             string
	Path            string
	Src             string
	SQL             []string
	WithConfirm     bool
	WithInput       bool
	WithInputFile   bool
	WithImageEditor bool
}

// Actions as
type Actions []Action

// Action dans le menu d'une vue ou formulaire
type Action struct {
	Group       string   // Groupe autorisée à lancer l'action
	Label       string   // label de l'action
	Checkbox    Setters  // checkbox pour mettre à jour la donnée
	URL         string   // URL d'appel du formulaire
	SQL         []string // les ordres SQL seront exécutées avant l'appel du formulaire
	WithConfirm bool     // demande de  confirmation
	Hide        bool     // Action non visible
	HideSQL     string   // requête pour cachée l'action
	Plugin      string   // Fonction Système à appeler nomFonction(p1, p2, ...)
}

// Setters as
type Setters struct {
	GetSQL  string // requête pour lire la données
	SetSQL  string // requête pour mettre à jour la données
	AliasDB string // Connecteur base de données
}

// Args paramètres à transmettre lors de l'appel
type Args map[string]string

// Item pour définir un combobox
type Item struct {
	Key   string // valeur dans la base de données
	Label string // label à afficher
}

// Config Paramètres de config pour alimenter le "A propos"
type Config struct {
	Appname string
	Appnote string
	Date    string
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

// HashPassword hashage de Value
func (element *Element) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		beego.Error(err)
	}
	return string(bytes)
}

// Parameters as Table Parameters dans base default sqlite
type Parameters struct {
	ID    string `orm:"pk;column(id)"`
	Label string
	Value string
}

// TableName as
func (u *Parameters) TableName() string {
	return "parameters"
}
