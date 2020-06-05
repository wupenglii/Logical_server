package controllers

import (
	"github.com/astaxie/beego"
)

//MainController use BaseController
type MainController struct {
	beego.Controller
}

func (c *MainController) InitData() {
	c.TplName = "index.tpl"
}
