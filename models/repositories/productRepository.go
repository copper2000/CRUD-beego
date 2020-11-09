package repositories

import (
	"beef/models/entities"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type ProductRepository struct {
}

func (p ProductRepository) GetAllProduct() []entities.Product {

	var products []entities.Product

	//### USING RAW SQL QUERY BUILDER ###
	/*
		qb, _ := orm.NewQueryBuilder(beego.AppConfig.String("mysqldriver"))

		qb.Select("*").From("product")

		sql := qb.String()

		o := orm.NewOrm()
		o.Raw(sql).QueryRows(&products)
	*/

	// USING QUERY TABLE
	o := orm.NewOrm()
	_, err := o.QueryTable("product").All(&products)

	if err != nil {
		fmt.Println("Query failed")
	}

	return products
}

func (p ProductRepository) InsertProduct(product entities.Product) entities.Product {
	//qb, _ := orm.NewQueryBuilder(beego.AppConfig.String("mysqldriver"))

	//qb.InsertInto("product").Values(product.ProductName, string(product.ProductOrder))

	//sql := qb.String()

	o := orm.NewOrm()

	err := o.Raw("INSERT INTO product(product_name, product_order, created_date, updated_date) VALUES(?, ?, ?, ?)",
		product.ProductName, product.ProductOrder, time.Now(), time.Now()).QueryRow(&product)

	if err != nil {
		fmt.Println("insert fail")
	}
	return product
}

func (p ProductRepository) UpdateProduct(product entities.Product) entities.Product {
	o := orm.NewOrm()

	err := o.Raw("UPDATE product SET product_name = ?, product_order = ?, updated_date = ? WHERE product_id = ?",
		product.ProductName, product.ProductOrder, time.Now(), product.ProductId).QueryRow(&product)

	if err != nil {
		fmt.Println("updated fail")
	}
	return product
}

func (p ProductRepository) DeleteProduct(product entities.Product) bool {
	o := orm.NewOrm()

	err := o.Raw("DELETE FROM product WHERE product_id = ?", product.ProductId).QueryRow(&product)

	if err != nil {
		fmt.Println("updated fail")
		return false
	}
	return true
}

func (p ProductRepository) GetProductById(id int) entities.Product {
	var products entities.Product

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("*").From("product").Where("product_id = ?")

	sql := qb.String()

	o := orm.NewOrm()

	o.Raw(sql, id).QueryRow(&products)

	return products
}

