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
	// beego.Router("/edf/cp/:app", &controllers.EdfListController{}, "post:Edfcp")
	// beego.Router("/edf/ed/:app", &controllers.EdfListController{}, "get:Edfed;post:Edfed")
	beego.Router("/bee/hugo/list/:app", &controllers.HugoController{}, "get:HugoList;post:HugoList")
	beego.Router("/bee/hugo/image/:app/:key", &controllers.HugoController{}, "get:HugoImage;post:HugoImage")
	beego.Router("/bee/hugo/file/:app/:key", &controllers.HugoController{}, "get:HugoFile;post:HugoFile")
	beego.Router("/bee/hugo/mv/:app/:key", &controllers.HugoController{}, "post:HugoFileMv")
	beego.Router("/bee/hugo/cp/:app/:key", &controllers.HugoController{}, "post:HugoFileCp")
	beego.Router("/bee/hugo/del/:app/:key", &controllers.HugoController{}, "post:HugoFileDel")
	// beego.Router("/edf/mk/:app", &controllers.EdfListController{}, "post:Edfmk")
	// beego.Router("/edf/mv/:app", &controllers.EdfListController{}, "post:Edfmv")
	// beego.Router("/edf/rm/:app", &controllers.EdfListController{}, "post:Edfrm")
	// beego.Router("/edf/to/:app", &controllers.EdfListController{}, "post:Edfto")
	// beego.Router("/edf/vi/:app", &controllers.EdfController{}, "get:Edfvi")

}
