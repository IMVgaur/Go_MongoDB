package dao

import (
	"testing"

	"github.com/Go_Mongo/TestApp/model"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestConnect(t *testing.T) {
	Connect()
}

func TestAddEmployee(t *testing.T) {
	Init()
	emp := model.Employee{
		ID:   "1999",
		Name: "GAUR",
	}
	err := AddEmployee(emp)
	if err != nil {
		t.Errorf("Expected no error but get error : %v", err)
	}
}

func TestGetEmployee(t *testing.T) {
	Init()
	id := "1954"
	_, err := GetByID(id)
	if err != nil {
		t.Errorf("Expected no error but got error : %v", err)
	}
}

func TestGetEmployeeN(t *testing.T) {
	Init()
	id := "1954111111"
	_, err := GetByID(id)
	if err == nil {
		t.Errorf("Expected no error but got error : %v", err)
	}
}

func TestGetAll(t *testing.T) {
	Init()
	_, err := GetAll()
	if err != nil {
		t.Errorf("Expected no error but got error : %v", err)
	}
}

func TestUpdateEmp(t *testing.T) {
	Init()
	id := "1999"
	err := UpdateEmp(id)
	if err != nil {
		t.Errorf("Expected no error but got error : %v", err)
	}
}

func TestUpdateEmpN(t *testing.T) {
	Init()
	id := "1999999999999999999"
	err := UpdateEmp(id)
	if err == nil {
		t.Errorf("Expected error but got no error : %v", err)
	}
}
func TestRemoveEmp(t *testing.T) {
	Init()
	id := "1999"
	err := RemoveEmp(id)
	if err != nil {
		t.Errorf("Expected no error but got error : %v", err)
	}
}

func TestRemoveEmpN(t *testing.T) {
	Init()
	id := "195999999"
	err := RemoveEmp(id)
	if err == nil {
		t.Errorf("Expected error but got no error : %v", err)
	}
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
