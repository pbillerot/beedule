package dico

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
)

// Ctx as Dictionnaire des applications
var Ctx Portail

// Portail as portail.yaml
type Portail struct {
	Title        string
	Info         string
	IconFile     string `yaml:"icon-file"`
	Applications map[string]Application
	Tables       map[string]*Table // chargé par portail.Load
	Parameters   map[string]string
	Files        []File // liste des fichiers trouvés dans beedic
}

// File as les fichiers du dictionnaire dicodir
type File struct {
	Base string // le filename
	Ext  string // l'extension
	Path string //le chemin compley
}

// Application as
type Application struct {
	Title       string
	Image       string
	IconFile    string `yaml:"icon-file"`
	IconName    string `yaml:"icon-name"` // Icône https://semantic-ui.com/elements/icon.html
	Group       string
	Path        string      // Path ou URL de l'application externe
	Target      string      // _blank pour ouvrir l'application dans un nouvel onglet
	TablesViews []TableView `yaml:"tables-views"` // Vues des tables liées à l'application
}

// TableView as
type TableView struct {
	TableName string `yaml:"table-name"`
	ViewName  string `yaml:"view-name"`
}

// Table as <table>.yaml
type Table struct {
	Setting   Setting
	Elements  map[string]Element
	Views     map[string]View
	Forms     map[string]Form
	Variables map[string]string
}

// Setting as
type Setting struct {
	AliasDB    string `yaml:"alias-db"`
	Key        string // clé de la table
	ColDisplay string `yaml:"col-display"` // la colonne qui identifie l'enregistrement
	IconName   string `yaml:"icon-name"`   // Icône https://semantic-ui.com/elements/icon.html
}

// Element as
type Element struct {
	Actions       []Action // bouton d'actions - utilise Params
	Args          Args
	Class         string            // Class du texte dans la cellule https://fomantic-ui.com/collections/table.html
	ClassSQL      string            `yaml:"class-sql"`  // SQL pour alimenter Class error warning info green blue
	ColAlign      string            `yaml:"col-align"`  //
	ColNoWrap     bool              `yaml:"col-nowrap"` // nowrap de la colonne
	Dataset       map[string]string `yaml:"dataset"`    // Dataset pour un Chartjs
	Default       string            // Valeur par défaut (macro possible)
	DefaultSQL    string            `yaml:"default-sql"` // Ordre SQL qui retournera la colonne pour alimenter Default
	Error         string            // contiendra "error" si le champ est en erreur de saisie
	Format        string            // "%3.2f %%" "%3.2f €" date datetime time
	FormatSQL     string            `yaml:"format-sql"`  // select strftime('%H:%M:%S', {Milliseconds}/1000, 'unixepoch')
	ComputeSQL    string            `yaml:"compute-sql"` // formule de calcul de Value en SQL dans VIEW EDIT ADD (pas dans LIST)
	Grid          string            // Class pour donner la largeur du champ dans le formulaire "four wide field" 16 colonnes
	Group         string            // Groupe autorisé à accéder à cette rubrique - Si owner c'est l'enregistreement qui sera protégé
	Height        int               // TODO hauteur de la colonne
	Help          string            // TODO aide sur la rubrique
	HelpSQL       string            `yaml:"help-sql"` // TODO aide sql sur la rubrique
	Hide          bool              // élémnt caché dans la vue ou formulaire
	HideSQL       string            `yaml:"hide-sql"`       // TODO cachée si condition
	HideOnMobile  bool              `yaml:"hide-on-mobile"` // La colonne dans une vue sera cachée sur Mobile
	Items         []Item            // slice d'item
	ItemsSQL      string            `yaml:"items-sql"` // Ordre SQL qui retournera la colonne pour alimenter Items
	Jointure      Jointure          // élément issu d'une jointure SQL (lecture seule)
	LabelLong     string            `yaml:"label-long"`  // Label dans un formulaire
	LabelShort    string            `yaml:"label-short"` // Label dans une vue
	Max           int               // TODO valeur max
	MaxLength     int               `yaml:"max-length"` // TODO longeur max
	Min           int               // TODO valeur min
	MinLength     int               `yaml:"min-length"` // TODO longueur min
	Order         int               // Ordre de l'élément dans une vue ou formulaire
	Params        Params            // paramètres optionnels
	Pattern       string            // Pattern de l'input HTML
	PlaceHolder   string            `yaml:"place-holder"` // Label dans le champ en saisie si vide
	PostAction    []Action          `yaml:"post-action"`  // actions sql ou plugin à exécuter après la mise à jour
	Protected     bool              // Est en mise à jour mais protégé en saisie
	ReadOnly      bool              `yaml:"read-only"` // Lecteur seule
	Record        orm.Params        // l'enregistrement  valorisant la rubrique
	Refresh       bool              // TODO avec un bouton refresh pour actualiser le formulaire en mise à jour
	Required      bool              // obligatoire
	SortDirection string            `yaml:"sort-direction"` // "", ascending, ou descending pour demander un tri à la requête sql
	SQLout        string            `yaml:"sql-out"`        // Valeur à enregistrer dans la base de données (zone calculée par le beedule)
	Type          string            // Type : action amount button checkbox combobox counter date datetime duration email float image list markdown month number pdf percent plugin section tag tel text time radio url week
}

