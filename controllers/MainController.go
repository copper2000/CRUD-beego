package controllers

import (
	"beef/models/entities"
	"beef/models/repositories"
	"fmt"
	"strconv"
)

type MainController struct {
	DefaultController
}

func (c *MainController) GetProductList() {
	var size = c.Ctx.Input.Param(":pageSize")
	var index = c.Ctx.Input.Param(":pageIndex")
	pageSize, _ := strconv.Atoi(size)
	pageIndex, _ := strconv.Atoi(index)

	c.Data["Title"] = "Product Management"
	c.Data["BrandName"] = "BE GROUP"

	c.Data["CurrentPageSize"] = int64(pageSize)
	c.Data["DataList"], c.Data["total"] = repositories.ProductRepository{}.PagingProduct(pageIndex, pageSize)
	c.TplName = productList
}

func (c *MainController) CreateForm() {
	c.TplName = "product/add.html"
}

func (c *MainController) CreateProduct() {
	prod := entities.Product{}

	err := c.ParseForm(&prod)

	repositories.ProductRepository{}.InsertProduct(prod)

	if err != nil {
		fmt.Println(err)
	}

	c.Redirect("product-list", 302)

}

func (c *MainController) EditForm() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	c.Data["product"] = repositories.ProductRepository{}.GetProductById(id)

	c.TplName = "product/edit.html"
}

func (c *MainController) UpdateProduct() {
	prod := entities.Product{}

	err := c.ParseForm(&prod)

	repositories.ProductRepository{}.UpdateProduct(prod)

	if err != nil {
		fmt.Println(err)
	}

	c.Redirect("product-list", 302)
}

func (c *MainController) DeleteForm() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	c.Data["product"] = repositories.ProductRepository{}.GetProductById(id)

	c.TplName = "product/remove.html"
}

func (c *MainController) DeleteProduct() {
	prod := entities.Product{}

	err := c.ParseForm(&prod)

	repositories.ProductRepository{}.DeleteProduct(&prod)

	if err != nil {
		fmt.Println(err)
	}

	c.Redirect("product-list", 302)
}

func (c *MainController) PagingProduct() {
	var size = c.Ctx.Input.Param(":pageSize")
	var index = c.Ctx.Input.Param(":pageIndex")

	pageSize, _ := strconv.Atoi(size)
	pageIndex, _ := strconv.Atoi(index)
	c.Data["DataPaging"], c.Data["total"] = repositories.ProductRepository{}.PagingProduct(pageIndex, pageSize)
	c.TplName = "product/remove.html"

}