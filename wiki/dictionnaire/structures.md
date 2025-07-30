## Modèles des structures
```
// Portail as portail.yaml
type Portail struct {
	DicoDir      []string `yaml:"dico-dir"` // répertoire du dictionnaires des applications
	Title        string
	Info         string
	IconFile     string                 `yaml:"icon-file"`
	Applications map[string]Application // working
}

// File as les fichiers du dictionnaire dicodir
type File struct {
	Base string // le filename
	Name string // le filename sans l'extension
	Ext  string // l'extensions
	Path string //le chemin complet
}

// Application as
type Application struct {
	AppID          string            `yaml:"app-id"`   // id de l'application
	AliasDB        string            `yaml:"alias-db"` // par défaut
	Batman         string            `yaml:"batman"`   // par défaut
	Title          string            // Label de l'application
	Image          string            // image de l'application affichée au niveau du portail
	IconFile       string            `yaml:"icon-file"` // url de l'icône de l'application
	IconName       string            `yaml:"icon-name"` // Icône https://semantic-ui.com/elements/icon.html
	Group          string            // groupe utilisateur habilité à accéder à l'application
	DicoDir        string            `yaml:"dico-dir"` // working
	Parameters     map[string]string // paramètres de l'application accessible par {__param1}
	Path           string            // Path ou URL de l'application externe
	Tables         map[string]*Table // Tables de l'application chargées par portail.load working
	Files          []File            // liste des fichiers trouvés dans dicodir de l'application working
	Menu           []TableView       // menu des Vues de l'application
	TasksTableName string            `yaml:"tasks-table-name"` // Nom de la table des Tâches planifiées
	Wiki           string            // répertoire du wiki de l'application
}

// Menu as
type Menu []TableView

// TableView as
type TableView struct {
	TableID  string `yaml:"table-id"`
	ViewID   string `yaml:"view-id"`
	InFooter bool   `yaml:"in-footer"`
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
	Actions       []Action          // bouton d'actions
	Args          map[string]string // Args pour passer des arguments à une vue
	AjaxSQL       string            `yaml:"ajax-sql"` // query sql pour ramenener des données dans le formulaire
	Class         string            // calculé Class du texte dans la cellule https://fomantic-ui.com/collections/table.html
	ClassSqlite   string            `yaml:"class-sqlite"` // SQL pour alimenter Class error warning info green blue
	ColAlign      string            `yaml:"col-align"`    //
	ColNoWrap     bool              `yaml:"col-nowrap"`   // nowrap de la colonne
	Dataset       map[string]string `yaml:"dataset"`      // Dataset pour un Chartjs ou pour passer des arguments à une vue ou à une "ajax-sql"
	Default       string            // Valeur par défaut (macro possible)
	DefaultSqlite string            `yaml:"default-sqlite"` // Ordre SQL qui retournera la colonne pour alimenter Default
	Error         string            // contiendra "error" si le champ est en erreur de saisie
	Format        string            // "%3.2f %%" "%3.2f €" date time
	FormatSqlite  string            `yaml:"format-sqlite"`  // select strftime('%H:%M:%S', {Milliseconds}/1000, 'unixepoch')
	ComputeSqlite string            `yaml:"compute-sqlite"` // formule de calcul de Value en SQL dans VIEW EDIT ADD (pas dans LIST)
	Group         string            // Groupe autorisé à accéder à cette rubrique - Si owner c'est l'enregistreement qui sera protégé
	Help          string            // TODO aide sur la rubrique
	Hide          bool              // élémnt caché dans la vue ou formulaire
	HideSqlite    string            `yaml:"hide-sqlite"`    // TODO cachée si condition
	HideOnMobile  bool              `yaml:"hide-on-mobile"` // La colonne dans une vue sera cachée sur Mobile
	ID            string            // id de la rubrique calculé = nom de la rubrique
	IconName      string            `yaml:"icon-name"` // icone d'une card par exemple
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
	Protected     bool              // Est en mise à jour mais protégé en saisie
	ReadOnly      bool              `yaml:"read-only"` // Lecteur seule
	Record        orm.Params        // l'enregistrement  valorisant la rubrique
	Required      bool              // obligatoire
	SortDirection string            `yaml:"sort-direction"` // "", ascending, ou descending pour demander un tri à la requête sql
	SQLout        string            `yaml:"sql-out"`        // Valeur à enregistrer dans la base de données (zone calculée par le beedule)
	StyleSqlite   string            `yaml:"style-sqlite"`   // style de la cellule
	Type          string            // Type : amount button card chart checkbox counter date email float image list number password pdf percent tag tel text textarea time radio url
	Width         string            // largeur s m l xl xxl max 150px 360px 450px 600px 750px 100% dans view et edit	WithSum       bool              // dans une table calcule la somme des valeurs
	WithScript    string            `yaml:"with-script"` // javascript de présentation
	WithSum       bool              `yaml:"with-sum"`    // dans une table calcule la somme des valeurs
}

// View Modele Vue d'une table
type View struct {
	ActionPress    Action             `yaml:"action-press"` // Action sur appui long sur l'article
	Actions        []Action           // Action sur la vue (ordres sql)
	Card           CardList           // Masque html d'une ligne dans la vue
	ClassSqlite    string             `yaml:"class-sqlite"` // couleur theme de la ligne
	Deletable      bool               // Suppression fiche autorisée
	Elements       map[string]Element // Eléments à récupérer de la base de données
	Filters        []string           // liste de nom d'élément sur la vue
	FooterSQL      string             `yaml:"footer-sql"` // requête sur la table courante
	FormAdd        string             `yaml:"form-add"`   // Formulaire d'ajout
	FormEdit       string             `yaml:"form-edit"`  // Formulaire d'édition
	FormView       string             `yaml:"form-view"`  // Masque de visualisation d'un enregistrement
	Group          string             // groupe qui peut accéder à la vue
	Hide           bool               // Vue cachée dans le menu
	HideOnMobile   bool               `yaml:"hide-on-mobile"` // Vue cachée dur mobile
	IconName       string             `yaml:"icon-name"`      // nom de l'icone
	Limit          int                // pour limiter le nbre de ligne dans la vue
	OrderBy        string             `yaml:"order-by"`       // Tri des données SQL
	PostSQL        []string           `yaml:"post-sql"`       // Ordre exécutée après la suppression si OK
	PreUpdateSQL   []string           `yaml:"pre-update-sql"` // requêtes SQL avant l'affichage
	Search         string             // calculé sql
	StyleSqlite    string             `yaml:"style-sqlite"` // style de la ligne
	Title          string             // Titre de la vue
	Type           string             // type de vue : card(default),image,table
	Where          string             // Condition SQL
	Width          string             // largeur s m l xl xxl max
	WithLineNumber bool               `yaml:"with-line-number"` // list.table n° de ligne en 1ère colonne
	WithSum        bool               `yaml:"with-sum"`
}

// Form formulaire
type Form struct {
	Actions     []Action           // Action appel d'un formulaire ou exécution d'une requête SQL
	Title       string             // Titre du formulaire
	Group       string             // groupe qui peut accéder au formulaire
	HideSubmit  bool               `yaml:"hide-submit"` // pour caher le bouton valider
	IconName    string             `yaml:"icon-name"`   // nom de l'icone
	Elements    map[string]Element // Eléments à récupérer de la base de données
	CheckSqlite []string           `yaml:"check-sqlite"` // retourne le libellé des erreurs lors du contrôle des rubriques
	PostSQL     []string           `yaml:"post-sql"`     // Ordre exécutée après la validation si contrôle OK
}

// Params paramètres d'un élément
type Params struct {
	Form         string   // card: form à ouvrir
	Header       []string // card pour image
	Description  []string // card pour image
	Meta         []string // card pour image
	Extra        []string // card pour image
	URL          string   `yaml:"url"`
	Src          string   // card: src de l'image
	SQL          []string `yaml:"sql"`
	Table        string   // card:
	Target       string   // target si URL
	Title        string   // title sur une image
	View         string   // card:
	Where        string   // card: + params.table + params.view
	WithConfirm  bool     `yaml:"with-confirm"`
	WithoutFrame bool     `yaml:"without-frame"` // card sans cadre
}

// Action dans le menu d'une vue ou formulaire
type Action struct {
	Group       string   // Groupe autorisée à lancer l'action
	Label       string   // label de l'action
	Checkbox    Setters  `yaml:"checkbox"`     // checkbox pour mettre à jour la donnée
	IconName    string   `yaml:"icon-name"`    // icone de l'action dans le menu
	URL         string   `yaml:"url"`          // URL d'appel du formulaire
	SQL         []string `yaml:"sql"`          // les ordres SQL seront exécutées avant l'appel du formulaire
	WithConfirm bool     `yaml:"with-confirm"` // demande de  confirmation
	Hide        bool     // Action non visible
	HideSqlite  string   `yaml:"hide-sqlite"` // requête pour cachée l'action
	Plugin      string   // Fonction Système à appeler nomFonction(p1, p2, ...)
	Shell       []string // commande dans le shell
}

// Setters as
type Setters struct {
	GetSQL  string `yaml:"get-sql"`  // requête pour lire la données
	SetSQL  string `yaml:"set-sql"`  // requête pour mettre à jour la données
	AliasDB string `yaml:"alias-db"` // Connecteur base de données
}

// Args paramètres à transmettre lors de l'appel
type Args map[string]string

// Item pour définir une list
type Item struct {
	Key   string // valeur dans la base de données
	Label string // label à afficher
}

// CardList Content d'une card https://fomantic-ui.com/views/card.html
// ..Left ..Right -> pas de label-long affiché
type CardList struct {
	Header       []string
	HeaderRight  []string `yaml:"header-right"`
	HeaderAction Action   `yaml:"header-action"` // bouton menu d'action
	Meta         []string
	MetaLeft     []string `yaml:"meta-left"`
	MetaRight    []string `yaml:"meta-right"`
	Description  []string
	Extra        []string
	ExtraLeft    []string `yaml:"extra-left"`
	ExtraRight   []string `yaml:"extra-right"`
	Footer       []string
	FooterLeft   []string `yaml:"footer-left"`
	FooterRight  []string `yaml:"footer-right"`
	FooterAction Action   `yaml:"footer-action"` // bouton menu d'action
}

// Jointure entre tables
type Jointure struct {
	Join   string // la commande du genre : left outer join on field = field
	Column string // colonne retournée par la jointure
}

```
