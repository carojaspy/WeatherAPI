package controllers

import (
	"WeatherAPI/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
)

// WeatherController operations for Weather
type WeatherController struct {
	beego.Controller
}

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
	city := controller.GetString("city") // Mexico
	country := controller.GetString("country") // mx
	if city == "" || country == "" {
		// error, incomplete params
		log.Print("error, incomplete params, is needed city and country")
		controller.Abort("401")
	}
	// If params was sended, continue with requests

	// Example URL to get Weather
	// "http://api.openweathermap.org/data/2.5/weather?q=Bogota,co&appid=8a14e8c7b941473ca2bc48b9e055e5ba"
	querystring := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s", city, country, "8a14e8c7b941473ca2bc48b9e055e5ba")
	// log.Println(querystring)
	resp, err := http.Get(querystring)
	if err != nil {
		// handle error
		log.Println("Error getting Weather from api.openweathermap.org")
		controller.Abort("404")
	}
	if resp.StatusCode != 200 {
		log.Println("Code Status invalid: ", resp.StatusCode)
		controller.Abort("404")
	}
	// Closing conection
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading result from api.openweathermap.org")
		controller.Abort("404")
	}

	// Building response
	wjson := models.WheatherJSON{}
	errUn := json.Unmarshal(body, &wjson)
	if errUn != nil {
		fmt.Println("Error in unmarshall: ", errUn)
		controller.Abort("401")
	}
	response := models.FillingResponse(wjson)
	controller.Data["json"] = response
	controller.ServeJSON()
}
