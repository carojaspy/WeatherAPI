package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/carojaspy/WeatherAPI/controllers"
	"github.com/carojaspy/WeatherAPI/models"
	_ "github.com/carojaspy/WeatherAPI/routers"
	_ "github.com/go-sql-driver/mysql"
)

var channelTask = make(chan models.Task, 10)

// getInterval @TODO Check how to return time in minutes using appconfig settings
func getInterval() time.Duration {
	/*
		interval, err := beego.AppConfig.Int("interval")
		if err != nil {
			//Default 60 minutes
			interval = 60
		}
		log.Println("Setting Interval to ", time.Minute*time.Duration(interval), " minutes, ")
		log.Printf("%T", interval)
	*/
	log.Println("Setting Interval to 1 minutes")
	return 1
}

// WorkerWeather .
func WorkerWeather(id int, o orm.Ormer, wg *sync.WaitGroup) {
	log.Printf("Running WorkerWeather %d: Waiting for task on the channel", id)

	for task := range channelTask {
		log.Printf("WORKER %d : New TASK! Sleeping 10s to emulate hard work \n", id)
		time.Sleep(time.Second * 10)
		result, err := controllers.GetWeather(task.City, task.Country)
		if err != nil {
			log.Printf("WORKER %d : Error Running Task %v - %s : %v \n", id, task.City, task.Country, err.Error())
		} else {
			weatherdb := models.Weather{}
			weatherdb = models.FillingDBModel(result)
			//Trying to insert in to DB
			weatherdb.Save(o)
			log.Printf("WORKER %d : Saved Row ID %d - %v \n", id, weatherdb.Id, weatherdb.Location)
		}
	}
	wg.Done()
}

// CreateTaskWorkerPool .
func CreateTaskWorkerPool(o orm.Ormer, noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go WorkerWeather(i, o, &wg)
	}
	wg.Wait()
}

// WeatherScraping .
func WeatherScraping(o orm.Ormer) {
	log.Println("WeatherScraping")
	// log.Println("Getting Tasks from DB: ...")
	var tasks []*models.Task
	num, err := o.QueryTable(new(models.Task)).All(&tasks)
	if err != nil {
		log.Println(err)
		return
	}
	if num > 0 {
		log.Printf("Sucess getting %v tasks objects from db \n", num)
		for _, task := range tasks {
			channelTask <- *task
			// go EmulateWeatherTask(task.City, task.Country, o)
		}
		log.Println("All task were sended to invoke. Done")
		// At the end
	} else {
		log.Println("No rows, Nothing to do")
	}
}

// CallMeAsync .
func CallMeAsync(o orm.Ormer) {
	go WeatherScraping(o)
	log.Println("CallMeAsync")
	interval := getInterval()
	// Call this Each Minute
	for {
		log.Println("Calling Periodic Tasks ...")
		time.Sleep(time.Minute * interval)
		go WeatherScraping(o)
	}
}

// RunPeriodicTasks .
func RunPeriodicTasks(o orm.Ormer) {
	numWorkers, err := beego.AppConfig.Int("numworkers")
	if err != nil {
		numWorkers = 5
	}
	// Init Pool Of Workers
	go CreateTaskWorkerPool(o, numWorkers)
	// To Call Go routines (Periodic Tasks)
	go CallMeAsync(o)
}

// init Function, Runs Beego and call periodic tasks
func init() {
	// Set UP database on INIT
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Setting Up Uri to connect to Database
	ormURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlhost"),
		beego.AppConfig.String("mysqlport"),
		beego.AppConfig.String("mysqldb"))
	log.Printf(ormURI)
	orm.RegisterDataBase("default", "mysql", ormURI)
}

// main function
func main() {

	// Init ORM
	o := orm.NewOrm()
	// orm.Debug = true
	o.Using("default") // Using default, you can use other database
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// Running  Periodic Tasks
	RunPeriodicTasks(o)

	// Run beego App
	beego.Run()
}
