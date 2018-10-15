package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Go_Mongo/TestApp/dao"

	"github.com/Go_Mongo/TestApp/model"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestAddEmployee(t *testing.T) {
	testSuits := []struct {
		method string
		url    string
		status int
		emp    model.Employee
	}{
		{
			method: "POST",
			url:    "/emp",
			status: 200,
			emp: model.Employee{
				ID:   "1950",
				Name: "test1_name1",
			},
		},
		{
			method: "GET",
			url:    "/emp",
			status: 405,
			emp: model.Employee{
				ID:   "19555",
				Name: "GAur",
			},
		},
	}
	for i, _ := range testSuits {
		dao.Init()
		body := &bytes.Buffer{}
		srv := httptest.NewServer(Handlers())
		defer srv.Close()
		json.NewEncoder(body).Encode(testSuits[i].emp)
		req, err := http.NewRequest(testSuits[i].method, srv.URL+testSuits[i].url, body)
		if err != nil {
			t.Errorf("New request could not be created : %v", err)
		}
		res, _ := http.DefaultClient.Do(req)
		if res.StatusCode != testSuits[i].status {
			t.Errorf("Expected status ok but got different status %v", res.Status)
		}
	}
}

func TestAddEmployeeN(t *testing.T) {
	testSuits := []struct {
		method string
		url    string
		status int
	}{
		{
			method: "POST",
			url:    "/emp",
			status: 500,
		},
	}
	for i, _ := range testSuits {
		dao.Init()
		body := &bytes.Buffer{}
		srv := httptest.NewServer(Handlers())
		defer srv.Close()
		req, err := http.NewRequest(testSuits[i].method, srv.URL+testSuits[i].url, body)
		if err != nil {
			t.Errorf("New request could not be created : %v", err)
		}
		res, _ := http.DefaultClient.Do(req)
		if res.StatusCode != testSuits[i].status {
			t.Errorf("Expected status ok but got different status %v", res.Status)
		}
	}
}

func TestGetEmployee(t *testing.T) {
	testSuits := []struct {
		method string
		url    string
		status int
	}{
		{
			method: "GET",
			url:    "/emp/1958",
			status: 200,
		},
		{
			method: "POST",
			url:    "/emp/1958",
			status: 405,
		},
		{
			method: "GET",
			url:    "/emp/100000",
			status: 500,
		},
	}
	for i, _ := range testSuits {
		dao.Init()
		srv := httptest.NewServer(Handlers())
		defer srv.Close()
		req, err := http.NewRequest(testSuits[i].method, srv.URL+testSuits[i].url, nil)
		if err != nil {
			t.Errorf("New request could not be created : %v", err)
		}
		res, _ := http.DefaultClient.Do(req)
		if res.StatusCode != testSuits[i].status {
			t.Errorf("Expected status ok but got different status %v", res.Status)
		}
	}
}

func TestGetEmployees(t *testing.T) {
	testSuits := []struct {
		method string
		url    string
		status int
	}{
		{
			method: "GET",
			url:    "/emps",
			status: 200,
		},
		{
			method: "POST",
			url:    "/emps",
			status: 405,
		},
	}

	for i, _ := range testSuits {
		dao.Init()
		srv := httptest.NewServer(Handlers())
		defer srv.Close()
		req, err := http.NewRequest(testSuits[i].method, srv.URL+testSuits[i].url, nil)
		if err != nil {
			t.Errorf("New request could not be created : %v", err)
		}
		res, _ := http.DefaultClient.Do(req)
		if res.StatusCode != testSuits[i].status {
			t.Errorf("Expected status ok but got different status %v", res.Status)
		}
	}
}

func TestRemoveEmp(t *testing.T) {
	testSuits := []struct {
		method string
		url    string
		status int
	}{
		{
			method: "DELETE",
			url:    "/emp/delete/1950",
			status: 200,
		},
		{
			method: "PUT",
			url:    "/emp/delete/1956",
			status: 405,
		},
		{
			method: "DELETE",
			url:    "/emp/delete/19560000",
			status: 500,
		},
	}
	for i, _ := range testSuits {
		dao.Init()
		srv := httptest.NewServer(Handlers())
		defer srv.Close()
		req, err := http.NewRequest(testSuits[i].method, srv.URL+testSuits[i].url, nil)
		if err != nil {
			t.Errorf("New request could not be created : %v", err)
		}
		res, _ := http.DefaultClient.Do(req)
		if res.StatusCode != testSuits[i].status {
			t.Errorf("Expected status ok but got different status %v", res.Status)
		}
	}
}

func TestUpdateEmp(t *testing.T) {
	testSuits := []struct {
		method string
		url    string
		status int
	}{
		{
			method: "UPDATE",
			url:    "/emp/update/1950",
			status: 200,
		},
		{
			method: "DELETE",
			url:    "/emp/update/1958",
			status: 405,
		},
		{
			method: "UPDATE",
			url:    "/emp/update/195800000",
			status: 500,
		},
	}
	for i, _ := range testSuits {
		dao.Init()
		srv := httptest.NewServer(Handlers())
		defer srv.Close()
		req, err := http.NewRequest(testSuits[i].method, srv.URL+testSuits[i].url, nil)
		if err != nil {
			t.Errorf("New request could not be created : %v", err)
		}
		res, _ := http.DefaultClient.Do(req)
		if res.StatusCode != testSuits[i].status {
			t.Errorf("Expected status ok but got different status %v", res.Status)
		}
	}
}

func TestHandler(t *testing.T) {
	testSuits := []struct {
		method   string
		testName string
		url      string
		status   int
	}{
		{
			method:   "GET",
			testName: "Test1",
			url:      "/emps",
			status:   200,
		},
		{
			method:   "GET",
			testName: "Test2",
			url:      "/emp/1958",
			status:   200,
		},
		{
			method:   "DELETE",
			testName: "Test3",
			url:      "/emp/delete/1950",
			status:   200,
		},
		{
			method:   "UPDATE",
			testName: "Test4",
			url:      "/emp/update/1958",
			status:   200,
		},
		{
			method:   "POST",
			testName: "Test5",
			url:      "/emp",
			status:   500,
		},
	}
	for i, _ := range testSuits {
		dao.Init()
		srv := httptest.NewServer(Handlers())
		defer srv.Close()
		req, err := http.NewRequest(testSuits[i].method, srv.URL+testSuits[i].url, nil)
		if err != nil {
			t.Errorf("New request could not be created : %v", err)
		}
		res, _ := http.DefaultClient.Do(req)
		if res.StatusCode != testSuits[i].status {
			t.Errorf("Expected status ok but got different status %v", res.Status)
		}
	}
}

func TestRespondWithError(t *testing.T) {
	rec := httptest.NewRecorder()
	respondWithError(rec, http.StatusInternalServerError, "Testing function")
}

func TestRespondWithJson(t *testing.T) {
	rec := httptest.NewRecorder()
	respondWithError(rec, http.StatusInternalServerError, "Testing function")
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