// HashPassword hashage de Value
func (element *Element) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logs.Error(err)
	}
	return string(bytes)
}

// View Vue d'une table
type View struct {
	Actions        []Action           // Action sur la vue (ordres sql)
	ClassSQL       string             `yaml:"class-sql"` // couleur theme de la ligne
	Deletable      bool               // Suppression fiche autorisée
	FormAdd        string             `yaml:"form-add"`   // Formulaire d'ajout
	FormEdit       string             `yaml:"form-edit"`  // Formulaire d'édition
	FormView       string             `yaml:"form-view"`  // Masque de visualisation d'un enregistrement
	FooterSQL      string             `yaml:"footer-sql"` // requête sur la table courante
	Hide           bool               // Vue cachée dans le menu
	HideOnMobile   bool               `yaml:"hide-on-mobile"` // Vue cachée dur mobile
	IconName       string             `yaml:"icon-name"`      // nom de l'icone
	Limit          int                // pour limiter le nbre de ligne dans la vue
	Group          string             // groupe qui peut accéder à la vue
	OrderBy        string             `yaml:"group-by"` // Tri des données SQL
	Where          string             // Condition SQL
	Type           string             // type de vue : card(default),image,table
	Title          string             // Titre de la vue
	Elements       map[string]Element // Eléments à récupérer de la base de données
	Mask           MaskList           // Masque html d'une ligne dans la vue
	PostSQL        []string           `yaml:"post-sql"`       // Ordre exécutée après la suppression si OK
	PreUpdateSQL   []string           `yaml:"pre-update-sql"` // requêtes SQL avant l'affichage
	Search         string             // Chaîne de recherche dans toutes les colonnes de la vue
	WithLineNumber bool               `yaml:"with-line-number"` // list.table n° de ligne en 1ère colonne
}

// Form formulaire
type Form struct {
	Actions    []Action           // Action appel d'un formulaire ou exécution d'une requête SQL
	Title      string             // Titre du formulaire
	Group      string             // groupe qui peut accéder au formulaire
	HideSubmit bool               `yaml:"hide-submit"` // pour caher le bouton valider
	IconName   string             `yaml:"icon-name"`   // nom de l'icone
	Elements   map[string]Element // Eléments à récupérer de la base de données
	CheckSQL   []string           `yaml:"check-sql"` // retourne le libellé des erreurs lors du contrôle des rubriques
	PostSQL    []string           `yaml:"post-sql"`  // Ordre exécutée après la validation si contrôle OK
}

// Params paramètres d'un élément
type Params struct {
	Action          string
	Form            string   // section: form à ouvrir
	IconName        string   `yaml:"icon-name"` // image: section
	Header          []string // card pour image
	Description     []string // card pour image
	Meta            []string // card pour image
	Extra           []string // card pour image
	URL             string   `yaml:"url"`
	Src             string   // section: src de l'image
	SQL             []string `yaml:"sql"`
	Table           string   // section:
	View            string   // section:
	Where           string   // section: + params.table + params.view
	WithConfirm     bool     `yaml:"with-confirm"`
	WithInput       bool     `yaml:"witn-input"`
	WithInputFile   bool     `yaml:"with-input-file"`
	WithImageEditor bool     `yaml:"with-image-editor"`
}

