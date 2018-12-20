package test

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/carojaspy/WeatherAPI/routers"
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
}

// TestTaskWeather Test for Create Tasks
func TestTaskWeather(t *testing.T) {

	r, _ := http.NewRequest("GET", "/v1/scheduler/weather", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestTaskWeather")
	log.Println(w.Code)

	Convey("Given a HTTP request for /v1/scheduler Without params", t, func() {
		req, _ := http.NewRequest("GET", "/v1/scheduler/weather", nil)
		resp := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(resp, req)

		Convey("The response should be a 404", func() {
			So(resp.Code, ShouldEqual, 404)
		})
	})

	Convey("Given a PUT request for /v1/scheduler with params ", t, func() {
		// Open the file.
		successCities, _ := os.Open(FILEPATH + SUCCESS_CITIES)
		// Create a new Scanner for the file.
		scanner := bufio.NewScanner(successCities)
		// Loop over all lines in the file and print them.
		var jsonStr []byte
		for scanner.Scan() {
			line := scanner.Text()
			res := strings.Split(line, ",")
			// log.Println(line, res, res[0], res[1])
			jsonStr = []byte(fmt.Sprintf(`{"city":"%s", "country":"%s"}`, res[0], res[1]))

			// Adding payload to Put request
			req, _ := http.NewRequest("PUT", "/v1/scheduler/weather", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(resp, req)
			log.Println(resp.Body)

			Convey(fmt.Sprintf("%s-%s Response should be a 202", res[0], res[1]), func() {
				So(resp.Code, ShouldEqual, 202)
			})
		}
	})

	Convey("Given a PUT request for /v1/scheduler with params", t, func() {
		// Open the file.
		faillCities, _ := os.Open(FILEPATH + FAIL_CITIES)
		// Create a new Scanner for the file.
		scanner := bufio.NewScanner(faillCities)
		// Loop over all lines in the file and print them.
		var jsonStr []byte
		for scanner.Scan() {
			line := scanner.Text()
			res := strings.Split(line, ",")
			// log.Println(line, res, res[0], res[1])
			jsonStr = []byte(fmt.Sprintf(`{"city":"%s", "country":"%s"}`, res[0], res[1]))

			// Adding payload to Put request
			req, _ := http.NewRequest("PUT", "/v1/scheduler/weather", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(resp, req)

			Convey(fmt.Sprintf("%s-%s Response should be a 404", res[0], res[1]), func() {
				So(resp.Code, ShouldEqual, 404)
			})
		}
	})

}
