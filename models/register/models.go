package register

import (
	"beef/models/entities"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(entities.Product))

	// drop table and re-create
	force := false

	// print log.
	verbose := true

	// error
	err := orm.RunSyncdb(beego.AppConfig.String("mysqlalias"), force, verbose)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrate done!")
	}

}
