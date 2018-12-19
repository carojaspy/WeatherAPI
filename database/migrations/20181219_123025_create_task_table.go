package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTaskTable_20181219_123025 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTaskTable_20181219_123025{}
	m.Created = "20181219_123025"

	migration.Register("CreateTaskTable_20181219_123025", m)
}

// Run the migrations
func (m *CreateTaskTable_20181219_123025) Up() {
	m.SQL("CREATE TABLE IF NOT EXISTS `task` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY, `is_active` bool NOT NULL DEFAULT FALSE , `city` varchar(255) NOT NULL DEFAULT '' , `country` varchar(255) NOT NULL DEFAULT '') ENGINE=InnoDB;")
}

// Reverse the migrations
func (m *CreateTaskTable_20181219_123025) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
