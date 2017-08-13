package apiHandler

import (
	"github.com/gorilla/mux"
)

func RegisterRouter(path string, r *mux.Router){
	customRouter := r.PathPrefix(path).Subrouter()
	customRouter.HandleFunc("/test/{id}", HandleId).Methods("GET")
}