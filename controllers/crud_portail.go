package controllers

// CrudPortailController as
type CrudPortailController struct {
	loggedRouter
}

// Get CrudPortailController
func (c *CrudPortailController) Get() {
	c.SetSession("from", "/crud")
	setContext(c.Controller)
	c.TplName = "crud_portail.html"
}
