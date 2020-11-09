package helper

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver(beego.AppConfig.String("mysqldriver"), orm.DRMySQL)

	var err = orm.RegisterDataBase(beego.AppConfig.String("mysqlalias"), beego.AppConfig.String("mysqldriver"),
		beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/"+
			beego.AppConfig.String("mysqldb")+"?charset=utf8")
	if err != nil {
		fmt.Println("DB connect fail")
	} else {
		fmt.Println("DB connect success")
	}
}
