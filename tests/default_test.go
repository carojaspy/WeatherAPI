package test

import (
	"encoding/json"
	"io/ioutil"
	_ "github.com/carojaspy/WeatherAPI/routers"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../.."+string(filepath.Separator))))
	log.Print(apppath)
	beego.TestBeegoInit(apppath)
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
	})
}