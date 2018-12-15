package main

import (
	_ "github.com/carojaspy/WeatherAPI/routers"
	"fmt"
	"log"
	"time"
	"math/rand"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func EmulateWeatherTask(id int) {
	log.Printf("Done Task %v", id)
}

func WeatherScraping(){
	log.Println("WeatherScraping")
	log.Println("Getting Tasks from DB: ...")
	time.Sleep(time.Second)
	for i:=1; i<10; i++ {
		time.Sleep(time.Second*1)
		go EmulateWeatherTask(rand.Intn(100))
	}
	log.Println("All task were sended to invoke. Done")
	// At the end 
}

// CallMeAsync . 
func CallMeAsync(){
	go WeatherScraping()
	log.Println("CallMeAsync")
	// Call this Each Minute
	time.Sleep(time.Minute)
	go WeatherScraping()
}


func init() {
	// Set UP database on INIT
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// "$USER:PASS@tcp($HOST:$PORT)/DBNAME",	
	// "%s:%s@tcp(%s:%s)/%s"
	log.Printf("================ %v - %v - %v \n",  beego.AppConfig.String("mysqluser"), beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqldb"))

	// ormURI := fmt.Sprintf("%s:%s@/%s?charset=utf8",
	ormURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
	beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlhost"),
		beego.AppConfig.String("mysqlport"),				
		beego.AppConfig.String("mysqldb"))
	log.Printf(ormURI)
	orm.RegisterDataBase("default", "mysql", ormURI)
	// orm.RegisterDataBase("default", "mysql", "root:@tcp(weather_api_db:3306)/weatherapi?charset=utf8")

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
	go CallMeAsync()
	beego.Run()
}

// ## docker-compose run app $COMANDO 
