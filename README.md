## WeatherAPI Proyect from go-bootcamp

Weather proyect using beego framekwork

#### Timing

The measure of time is in __time.info__ file.


## Change Weather provider

The Weather provider is defined by weatherprovider in app.conf, there's two providers defined:

To get Weather info from http://api.openweathermap.org/data/2.5/, must be

    weatherprovider = APIProvider

To get Weather from JSON files:

    weatherprovider = FileProvider


