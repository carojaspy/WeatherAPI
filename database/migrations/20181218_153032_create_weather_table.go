package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateWeatherTable_20181218_153032 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateWeatherTable_20181218_153032{}
	m.Created = "20181218_153032"

	migration.Register("CreateWeatherTable_20181218_153032", m)
}

// Run the migrations
func (m *CreateWeatherTable_20181218_153032) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS `weather` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,`location` varchar(255) NOT NULL DEFAULT '' ,`temperature` varchar(255) NOT NULL DEFAULT '' ,`wind` varchar(255) NOT NULL DEFAULT '' ,`cloudines` varchar(255) NOT NULL DEFAULT '' ,`presure` varchar(255) NOT NULL DEFAULT '' ,`humidity` varchar(255) NOT NULL DEFAULT '' ,`sunrise` varchar(255) NOT NULL DEFAULT '' ,`sunset` varchar(255) NOT NULL DEFAULT '' ,`geo_coordinates` varchar(255) NOT NULL DEFAULT '' ,`requested_time` datetime NOT NULL) ENGINE=InnoDB;")

}

// Reverse the migrations
func (m *CreateWeatherTable_20181218_153032) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
