package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/carojaspy/WeatherAPI/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// timeOffset .
const timeOffset = 21600.0

// LapseSeconds .
const LapseSeconds = -300

func getLapse() float64 {
	// Returns the time to compare between 2 dates
	return timeOffset + LapseSeconds
}

// WeatherController operations for Weather
type WeatherController struct {
	beego.Controller
}

// GetWeatherFromProvider .
func GetWeatherFromProvider(city string, country string) (models.WheatherJSON, error) {
	//
	wjson := models.WheatherJSON{}

	// If params was sended, continue with requests
	if city == "" || country == "" {
		// error, incomplete params
		log.Print("error, incomplete params, is needed city and country")
		err := errors.New("error, incomplete params, is needed city and country")
		return wjson, err
	}
	// Example URL to get Weather
	// "http://api.openweathermap.org/data/2.5/weather?q=Bogota,co&appid=8a14e8c7b941473ca2bc48b9e055e5ba"
	querystring := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s", city, country, "8a14e8c7b941473ca2bc48b9e055e5ba")
	// log.Println(querystring)
	resp, err := http.Get(querystring)
	if err != nil {
		// handle error
		log.Println("Error getting Weather from api.openweathermap.org")
		return wjson, err
	}
	if resp.StatusCode != 200 {
		log.Println("Code Status invalid: ", resp.StatusCode)
		str := fmt.Sprintf("404, Not found : City: %v - Country: %v", city, country)
		return wjson, errors.New(str)
	}
	// Closing conection
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading result from api.openweathermap.org")
		return wjson, errors.New("Error reading result from api.openweathermap.org")
	}

	// Building response
	errUn := json.Unmarshal(body, &wjson)
	if errUn != nil {
		fmt.Println("Error in unmarshall: ", errUn)
		return wjson, errors.New("Error Unpacking WeatherAPI info")
	}
	return wjson, nil
} //end GetWeatherFromProvider

// Get ...
// @Title Get
// @Description get Weather by id
// @Success 200 {object} models.Weather
// @Failure 403 :id is empty
// @router / [get]
func (controller *WeatherController) Get() {
	/**/
	log.Print("Handle for Get WeatherController Requests")
	// Trying to retrieve the params from URL
	city := controller.GetString("city")       // Mexico
	country := controller.GetString("country") // mx

	// Calling Handler
	wjson, err := GetWeatherFromProvider(city, country)
	if err != nil {
		controller.CustomAbort(404, err.Error())
	}

	// response := models.FillingResponse(wjson)
	weatherdb := models.Weather{}
	weatherdb = models.FillingDBModel(wjson)

	// Insert Weather to DB
	o := orm.NewOrm()
	qs := o.QueryTable(weatherdb) // return a QuerySeter

	//var weather *models.Weather
	var lastRow models.Weather
	insertRow := false

	// Just One row
	err = qs.OrderBy("-Id").Filter("Location", weatherdb.Location).One(&lastRow)
	if err != nil {
		log.Println(err)
		// No previous Rows, insert to DB
		insertRow = true
	} else {
		log.Println("Sucess getting weather objects from db")
		elapsedTime := time.Until(lastRow.RequestedTime.UTC())

		// log.Println(elapsedTime)
		log.Printf("seconds: %v, LAPSE: %v, %v", elapsedTime.Seconds(), getLapse(), elapsedTime.Seconds() < getLapse())
		if elapsedTime.Seconds() < getLapse() {
			log.Println("New row !!")
			insertRow = true
		} else {
			// still not pass enough time to save another row
			log.Println("You cant insert yet")
			// controller.CustomAbort(404, "You cant insert yet")
		}
	} // End
	// If requests pass all validation, save in Databsae
	if insertRow {
		fmt.Println("Inserting row ...")
		id, err := o.Insert(&weatherdb)
		if err == nil {
			fmt.Println(id)
		}
	}
	controller.Data["json"] = weatherdb
	controller.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description retrieve all Weather objects
// @Success 200 {object} models.WeatherDB
// @Failure 403 :id is empty
// @router /all [get]
func (controller *WeatherController) GetAll() {
	/**/
	log.Println("GetAll method")

	o := orm.NewOrm()
	var weathers []*models.Weather
	num, err := o.QueryTable(new(models.Weather)).All(&weathers)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("SUcess getting weather objects from db")
	log.Println(num)
	controller.Data["json"] = weathers
	controller.ServeJSON()
}
