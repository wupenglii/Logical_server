package controllers

//MainController use BaseController
type MainController struct {
	BaseController
}

//Get use index.tpl
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
