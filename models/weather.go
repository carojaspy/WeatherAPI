package models

import (
	"fmt"
	"time"
	"log"
    "errors"
	"github.com/astaxie/beego/orm"
)

// LapseSeconds . 
const LapseSeconds = 300
func getLapse() float64 {
	// Returns the time to compare between 2 dates
	return LapseSeconds
}

// Database .
type Database interface {
	Get(o *orm.Ormer)  error
	GetAll(o *orm.Ormer)  error
	Save(o *orm.Ormer) error
	IsValid(o *orm.Ormer) error
}

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

// Get Fetch a single object from Db
func (w *Weather) Get(o orm.Ormer) error {
	return errors.New("Not implemented")
}
  
// Save Persist the Objet to the Database
func (w *Weather) Save(o orm.Ormer) error {
	// fmt.Println("Inserting row ...")
	id, err := o.Insert(w)
	if err == nil {
		fmt.Printf("Weather Row inserted with ID: %v", id)
		return nil
	}
	return err
}

// IsValid . Check if has elapsed 300 seconds to insert a new row
func (w *Weather) IsValid(o orm.Ormer) error {
	qs := o.QueryTable(*w) // return a QuerySeter

	/*	Check if there's */
	var lastRow Weather

	// Just One row
	err := qs.OrderBy("-Id").Filter("Location", w.Location).One(&lastRow)
	if err != nil {
		// No previous rows inserted to DB
		log.Println(err)
		return nil
	} else {
		log.Println("Sucess getting weather objects from db")
		// elapsedTime := time.Until(lastRow.RequestedTime.UTC())
		elapsedTime := time.Since(lastRow.RequestedTime.UTC())
		// log.Printf("seconds: %v, LAPSE: %v, %v", elapsedTime.Seconds(), getLapse(), elapsedTime.Seconds() < getLapse())
		if elapsedTime.Seconds() > getLapse(){
			log.Println("New row !!")
			return nil
		} else {
			// still not pass enough time to save another row
			return fmt.Errorf("You cant insert yet: seconds: %v", elapsedTime.Seconds())
		}	//End time elapsed comparission
	}// End QueryFilter if
}

// GetAll Implementing GetAll method from Database to get All rows
func (w *Weather) GetAll(o orm.Ormer)  ([]*Weather, error){
	log.Println("GetAll method")
	var weathers []* Weather
	_, err := o.QueryTable(new(Weather)).All(&weathers)
	if err != nil {
		log.Println(err)
	}
	return weathers, err
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

