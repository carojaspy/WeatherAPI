package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateRequestWeatherTable_20181219_123303 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateRequestWeatherTable_20181219_123303{}
	m.Created = "20181219_123303"

	migration.Register("CreateRequestWeatherTable_20181219_123303", m)
}

// Run the migrations
func (m *CreateRequestWeatherTable_20181219_123303) Up() {
	m.SQL("CREATE TABLE IF NOT EXISTS `request_weather` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY, `country` varchar(255) NOT NULL DEFAULT '' , `city` varchar(255) NOT NULL DEFAULT '' , `requested_time` datetime NOT NULL) ENGINE=InnoDB;")
}

// Reverse the migrations
func (m *CreateRequestWeatherTable_20181219_123303) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
}
