package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

const (
	productList   = "product/list.html"
	productDetail = "product/detail"
	productAdd    = "product/add"
	productEdit   = "product/edit"
	productDelete = "product/delete"
)

func (c *DefaultController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
