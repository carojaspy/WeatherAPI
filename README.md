## WeatherAPI Proyect from go-bootcamp

Weather proyect using beego framekwork

#### Timing

The measure of time is in __time.info__ file.


bee run -gendoc=true
bee generate controller Weather



docker build -t weather-api-image .
docker-compose build
docker-compose run


dep.exe init
dep.exe ensure -v

##### Enter to container in mode interactive
docker exec -ti weather_api /bin/bash

