package controllers

import (
	"time"
	"github.com/carojaspy/WeatherAPI/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


// WeatherController operations for Weather
type WeatherController struct {
	beego.Controller
}


// GetWeatherFromProvider .
func GetWeatherFromProvider(city string, country string) (models.WheatherJSON ,error){
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
}//end GetWeatherFromProvider


// Get ...
// @Title Get 
// @Description get Weather by id
// @Success 200 {object} models.Weather
// @Failure 403 :id is empty
// @router / [get]
func (controller *WeatherController) Get() {
	/**/
	log.Print("Handle for Get WeatherController Requests")
	log.Print(time.Now())
	// Trying to retrieve the params from URL
	city := controller.GetString("city") // Mexico
	country := controller.GetString("country") // mx
	// city := "Mexico" // Mexico
	// country := "mx" // mx

	// Calling Handler
	wjson, err := GetWeatherFromProvider(city, country)
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
	} else {
		log.Println(err.Error())
	}
	controller.Data["json"] = weatherdb
	controller.ServeJSON()
}//End Get Method


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
}// End GetAll Method
