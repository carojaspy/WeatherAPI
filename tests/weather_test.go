package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/carojaspy/WeatherAPI/routers"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	log.Print(file)
	// apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	//	Linux Path
	// apppath := "/mnt/c/Users/carlosalberto.rojas/Desktop/go/src/github.com/carojaspy/WeatherAPI/"

	//	Docker path
	apppath := "/go/src/github.com/carojaspy/WeatherAPI/"
	// apppath := "\\mnt\\c\\Users\\carlosalberto.rojas\\Desktop\\go\\src\\WeatherAPI\\"

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

	r, _ := http.NewRequest("GET", "/v1/weather?city=Peru&country=us", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGet")

	Convey("Subject: Test Get Weather Request with City and Country params\n", t, func() {
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

}
