package restlayer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restProject/serverlayer/dbtools"
	"restProject/serverlayer/model"
	"strconv"

	"github.com/gorilla/mux"
)

func SelectAllUsers(response http.ResponseWriter, request *http.Request) {
	students := dbtools.SelectAllUsers()
	json.NewEncoder(response).Encode(students)
}

func SelectUserByName(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name, ok := vars["name"]
	if !ok {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Error: %v", ok)
	}
	user := dbtools.SelectUserByName(name)
	json.NewEncoder(response).Encode(user)
}

func SelectUserByAge(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	age, ok := vars["age"]
	if !ok {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Error:%v", ok)
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		fmt.Fprintf(response, "Not converted string to int: %v", err)
	}
	user := dbtools.SelectUserByAge(int32(ageInt))
	json.NewEncoder(response).Encode(user)
}

func SaveUser(response http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "Server error:%v", err)
	}
	dbtools.AddUser(user)
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(response, "Error: %v", err)
	}
	dbtools.UpdateUser(user)
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(response, "Err: %v", err)
	}
	dbtools.DeleteUser(user)
}
