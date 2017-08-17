package apiHandler

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func RegisterRouter(path string, r *mux.Router, session *mgo.Session){
	customRouter := r.PathPrefix(path).Subrouter()
	customRouter.HandleFunc("/getHtmlCode", GetHtmlCode(session)).Methods("GET")
}