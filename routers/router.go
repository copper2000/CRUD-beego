package routers

import (
	"beef/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/product-list", &controllers.MainController{}, "get:GetProductList")

    beego.Router("/product-add", &controllers.MainController{}, "get:CreateForm")
    beego.Router("/product-create", &controllers.MainController{}, "post:CreateProduct")

	beego.Router("/product-edit/:id", &controllers.MainController{}, "get:EditForm")
	beego.Router("/product-update", &controllers.MainController{}, "post:UpdateProduct")
}
