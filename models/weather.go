package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

// WheatherJSON .
type WheatherJSON struct {
	Base    string                   `json:"base,omitempty"`
	Clouds  map[string]interface{}   `json:"clouds,omitempty"`
	Cod     int                      `json:"cod,omitempty"`
	Coord   map[string]interface{}   `json:"coord,omitempty"`
	Dt      int                      `json:"dt,omitempty"`
	ID      int                      `json:"id,omitempty"`
	Main    map[string]interface{}   `json:"main,omitempty"`
	Name    string                   `json:"name,omitempty"`
	Sys     map[string]interface{}   `json:"sys,omitempty"`
	Weather []map[string]interface{} `json:"weather,omitempty"`
	Wind    map[string]interface{}   `json:"wind,omitempty"`
}

// Weather .
type Weather struct {
	Id            int
	Location      string
	Temperature   string
	Wind          string
	Cloudines     string
	Presure       string
	Humidity      string
	Sunrise       string
	Sunset string 
	GeoCoordinates string
	RequestedTime time.Time
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Weather))
}

// FillingResponse ..
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

// FillingDBModel ..
func FillingDBModel(source WheatherJSON) (resp Weather) {
	// resp = make(map[string]interface{})
	resp = Weather{}
	resp.Location = fmt.Sprintf("%v, %v", source.Name, source.Sys["country"])
	resp.Temperature = fmt.Sprintf("%v", source.Main["temp"])
	resp.Wind = fmt.Sprintf("%v m/s", source.Wind["speed"])
	resp.Cloudines = fmt.Sprintf("%v", source.Weather[0]["description"])
	resp.Presure = fmt.Sprintf("%v hpa", source.Main["pressure"])
	resp.Humidity = fmt.Sprintf("%v%%", source.Main["humidity"])
	resp.Sunrise = fmt.Sprintf("%v", source.Sys["sunrise"])
	resp.Sunset = fmt.Sprintf("%v", source.Sys["sunset"])
	resp.GeoCoordinates = fmt.Sprintf("[%v, %v]", source.Coord["lat"], source.Coord["lon"])
	resp.RequestedTime = time.Now() // .Format("2006-01-02 15:04:05")
	return
}

