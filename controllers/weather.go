package controllers

import (
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/carojaspy/WeatherAPI/models"
)

// WeatherController operations for Weather
type WeatherController struct {
	beego.Controller
}

// GetWeather .
func GetWeather(city string, country string) (models.WheatherJSON, error) {
	provider := beego.AppConfig.String("weatherprovider")
	wjson := models.WheatherJSON{}
	if provider == "APIProvider" {
		log.Println("Getting Weather from APIProvider")
		return models.GetWeatherFromAPI(city, country)
	} else if provider == "FileProvider" {
		return models.GetWeatherFromFile(city, country)
	}
	return wjson, errors.New("weatherprovider variable is not set, check your app.conf")
}

// WeatherTask  Get the Weather and persists to Database  (SYNC)
func WeatherTask(city string, country string, o orm.Ormer) {
	// log.Printf("Done Task %v - %s", city, country)
	result, err := GetWeather(city, country)
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

// Get ...
// @Title Get
// @Description get Weather by id
// @Success 200 {object} models.Weather
// @Failure 403 :id is empty
// @router / [get]
func (controller *WeatherController) Get() {
	/**/
	// Trying to retrieve the params from URL
	city := controller.GetString("city")
	country := controller.GetString("country")

	// Calling Handler
	wjson, err := GetWeather(city, country)
	if err != nil {
		controller.CustomAbort(404, err.Error())
	}

	// response := models.FillingResponse(wjson)
	weatherdb := models.Weather{}
	weatherdb = models.FillingDBModel(wjson)

	o := orm.NewOrm()
	qs := o.QueryTable(weatherdb) // return a QuerySeter
	qs.Filter("Location", weatherdb.Location)

	//	Check if is a valid new row ( >300 seconds)
	if err := weatherdb.IsValid(o); err == nil {
		//Trying to to DB
		weatherdb.Save(o)
		// Saving the request
		req := models.RequestWeather{City: city, Country: country, RequestedTime: time.Now()}
		req.Save(o)
	} else {
		log.Println(err.Error())
	}
	controller.Data["json"] = weatherdb
	controller.ServeJSON()
} //End Get Method

// GetAll ...
// @Title GetAll
// @Description retrieve all Weather objects
// @Success 200 {object} models.WeatherDB
// @Failure 403 :id is empty
// @router /all [get]
func (controller *WeatherController) GetAll() {
	/**/
	log.Println("GetAll Weather Controller")
	o := orm.NewOrm()
	var weathers []*models.Weather
	num, err := o.QueryTable(new(models.Weather)).All(&weathers)
	if err != nil {
		log.Println(err)
		controller.CustomAbort(404, err.Error())
	}
	log.Println("SUcess getting weather objects from db")
	log.Println(num)
	controller.Data["json"] = weathers
	controller.ServeJSON()
} // End GetAll Method
