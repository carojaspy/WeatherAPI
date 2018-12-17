package controllers

import (
	"fmt"
	"log"
	"github.com/carojaspy/WeatherAPI/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// SchedulerController operations for Scheduler
type SchedulerController struct {
	beego.Controller
}


/*
Task to get the Weather 
*/
func callPeriodicTask(city string, country string){
	//GetWeatherFromProvider(city string, country string)
}



// GetAll ...
// @Title GetAll
// @Description get Scheduler
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Scheduler
// @Failure 403
// @router / [get]
func (controller *SchedulerController) GetAll() {
	fmt.Println("GetAll routine controller")


}

// Put ...
// @Title Put
// @Description update the Scheduler
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Scheduler	true		"body for Scheduler content"
// @Success 200 {object} models.Scheduler
// @Failure 403 :id is not int
// @router /weather [put]
func (controller *SchedulerController) Put() {
	fmt.Println("Put new routine controller")
	o := orm.NewOrm()

	// Trying to retrieve the params from URL
	city := controller.GetString("city") // Mexico
	country := controller.GetString("country") // mx
	log.Printf("%v - %v ", city, country)

	// Saving info in Task model
	t := models.Task{City:city, Country:country, IsActive:true}
	log.Printf("%v", t)
	if err:= t.IsValid(o); err!=nil {
		log.Println("Invalid Task to persist to Database: ", err.Error())
	} else {
		log.Println("Valid Task, persists to DB: ")
		log.Printf("%v - %v ", city, country)
		t.Save(o)
	}
	controller.Data["json"] = t
	controller.ServeJSON()
}// EndPut Method

