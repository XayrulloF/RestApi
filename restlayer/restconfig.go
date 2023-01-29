package restlayer

import (
	"net/http"

	"github.com/gorilla/mux"
)

func restConfig(router *mux.Router) {
	restRouter := router.PathPrefix("/restapi").Subrouter()

	//localhost:/restapi/students
	restRouter.Methods("GET").Path("/students").HandlerFunc(SelectAllUsers)

	//localhost:/restapi/student/{name}
	restRouter.Methods("GET").Path("/student/{name}").HandlerFunc(SelectUserByName)

	//localhost:/restapi/student/{age}
	restRouter.Methods("GET").Path("/student/{age:[0-9]+}").HandlerFunc(SelectUserByAge)

	//localhost:/restapi/student/add
	restRouter.Methods("POST").Path("/student/add").HandlerFunc(SaveUser)

	//localhost:/restapi/student/update
	restRouter.Methods("POST").Path("/student/update").HandlerFunc(UpdateUser)

	//localhost:/restapi/student/delete
	restRouter.Methods("POST").Path("/student/delete").HandlerFunc(DeleteUser)
}

func RestStart(endpoint string) error {
	router := mux.NewRouter()
	restConfig(router)
	return http.ListenAndServe(endpoint, router)
}
