package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

/*
	Controllers to handle requests
	this also contains logic to connect with
	openweathermap.org API

*/
const APIKEY = "8a14e8c7b941473ca2bc48b9e055e5ba"

// Operations about object
type WeatherController struct {
	beego.Controller
}

type WheatherJSON struct {
	Base    string
	Clouds  map[string]interface{}
	Cod     int
	Coord   map[string]interface{}
	Dt      int
	Id      int
	Main    map[string]interface{}
	Name    string
	Sys     map[string]interface{}
	Weather []map[string]interface{} // map[string]interface{}
	Wind    map[string]interface{}
}

// @Title FillingResponse

func FillingResponse(source WheatherJSON) (resp map[string]interface{}) {
	resp = make(map[string]interface{})
	resp["location"] = fmt.Sprintf("%v, %v", source.Name, source.Sys["country"])
	resp["temperature"] = fmt.Sprintf("%v", source.Main["temp"])
	resp["wind"] = fmt.Sprintf("%v m/s", source.Wind["speed"])
	resp["cloudines"] = source.Weather[0]["description"]
	resp["presure"] = fmt.Sprintf("%v hpa", source.Main["pressure"])
	resp["humidity"] = fmt.Sprintf("%v%%", source.Main["humidity"])
	resp["sunrise"] = source.Sys["sunrise"]
	resp["sunset"] = source.Sys["sunset"]
	resp["geo_coordinates"] = fmt.Sprintf("[%v, %v]", source.Coord["lat"], source.Coord["lon"])
	resp["requested_time"] = fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))
	return
}

/*
 */

// @Title Get
func (controller WeatherController) Get() {
	log.Print("Handle for Get WeatherController Requests")
	// Trying to retrieve the params from URL
	city := controller.GetString("city")
	country := controller.GetString("country")
	if city == "" || country == "" {
		// error, incomplete params
		log.Print("error, incomplete params, is needed city and country")
		controller.Abort("401")
	}
	// If params was sended, continue with requests
	log.Print("city: ", city)
	log.Print("country: ", country)

	// Example URL to get Weather
	// "http://api.openweathermap.org/data/2.5/weather?q=Bogota,co&appid=8a14e8c7b941473ca2bc48b9e055e5ba"
	querystring := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s", city, country, APIKEY)
	resp, err := http.Get(querystring)
	if err != nil {
		// handle error
		log.Print("Error getting Weather from api.openweathermap.org")
		controller.Abort("401")
	}
	if resp.StatusCode != 200 {
		fmt.Println("Code Status invalid: ", resp.StatusCode)
		controller.Abort("401")
	}
	// Closing conection
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading result from api.openweathermap.org")
	}

	// Building response
	// var raw map[string]interface{}
	wjson := WheatherJSON{}
	log.Print("Before saving data")
	log.Print(wjson)

	errUn := json.Unmarshal(body, &wjson)
	if errUn != nil {
		fmt.Println("Error in unmarshall: ", errUn)
		controller.Abort("401")
	}
	// out, _ := json.Marshal(raw)
	// out2, _ := json.Marshal(wjson)
	log.Print("After saving data: ")
	log.Print(wjson)
	log.Print(FillingResponse(wjson))
	response := FillingResponse(wjson)

	// controller.Data["json"] = raw
	controller.Data["json"] = response
	controller.ServeJSON()
}
