package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/carojaspy/WeatherAPI/controllers:WeatherController"] = append(beego.GlobalControllerRouter["github.com/carojaspy/WeatherAPI/controllers:WeatherController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carojaspy/WeatherAPI/controllers:WeatherController"] = append(beego.GlobalControllerRouter["github.com/carojaspy/WeatherAPI/controllers:WeatherController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
