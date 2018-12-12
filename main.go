package main

import (
	_ "WeatherAPI/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// Set UP database on INIT
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// "USER:PASS@/DBNAME?charset=utf8"
	ormURI := fmt.Sprintf("%s:%s@/%s?charset=utf8",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqldb"))
	orm.RegisterDataBase("default", "mysql", ormURI)
}
func main() {
	// Init ORM

	o := orm.NewOrm()
	orm.Debug = true
	o.Using("default") // Using default, you can use other database

	/*
	CREATE SQL SCHEMA
		// Database alias.
		name := "default"
		// Drop table and re-create.
		force := true
		// Print log.
		verbose := true
		// Error.
		err := orm.RunSyncdb(name, force, verbose)
		if err != nil {
			fmt.Println(err)
		}
	*/
		
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