// Action dans le menu d'une vue ou formulaire
type Action struct {
	Group       string   // Groupe autorisée à lancer l'action
	Label       string   // label de l'action
	Checkbox    Setters  `yaml:"checkbox"`     // checkbox pour mettre à jour la donnée
	URL         string   `yaml:"url"`          // URL d'appel du formulaire
	SQL         []string `yaml:"sql"`          // les ordres SQL seront exécutées avant l'appel du formulaire
	WithConfirm bool     `yaml:"with-confirm"` // demande de  confirmation
	Hide        bool     // Action non visible
	HideSQL     string   `yaml:"hide-sql"` // requête pour cachée l'action
	Plugin      string   // Fonction Système à appeler nomFonction(p1, p2, ...)
}

// Setters as
type Setters struct {
	GetSQL  string `yaml:"get-sql"`  // requête pour lire la données
	SetSQL  string `yaml:"set-sql"`  // requête pour mettre à jour la données
	AliasDB string `yaml:"alias-db"` // Connecteur base de données
}

// Args paramètres à transmettre lors de l'appel
type Args map[string]string

// Item pour définir un combobox
type Item struct {
	Key   string // valeur dans la base de données
	Label string // label à afficher
}

// MaskList Content d'une card https://fomantic-ui.com/views/card.html
type MaskList struct {
	Header      []string
	Meta        []string
	Description []string
	Extra       []string
}

// Jointure entre tables
type Jointure struct {
	Join   string // la commande du genre : left outer join on field = field
	Column string // colonne retournée par la jointure
}

// CHARGEMENT DU DICTIONNAIRE

var dicoError []string

// Load as
func (c *Portail) Load() ([]string, error) {
	// Raz error
	dicoError = []string{}
	// Read file
	yf := beego.AppConfig.String("dicodir") + "/portail.yaml"
	logs.Info("...load", yf)
	yamlFile, err := ioutil.ReadFile(yf)
	if err != nil {
		logs.Error("yamlFile.Get err", err)
		return dicoError, err
	}
	// chargement de la structure Portail
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		msg := fmt.Sprintf("portail.yaml : [%v]", err)
		dicoError = append(dicoError, msg)
		logs.Error("Unmarshal", msg)
		return dicoError, err
	}
	// Chargement des structures Table
	c.Tables = map[string]*Table{}
	for _, app := range c.Applications {
		logs.Info("...load application:", app.Title)
		for _, tableview := range app.TablesViews {
			if _, ok := c.Tables[tableview.TableName]; !ok {
				var table Table
				table.Load(tableview.TableName)
				c.Tables[tableview.TableName] = &table
			}
		}
	}
	// load liste des fichiers dicodir
	// ouverture du dossier
	f, err := os.Open(beego.AppConfig.String("dicodir"))
	if err != nil {
		msg := fmt.Sprintf("%s : [%v]", beego.AppConfig.String("dicodir"), err)
		dicoError = append(dicoError, msg)
		logs.Error("DICODIR", msg)
		return dicoError, err
	}
	// lecture ds fichiers et dossiers du dossier courant
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		msg := fmt.Sprintf("%s : [%v]", beego.AppConfig.String("dicodir"), err)
		dicoError = append(dicoError, msg)
		logs.Error("DICODIR", msg)
		return dicoError, err
	}
	// tri des fichiers sur le nom
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name() < list[j].Name()
	})

	for _, fileinfo := range list {
		if !fileinfo.IsDir() {
			file := File{}
			file.Base = fileinfo.Name()
			file.Ext = filepath.Ext(fileinfo.Name())
			file.Path = beego.AppConfig.String("dicodir") + "/" + fileinfo.Name()
			c.Files = append(c.Files, file)
		}
	}
	return dicoError, err
}

// Load as
func (c *Table) Load(table string) ([]string, error) {
	yf := beego.AppConfig.String("dicodir") + "/" + table + ".yaml"
	logs.Info("...load", yf)
	// Read file
	yamlFile, err := ioutil.ReadFile(yf)
	if err != nil {
		msg := fmt.Sprintf("%s.yaml : [%v]", table, err)
		dicoError = append(dicoError, msg)
		logs.Error("Unmarshal", msg)
		return dicoError, err
	}
	// chargement de la structure Table
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		logs.Error("Unmarshal", err)
		return dicoError, err
	}
	return dicoError, err
}
