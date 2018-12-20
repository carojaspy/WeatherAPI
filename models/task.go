package models

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/astaxie/beego/orm"
)

// Database .
type DatabaseTask interface {
	Get(o orm.Ormer) error
	GetAll(o orm.Ormer) error
	Save(o orm.Ormer) error
	IsValid(o orm.Ormer) (bool, error)
}

// Task .
type Task struct {
	Id       int
	IsActive bool
	City     string
	Country  string
}

/*	Implementing Interface Methods from Database Interface	*/

// Get Fetch a single object from Db
func (t *Task) Get(o orm.Ormer) error {
	return errors.New("Not implemented")
}

// GetAll Implementing GetAll method from Database to get All rows
func (t *Task) GetAll(o orm.Ormer) ([]*Task, error) {
	log.Println("GetAll method")
	var tasks []*Task
	_, err := o.QueryTable(new(Task)).All(&tasks)
	if err != nil {
		log.Println(err)
	}
	return tasks, err
}

// Save Persist the Objet to the Database
func (t *Task) Save(o orm.Ormer) error {
	log.Println("Inserting row ...")
	id, err := o.Insert(t)
	if err == nil {
		log.Printf("Task Row inserted with ID: %v", id)
		// 	Send this task to the periodic task
		return nil
	}
	return err
} //EndMethod

// IsValid . Check if has elapsed 300 seconds to insert a new row
func (t *Task) IsValid(o orm.Ormer) error {
	qs := o.QueryTable(*t) // return a QuerySeter
	log.Println("City and country: ")
	log.Println(t.City, t.Country)

	/*	Check if there's	*/
	var lastRow Task
	// Just One row
	err := qs.OrderBy("-Id").Filter("City", t.City).Filter("Country", t.Country).Filter("IsActive", true).One(&lastRow)
	if err != nil {
		// No previous rows inserted to DB
		if strings.Contains(err.Error(), "no row found") {
			log.Println("No Row for this tasks: ", err.Error())
			return nil
		} else {
			// Another Case that is not a 'no row found' could be a real error
			log.Println(err)
			return err
		}
	} else {
		return fmt.Errorf("Already Exists a Task to that Country : %v", lastRow)
	}
	//return errors.New("Not implemented")
} //End IsValid

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Task))
}
