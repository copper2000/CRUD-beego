package main

import (
	_ "beef/models/helper"
	_ "beef/models/register"
	_ "beef/routers"
	"github.com/astaxie/beego"
)

func main() {
	//o := orm.NewOrm()

	// Get a QuerySeter object. product is table name
	//qs := o.QueryTable("product")

	// Can also use object as table name
	//products := new([]*entities.Product)

	//num, err := o.QueryTable(new(entities.Product)).All(products)

	//fmt.Printf("Returned Rows Num: %d %s", num, err)


	beego.Run()
}

