package models


import (
	"fmt"
	"log"
    "errors"
	"github.com/astaxie/beego/orm"
)


// Database .
type DatabaseTask interface {
	Get(o *orm.Ormer)  error
	GetAll(o *orm.Ormer)  error
	Save(o *orm.Ormer) error
	IsValid(o *orm.Ormer) (bool, error)
}

// Task . 
type Task struct {
	Id	int
	IsActive bool
	City	string
	Country	string
}

 /*	Implementing Interface Methods from Database Interface	*/


// Get Fetch a single object from Db
func (t *Task) Get(o orm.Ormer) error {
	return errors.New("Not implemented")
}
  
// Save Persist the Objet to the Database
func (t *Task) Save(o orm.Ormer) error {
	log.Println("Inserting row ...")
	id, err := o.Insert(t)
	if err == nil {
		fmt.Printf("Task Row inserted with ID: %v", id)
		return nil
	}
	return err
}//EndMethod

// IsValid . Check if has elapsed 300 seconds to insert a new row
func (t *Task) IsValid(o orm.Ormer) error {
	// qs := o.QueryTable(t) // return a QuerySeter
	// qs.Filter("Location", w)	
	return errors.New("Not implemented")

}//End IsValid


func init() {
	// Need to register model in init
	orm.RegisterModel(new(Task))
}



