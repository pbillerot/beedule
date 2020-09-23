package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pbillerot/beedule/app"
	"golang.org/x/crypto/bcrypt"
)

// loggedRouter implements global settings for all other routers.
type loggedRouter struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *loggedRouter) Prepare() {
	if c.GetSession("LoggedIn") != nil {
		if c.GetSession("LoggedIn").(bool) != true {
			c.Ctx.Redirect(302, "/crud/login")
		}
	} else {
		c.Ctx.Redirect(302, "/crud/login")
	}
	appid := c.Ctx.Input.Param(":app")
	if appid != "" {
		c.Data["TabIcon"] = app.Applications[appid].Image
		c.Data["TabTitle"] = app.Applications[appid].Title
	} else {
		c.Data["TabIcon"] = app.Portail.IconFile
		c.Data["TabTitle"] = app.Portail.Title
	}
}

// adminRouter implements global settings for all other routers.
type adminRouter struct {
	beego.Controller
}

// Prepare implements Prepare method for loggedRouter.
func (c *adminRouter) Prepare() {
	// beego.Debug("adminRouter")
	if c.GetSession("LoggedIn") != nil {
		if c.GetSession("LoggedIn").(bool) != true {
			c.Ctx.Redirect(302, "/crud/login")
		}
		if c.GetSession("IsAdmin").(bool) != true {
			c.Ctx.Redirect(302, "/crud/login")
		}
	} else {
		c.Ctx.Redirect(302, "/crud/login")
	}
	appid := c.Ctx.Input.Param(":app")
	if appid != "" {
		c.Data["TabIcon"] = app.Applications[appid].Image
		c.Data["TabTitle"] = app.Applications[appid].Title
	} else {
		c.Data["TabIcon"] = app.Portail.IconFile
		c.Data["TabTitle"] = app.Portail.Title
	}
}

// LoginController as
type LoginController struct {
	beego.Controller
}

// LogoutController as
type LogoutController struct {
	beego.Controller
}

// AboutController as
type AboutController struct {
	loggedRouter
}

// Get AboutController
func (c *AboutController) Get() {
	setContext(c.Controller)
	c.Data["Title"] = "A propos de Beedule"
	c.TplName = "crud_about.html"
}

// Get of LoginController
func (c *LoginController) Get() {
	setContext(c.Controller)
	c.Data["Title"] = "Beedule"
	c.TplName = "crud_login.html"
}

// Post of LoginController
func (c *LoginController) Post() {
	c.Data["username"] = c.GetString("username")

	flash := beego.NewFlash()

	user, err := GetUser(c.GetString("username"))
	if err != nil {
		flash.Error("Compte ou mot de passe erroné")
		flash.Store(&c.Controller)
		setContext(c.Controller)
		c.TplName = "crud_login.html"
		return
	}
	// Comparaison mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.GetString("password")))
	if err != nil {
		flash.Error("Compte ou mot de passe erroné")
		flash.Store(&c.Controller)
		setContext(c.Controller)
		c.TplName = "crud_login.html"
		return
	}
	// C'est OK enregistrement du compte dans la session
	c.SetSession("LoggedIn", true)
	c.SetSession("Username", user.Username)
	c.SetSession("Groups", user.Groupes)
	if user.IsAdmin == 1 {
		c.SetSession("IsAdmin", true)
	} else {
		c.SetSession("IsAdmin", false)
	}
	c.Ctx.Redirect(302, "/crud")
}

// Get of LogoutController
func (c *LogoutController) Get() {
	// Suppression du user dans le session
	c.SetSession("LoggedIn", false)
	c.DelSession("Username")
	c.DelSession("IsAdmin")
	c.DelSession("Groups")
	setContext(c.Controller)
	c.Ctx.Redirect(302, "/crud/login")
}

// User as
type User struct {
	Username string `orm:"pk;column(user_name);size(20)"`
	Password string `orm:"column(user_password);size(100)"`
	Email    string `orm:"column(user_email);size(100)"`
	IsAdmin  int    `orm:"column(user_is_admin)"`
	Groupes  string `orm:"column(user_groupes)"`
}

// TableName de la table User
func (u *User) TableName() string {
	return "users"
}
func init() {
	orm.RegisterModel(new(User))
}

// GetUser fournit le user
func GetUser(username string) (User, error) {
	o := orm.NewOrm()
	o.Using(app.Tables["users"].AliasDB)
	user := User{Username: username}
	err := o.Read(&user, "Username")

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(user.Username)
	}
	return user, err
}
