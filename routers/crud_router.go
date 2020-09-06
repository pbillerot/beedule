package routers

import (
	"github.com/pbillerot/beedule/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// PORTAIL liste des applications externes et CRUD

	// Aut
	// beego.Router("/auth", &controllers.AuthRouter{})

	// LOGIN MANAGER
	beego.Router("/crud/login", &controllers.LoginController{})
	beego.Router("/crud/logout", &controllers.LogoutController{})
	beego.Router("/crud/about", &controllers.AboutController{})

	// Accueuil du CRUD (liste des vues des tables)
	beego.Router("/crud", &controllers.CrudPortailController{})

	// CRUD MANAGER
	beego.Router("/crud/list/:app/:table/:view", &controllers.CrudListController{})
	beego.Router("/crud/view/:app/:table/:view/:id", &controllers.CrudViewController{})
	beego.Router("/crud/add/:app/:table/:view/:form", &controllers.CrudAddController{})
	beego.Router("/crud/edit/:app/:table/:view/:form/:id", &controllers.CrudEditController{})
	beego.Router("/crud/delete/:app/:table/:view/:id", &controllers.CrudDeleteController{})
	beego.Router("/crud/action/:app/:table/:view/:action", &controllers.CrudActionController{})

}