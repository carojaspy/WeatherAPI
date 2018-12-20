FROM golang:1.11.2
ADD . /go/src/github.com/carojaspy/WeatherAPI
WORKDIR /go/src/github.com/carojaspy/WeatherAPI
RUN  go get -u github.com/go-sql-driver/mysql && go get -u github.com/beego/bee # &&  ./setup.sh
CMD ["bee", "run"]

