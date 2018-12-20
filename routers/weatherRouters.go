package routers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	fmt.Println("Init WeatherRouters")
	beego.GlobalControllerRouter["WeatherAPI/controllers:WeatherV1Controller"] = append(beego.GlobalControllerRouter["github.com/carojaspy/WeatherAPI/controllers:WeatherV1Controller"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["WeatherAPI/controllers:WeatherV2Controller"] = append(beego.GlobalControllerRouter["github.com/carojaspy/WeatherAPI/controllers:WeatherV2Controller"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["WeatherAPI/controllers:WeatherV2Controller"] = append(beego.GlobalControllerRouter["github.com/carojaspy/WeatherAPI/controllers:WeatherV2Controller"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}
