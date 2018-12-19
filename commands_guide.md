
## Guide of commands 


### Bee (tool) commands

##### Create Swagger DOC

    bee run -gendoc=true

##### Create a Controller from bee tool

    bee generate controller Weather



### Docker Commands

##### Build the Proyect using Dockerfile

    docker build -t weather-api-image .

##### Build the Proyect using docker-compose.yml    
    docker-compose build

##### Run the proyect using docker-compose

    docker-compose run 
    docker-compose run <container_name>

##### Create Gopkg.lock, vendor/ and Gopkg.toml

    dep.exe init

##### After pull the proyect, check if all requirements are satisfied
dep.exe ensure -v

##### Enter to container in mode interactive
docker exec -ti weather_api /bin/bash

