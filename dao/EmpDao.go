package dao

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/Go_Mongo/TestApp/model"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

//Init ...
func Init() {
	Connect()
}

//Connect ...
//Initialization of mongodb
func Connect() {
	session, err := mgo.Dial("localhost:27017")
	if err == nil {
		db = session.DB("test")
	}
}

//AddEmployee ...
func AddEmployee(emp model.Employee) error {
	return db.C("Employee").Insert(emp)
}

//GetByID ...
func GetByID(id string) (model.Employee, error) {
	var emp model.Employee
	err := db.C("Employee").Find(bson.M{"id": id}).One(&emp)
	return emp, err
}

//GetAll ...
func GetAll() ([]model.Employee, error) {
	var emps []model.Employee
	err := db.C("Employee").Find(bson.M{}).All(&emps)
	return emps, err
}

//UpdateEmp ...
func UpdateEmp(id string) error {
	emp, err := GetByID(id)
	if err != nil {
		return err
	}
	return db.C("Employee").Update(bson.M{"id": id}, emp)
}

//RemoveEmp ...
func RemoveEmp(id string) error {
	emp, err := GetByID(id)
	if err == nil {
		return db.C("Employee").Remove(emp)
	}
	return err
}
