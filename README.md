# WeatherAPI Proyect from go-bootcamp

Weather proyect using beego framekwork

## How to Run the Proyect

1.- This Proyect should be put it in the path 

    $GOPATH/github.com/carojaspy/$THISREPO

for example, if the $GOPATH is in /opt/go/ : 

    /opt/go/github.com/carojaspy/WeatherAPI

2.- Run docker containers:

Move into the project folder

    cd $GOPATH/carojaspy/WeatherAPI/

Build 
    
    docker-compose build
    
run    

    docker-compose up

Ensure that there's no any service or another MySQL instance running on that port, or modify the ports listen to on **docker-compose.yml** file

    ports:
      - 3306:3306

3.- Create Migrations


#### Timing

The measure of time is in __time.info__ file.


## Change Weather provider

The Weather provider is defined by weatherprovider in app.conf, there's two providers defined:

To get Weather info from http://api.openweathermap.org/data/2.5/, must be

    weatherprovider = APIProvider

To get Weather from JSON files:

    weatherprovider = FileProvider


## Preloaded Data

In the path /tests/data/ there's two .sql scripts

- **db_backup.sql** : Contains all the database (Create tables, and insert data)

- **db_data.sql** : Contains just the rows (Periodic Tasks, Weather info, Requests)


