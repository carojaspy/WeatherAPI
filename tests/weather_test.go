package test

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/carojaspy/WeatherAPI/controllers"
	_ "github.com/carojaspy/WeatherAPI/models"
	_ "github.com/carojaspy/WeatherAPI/routers"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"
)

// Path to JSON files
const FILEPATH = "/go/src/github.com/carojaspy/WeatherAPI/tests/"

// SUCCESS_CITIES .
const SUCCESS_CITIES = "txt/SUCCES_CODES.txt"

// FAIL_CITIES .
const FAIL_CITIES = "txt/FAIL_CODES.txt"

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../.."+string(filepath.Separator))))

	beego.TestBeegoInit(apppath)

	// Connect to Database
	orm.RegisterDataBase("default", "mysql", "root:@/weatherapi?charset=utf8")
	o := orm.NewOrm()
	orm.Debug = true
	o.Using("default") // Using default, you can use other database

}

// TestWeatherProvider Test connection with Weather Provider
func TestWeatherProvider(t *testing.T) {
	Convey("Subject: Test success response (api.openweathermap.org)\n", t, func() {
		righCities, _ := os.Open(FILEPATH + SUCCESS_CITIES)
		scanner := bufio.NewScanner(righCities)
		for scanner.Scan() {
			line := scanner.Text()
			res := strings.Split(line, ",")
			log.Println(res[0], res[1])
			_, err := controllers.GetWeatherFromProvider(res[0], res[1])
			Convey("Subject: err Should be nil", func() {
				So(err, ShouldBeNil)
			})
		}
	})

	Convey("Subject: Test err response (api.openweathermap.org)\n", t, func() {
		failCities, _ := os.Open(FILEPATH + FAIL_CITIES)
		scanner := bufio.NewScanner(failCities)
		for scanner.Scan() {
			line := scanner.Text()
			res := strings.Split(line, ",")
			_, err := controllers.GetWeatherFromProvider(res[0], res[1])
			Convey("err shouldn't be nil", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("err.body should contains 404 or Not found", func() {
				So(err.Error(), ShouldContainSubstring, "404, Not found")
			})

		}
	})

}

// TestEmptyGet Request Without params, should be fail
func TestEmptyGet(t *testing.T) {

	r, _ := http.NewRequest("GET", "/v1/weather", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestEmptyGet")
	log.Println(w.Code)

	Convey("Subject: Test Get Weather Request\n", t, func() {
		Convey("Status Code Should Be 404 for incomplete params", func() {
			So(w.Code, ShouldEqual, 404)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestEmptyGet Request Without params, should be fail
func TestGet(t *testing.T) {

	Convey("Subject: Test Get Weather Request with City and Country params\n", t, func() {
		r, _ := http.NewRequest("GET", "/v1/weather?city=Peru&country=us", nil)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		beego.Trace("testing", "TestGet")

		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})

		Convey("The Result Should be a JSON response", func() {
			body, err := ioutil.ReadAll(w.Body)
			// Building response
			var wjson interface{}
			err = json.Unmarshal(body, &wjson)
			So(err, ShouldBeNil)
		})
	})

	Convey("Subject: Test Get Weather Request with wrong City(Right) and Country(Wrong) params\n", t, func() {
		r, _ := http.NewRequest("GET", "/v1/weather?city=Peru&country=test", nil)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		beego.Trace("testing", "TestGet")

		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})

		Convey("The Result Should be a JSON response", func() {
			body, err := ioutil.ReadAll(w.Body)
			// Building response
			var wjson interface{}
			err = json.Unmarshal(body, &wjson)
			So(err, ShouldBeNil)
		})
	})

	Convey("Subject: Test Get Weather Request with wrong City(Wrong) and Country(Right) params\n", t, func() {
		r, _ := http.NewRequest("GET", "/v1/weather?city=test&country=us", nil)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		beego.Trace("testing", "TestGet")

		Convey("Status Code Should Be 404", func() {
			So(w.Code, ShouldEqual, 404)
		})

		Convey("The Result ShouldContainSubstring 'not found' text response", func() {
			body, _ := ioutil.ReadAll(w.Body)
			So(string(body), ShouldContainSubstring, "Not found")
		})
	})

}
