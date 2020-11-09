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
	c.Data["DataList"] = repositories.ProductRepository{}.GetAllProduct()
	c.Data["Title"] = "Product Management"
	c.Data["BrandName"] = "BE GROUP"
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