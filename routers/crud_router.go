package routers

import (
	"github.com/pbillerot/beedule/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// PORTAIL liste des applications externes et CRUD

	// LOGIN MANAGER
	beego.Router("/bee/login", &controllers.LoginController{})
	beego.Router("/bee/logout", &controllers.LogoutController{})
	beego.Router("/bee/about", &controllers.AboutController{})

	// Accueuil du CRUD (liste des vues des tables)
	beego.Router("/bee", &controllers.CrudPortailController{})

	// CRUD MANAGER
	beego.Router("/bee/list/:app/:table/:view", &controllers.CrudListController{}, "get:CrudList;post:CrudList")
	beego.Router("/bee/view/:app/:table/:view/:id", &controllers.CrudViewController{})
	beego.Router("/bee/add/:app/:table/:view/:form", &controllers.CrudAddController{})
	beego.Router("/bee/edit/:app/:table/:view/:form/:id", &controllers.CrudEditController{})
	beego.Router("/bee/delete/:app/:table/:view/:form/:id", &controllers.CrudDeleteController{})
	beego.Router("/bee/actionv/:app/:table/:view/:action", &controllers.CrudActionViewController{})
	beego.Router("/bee/actionf/:app/:table/:view/:form/:id/:action", &controllers.CrudActionFormController{})
	beego.Router("/bee/actione/:app/:table/:view/:form/:id/:action", &controllers.CrudActionElementController{})

	// EXPLORATEUR DE FICHIERS : HUGO
	beego.Router("/bee/hugo/list/:app", &controllers.HugoController{}, "get:HugoList;post:HugoList")
	beego.Router("/bee/hugo/edit/:app", &controllers.HugoController{}, "get:HugoEditor;post:HugoEditor")
	beego.Router("/bee/hugo/document/:app/:key", &controllers.HugoController{}, "get:HugoDocument;post:HugoDocument")
	beego.Router("/bee/hugo/image/:app/:key", &controllers.HugoController{}, "get:HugoImage;post:HugoImage")
	beego.Router("/bee/hugo/pdf/:app/:key", &controllers.HugoController{}, "get:HugoPdf")
	beego.Router("/bee/hugo/directory/:app/:key", &controllers.HugoController{}, "get:HugoDirectory")
	beego.Router("/bee/hugo/file/:app/:key", &controllers.HugoController{}, "get:HugoFile")
	beego.Router("/bee/hugo/mv/:app/:key", &controllers.HugoController{}, "post:HugoFileMv")
	beego.Router("/bee/hugo/cp/:app/:key", &controllers.HugoController{}, "post:HugoFileCp")
	beego.Router("/bee/hugo/rm/:app/:key", &controllers.HugoController{}, "post:HugoFileRm")
	beego.Router("/bee/hugo/mkdir/:app/:key", &controllers.HugoController{}, "post:HugoFileMkdir")
	beego.Router("/bee/hugo/upload/:app/:key", &controllers.HugoController{}, "post:HugoFileUpload")

	// REST API
	// beego.Router("/bee/api/isc", &controllers.RestController{}, "get:RestIsc")
	// beego.Router("/bee/api/put/log/:url", &controllers.RestController{}, "get:RestPutLog")

}
