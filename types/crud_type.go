package types

// BeeConfig Paramètres de config de la webapp
type BeeConfig struct {
	Appname  string
	Appnote  string
	Date     string
	Icone    string
	Site     string
	Email    string
	Author   string
	Version  string
	Theme    string
	HelpDir  string
	HelpPath string
}

// Session Données de session dans le contexte
type Session struct {
	LoggedIn bool
	Username string
	IsAdmin  bool
	Groups   string
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
