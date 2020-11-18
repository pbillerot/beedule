package routers

import (
	"github.com/pbillerot/beedule/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// PORTAIL liste des applications externes et CRUD

	// LOGIN MANAGER
	beego.Router("/crud/login", &controllers.LoginController{})
	beego.Router("/crud/logout", &controllers.LogoutController{})
	beego.Router("/crud/about", &controllers.AboutController{})

	// Accueuil du CRUD (liste des vues des tables)
	beego.Router("/crud", &controllers.CrudPortailController{})

	// CRUD MANAGER
	beego.Router("/crud/list/:app/:table/:view", &controllers.CrudListController{}, "get:CrudList;post:CrudList")
	beego.Router("/crud/view/:app/:table/:view/:id", &controllers.CrudViewController{})
	beego.Router("/crud/add/:app/:table/:view/:form", &controllers.CrudAddController{})
	beego.Router("/crud/edit/:app/:table/:view/:form/:id", &controllers.CrudEditController{})
	beego.Router("/crud/delete/:app/:table/:view/:form/:id", &controllers.CrudDeleteController{})
	beego.Router("/crud/actionv/:app/:table/:view/:action", &controllers.CrudActionViewController{})
	beego.Router("/crud/actionf/:app/:table/:view/:form/:id/:action", &controllers.CrudActionFormController{})
	beego.Router("/crud/actione/:app/:table/:view/:form/:id/:action", &controllers.CrudActionElementController{})

	// EXPLORATEUR DE FICHIERS : HUGO
	beego.Router("/bee/hugo/list/:app", &controllers.HugoController{}, "get:HugoList;post:HugoList")
	beego.Router("/bee/hugo/edit/:app", &controllers.HugoController{}, "get:HugoEditor;post:HugoEditor")
	beego.Router("/bee/hugo/document/:app/:key", &controllers.HugoController{}, "get:HugoDocument;post:HugoDocument")
	beego.Router("/bee/hugo/image/:app/:key", &controllers.HugoController{}, "get:HugoImage;post:HugoImage")
	beego.Router("/bee/hugo/directory/:app/:key", &controllers.HugoController{}, "get:HugoDirectory")
	beego.Router("/bee/hugo/file/:app/:key", &controllers.HugoController{}, "get:HugoFile")
	beego.Router("/bee/hugo/mv/:app/:key", &controllers.HugoController{}, "post:HugoFileMv")
	beego.Router("/bee/hugo/cp/:app/:key", &controllers.HugoController{}, "post:HugoFileCp")
	beego.Router("/bee/hugo/rm/:app/:key", &controllers.HugoController{}, "post:HugoFileRm")
	beego.Router("/bee/hugo/mkdir/:app/:key", &controllers.HugoController{}, "post:HugoFileMkdir")
	beego.Router("/bee/hugo/upload/:app/:key", &controllers.HugoController{}, "post:HugoFileUpload")

}
