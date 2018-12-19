package main

import (
	_ "github.com/carojaspy/WeatherAPI/routers"
	"github.com/carojaspy/WeatherAPI/models"
	"github.com/carojaspy/WeatherAPI/controllers"
	"fmt"
	"log"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// EmulateWeatherTask . 
func EmulateWeatherTask(city string, country string, o orm.Ormer) {
	// log.Printf("Done Task %v - %s", city, country)
	result, err := controllers.GetWeatherFromProvider(city, country)
	if err != nil {
		log.Printf("Error Running Task %v - %s : %v", city, country, err.Error())
		return 
	}
	// response := models.FillingResponse(wjson)
	weatherdb := models.Weather{}
	weatherdb = models.FillingDBModel(result)
	//Trying to to DB
	weatherdb.Save(o)
	// Continue 
}

// WeatherScraping . 
func WeatherScraping(o orm.Ormer){
	log.Println("WeatherScraping")
	log.Println("Getting Tasks from DB: ...")

	//	@TODO	Move the FetchAll objects to a Interface Method
	var tasks []*models.Task
	num, err := o.QueryTable(new(models.Task)).All(&tasks)
	if err != nil {
		log.Println(err)
		return
	}
	if num > 0 {
		log.Printf("Sucess getting %v tasks objects from db", num)
		for _, task := range tasks {
			time.Sleep(time.Second*1)
			go EmulateWeatherTask(task.City, task.Country, o)
		}
		log.Println("All task were sended to invoke. Done")
		// At the end 	
	} else {
		log.Println("No rows, Nothing to do")
	}
}

// CallMeAsync . 
func CallMeAsync(o orm.Ormer){
	go WeatherScraping(o)
	log.Println("CallMeAsync")
	// Call this Each Minute
	for {
		time.Sleep(time.Minute)
		go WeatherScraping(o)
	}
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

	// To Call Go routines (Periodic Tasks)
	// go CallMeAsync(o)
	beego.Run()
}

// ## docker-compose run app $COMANDO 
