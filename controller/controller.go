package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Go_Mongo/TestApp/dao"
	"github.com/Go_Mongo/TestApp/model"

	"github.com/gorilla/mux"
)

//Handlers ...
func Handlers() http.Handler {
	server := mux.NewRouter()
	server.HandleFunc("/emp", AddEmployee).Methods("POST")
	server.HandleFunc("/emps", GetEmployees).Methods("GET")
	server.HandleFunc("/emp/update/{id}", UpdateEmp).Methods("UPDATE")
	server.HandleFunc("/emp/delete/{id}", RemoveEmp).Methods("DELETE")
	server.HandleFunc("/emp/{id}", GetEmployee).Methods("GET")
	return server
}

//AddEmployee ...
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respondWithError(w, http.StatusMethodNotAllowed, "Bad Request")
	}
	if r.Body != nil {
		defer r.Body.Close()
		var emp model.Employee
		err := json.NewDecoder(r.Body).Decode(&emp)
		if err == nil {
			dao.AddEmployee(emp)
			respondWithJSON(w, http.StatusOK, "Added successfully")
		}
		respondWithError(w, http.StatusInternalServerError, "Bad Request")
	}
	respondWithError(w, http.StatusBadRequest, "Bad Request")
}

//GetEmployee ...
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respondWithError(w, http.StatusMethodNotAllowed, "Bad Request")
	}
	params := mux.Vars(r)
	emp, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, emp)
}

//GetEmployees ...
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respondWithError(w, http.StatusMethodNotAllowed, "Bad Request")
	}
	emps, err := dao.GetAll()
	respondWithJSON(w, http.StatusOK, emps, err)
}

//UpdateEmp ...
func UpdateEmp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "UPDATE" {
		respondWithError(w, http.StatusMethodNotAllowed, "Bad request")
	}
	params := mux.Vars(r)
	err := dao.UpdateEmp(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, "Updation successfull...")
}

//RemoveEmp ...
func RemoveEmp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		respondWithError(w, http.StatusMethodNotAllowed, "Bad request")
	}
	params := mux.Vars(r)
	err := dao.RemoveEmp(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, "Removal successfull...")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"Error ": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, message ...interface{}) {
	response, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
