package entities

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)

type Product struct {
	ProductId    int       `json:"product_id" orm:"pk;auto" form:"productId"`
	ProductName  string    `json:"product_name" orm:"null" form:"productName"`
	ProductOrder int       `json:"product_order" orm:"null" form:"productOrder"`
	CreatedDate  time.Time `json:"created_date" orm:"null"`
	UpdatedDate  time.Time `json:"updated_date" orm:"null"`
}
