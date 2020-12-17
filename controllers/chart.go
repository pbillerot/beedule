package controllers

import (
	"fmt"
	"html/template"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pbillerot/beedule/app"
	"github.com/pbillerot/beedule/types"
)

// Quotes table QUOTES
type Quotes struct {
	ID       string `orm:"pk;column(id)"`
	Name     string
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Close1   float64
	Adjclose float64
	Volume   int
}

// Orders table ORDERS
type Orders struct {
	OrdersID        string  `orm:"pk;column(orders_id)"`
	OrdersPtfID     string  `orm:"column(orders_ptf_id)"`
	OrdersBuy       float64 `orm:"column(orders_buy)"`
	OrdersCostPrice float64 `orm:"column(orders_cost_price)"`
}

// ChartController as
type ChartController struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *ChartController) Prepare() {
	if c.GetSession("LoggedIn") != nil {
		if c.GetSession("LoggedIn").(bool) != true {
			c.Ctx.Redirect(302, "/bee/login")
		}
	} else {
		c.Ctx.Redirect(302, "/bee/login")
	}
	appid := c.Ctx.Input.Param(":app")
	// CTRL APPLICATION
	flash := beego.ReadFromRequest(&c.Controller)
	if _, ok := app.Applications[appid]; !ok {
		beego.Error("App not found", c.GetSession("Username").(string), appid)
		c.Ctx.Redirect(302, "/bee/login")
		return
	}
	if !IsInGroup(c.Controller, app.Applications[appid].Group, "") {
		beego.Error("Accès non autorisé", c.GetSession("Username").(string), appid)
		flash.Error("Accès non autorisé")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/bee/login")
		return
	}

	if appid != "" {
		c.Data["TabIcon"] = app.Applications[appid].Image
		c.Data["TabTitle"] = app.Applications[appid].Title
	} else {
		c.Data["TabIcon"] = app.Portail.IconFile
		c.Data["TabTitle"] = app.Portail.Title
	}
	c.Data["AppId"] = appid

	// Contexte de l'application
	c.Data["Application"] = app.Applications[appid]

	// Contexte de la session
	session := types.Session{}
	if c.GetSession("LoggedIn") != nil {
		session.LoggedIn = c.GetSession("LoggedIn").(bool)
	}
	if c.GetSession("Username") != nil {
		session.Username = c.GetSession("Username").(string)
	}
	if c.GetSession("IsAdmin") != nil {
		session.IsAdmin = c.GetSession("IsAdmin").(bool)
	}
	if c.GetSession("Groups") != nil {
		session.Groups = c.GetSession("Groups").(string)
	}
	c.Data["Session"] = &session

	// Identité du framework Beedule
	config := types.Config{}
	config.Appname = beego.AppConfig.String("appname")
	config.Appnote = beego.AppConfig.String("appnote")
	config.Date = beego.AppConfig.String("date")
	config.Icone = beego.AppConfig.String("icone")
	config.Site = beego.AppConfig.String("site")
	config.Email = beego.AppConfig.String("email")
	config.Author = beego.AppConfig.String("author")
	config.Version = beego.AppConfig.String("version")
	config.Theme = beego.AppConfig.String("theme")
	c.Data["Config"] = &config

	// XSRF protection des formulaires
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	// Title
	c.Data["Title"] = config.Appname
	c.Data["Portail"] = &app.Portail
	// Contexte de navigation
	c.Data["From"] = c.Ctx.Input.Cookie("from")
	c.Data["URL"] = c.Ctx.Request.URL.String()

	// Sera ajouté derrière les urls pour ne pas utiliser le cache des images dynamiques
	c.Data["Composter"] = time.Now().Unix()
}

// Get of ChartController
func (c *ChartController) Get() {

	o := orm.NewOrm()
	o.Using("picsou")
	var records []Quotes
	num, err := o.QueryTable("quotes").Filter("ID", "ACA.PA").All(&records)
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
	} else {
		beego.Info(num, "records founds")
	}

	var orders Orders
	err = o.QueryTable("orders").Filter("orders_ptf_id", "ACA.PA").One(&orders)
	if err != nil {
		beego.Error(err)
	}
	// Remplissage du contexte pour le template
	// Dataset as
	type Dataset struct {
		Quotes string  // cotation
		Quotep string  // cotation en %
		Labels string  // axe des dates
		Minp   string  // dataset min en %
		Maxp   string  // dataset max en %
		Min    float64 // Quote max dans le graphique
		Max    float64 // Quote min dans le graphique
		SeuilV float64 // Seuil vente
		SeuilR float64 // Seuil rentabilité
	}
	var dataset Dataset
	start := true
	min := 2000.0
	max := 0.0
	for _, record := range records {
		if start {
			start = false
		} else {
			dataset.Quotes += ","
			dataset.Quotep += ","
			dataset.Labels += ","
			dataset.Minp += ","
			dataset.Maxp += ","
		}
		// Quotes
		dataset.Quotes += fmt.Sprintf("%.2f", record.Open)
		dataset.Quotes += fmt.Sprintf(",%.2f", record.Close)

		// Quotep
		dataset.Quotep += fmt.Sprintf("%.2f", (record.Open-record.Close1)*100/record.Close1)
		dataset.Quotep += fmt.Sprintf(",%.2f", (record.Close-record.Close1)*100/record.Close1)

		// Min
		if record.Low < min {
			min = record.Low
		}
		dataset.Minp += fmt.Sprintf("%.2f", (record.Low-record.Close1)*100/record.Close1)
		dataset.Minp += fmt.Sprintf(",%.2f", (record.Low-record.Close1)*100/record.Close1)

		// Max
		if record.High > max {
			max = record.High
		}
		dataset.Maxp += fmt.Sprintf("%.2f", (record.High-record.Close1)*100/record.Close1)
		dataset.Maxp += fmt.Sprintf(",%.2f", (record.High-record.Close1)*100/record.Close1)

		// Labels
		dataset.Labels += fmt.Sprintf("%s-%s", record.Date[8:], record.Date[5:7])
		dataset.Labels += ",-"
	}
	dataset.SeuilR = orders.OrdersBuy
	optimum := app.Params["__optimum"]
	if f, err := strconv.ParseFloat(optimum, 64); err == nil {
		dataset.SeuilV = orders.OrdersBuy + f*orders.OrdersBuy
	}
	if dataset.SeuilV > max {
		dataset.Max = dataset.SeuilV + dataset.SeuilV*0.01
	} else {
		dataset.Max = max + max*0.01
	}
	dataset.Min = min
	c.Data["Dataset"] = dataset
	c.Data["ID"] = "ACA.PA"
	c.TplName = "chart.html"

}
