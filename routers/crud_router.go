package routers

import (
	"github.com/pbillerot/beedule/controllers"

	beego "github.com/beego/beego/v2/adapter"
)

func init() {
	// PORTAIL liste des applications externes et CRUD

	// LOGIN MANAGER
	beego.Router("/bee/login", &controllers.LoginController{})
	beego.Router("/bee/logout", &controllers.LogoutController{})
	beego.Router("/bee/about", &controllers.AboutController{})

	// Accueuil du CRUD (liste des vues des tables)
	beego.Router("/", &controllers.CrudPortailController{})
	beego.Router("/bee", &controllers.CrudPortailController{})

	// CRUD MANAGER
	beego.Router("/bee/list/:app/:table/:view", &controllers.CrudListController{}, "get:CrudList;post:CrudList")
	beego.Router("/bee/dashboard/:app/:table/:view", &controllers.CrudDashboardController{}, "get:Get;post:Get")
	beego.Router("/bee/view/:app/:table/:view/:id", &controllers.CrudViewController{})
	beego.Router("/bee/add/:app/:table/:view/:form", &controllers.CrudAddController{})
	beego.Router("/bee/edit/:app/:table/:view/:form/:id", &controllers.CrudEditController{})
	beego.Router("/bee/delete/:app/:table/:view/:id", &controllers.CrudDeleteController{})
	beego.Router("/bee/actionv/:app/:table/:view/:action", &controllers.CrudActionViewController{})
	beego.Router("/bee/actionp/:app/:table/:view/:id", &controllers.CrudActionPressController{})
	beego.Router("/bee/actionf/:app/:table/:view/:form/:id/:action", &controllers.CrudActionFormController{})
	beego.Router("/bee/actione/:app/:table/:view/:form/:id/:action", &controllers.CrudActionElementController{})
	beego.Router("/bee/actionx/:app/:table/:view/:id/:action", &controllers.CrudActionAjaxController{})
	beego.Router("/bee/ajax/:app/:table/:view/:form/:action", &controllers.CrudAjaxSqlController{}, "post:Post")
	beego.Router("/bee/share/:appid/:shareid", &controllers.CrudShareController{}, "get:CrudShareApp")
	beego.Router("/bee/search/:app/:table/:view", &controllers.CrudSearchController{}, "post:Post")
	beego.Router("/bee/filter/:app/:table/:view", &controllers.CrudFilterController{}, "post:Post")
	beego.Router("/bee/sort/:app/:table/:view", &controllers.CrudSortController{}, "post:Post")

	// REST API
	//beego.Router("/bee/api/refresh", &controllers.RestController{}, "get:RestRefreshDico")
	//beego.Router("/bee/api/isc", &controllers.RestController{}, "get:RestIsc")

	// EXPLORATEUR DE FICHIERS : EDDI
	beego.Router("/bee/eddy/document/:key", &controllers.EddyController{}, "get:EddyDocument;post:EddyDocument")
	beego.Router("/bee/eddy/log", &controllers.EddyController{}, "get:EddyLog")

}
