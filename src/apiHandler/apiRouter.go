package apiHandler

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func RegisterRouter(path string, r *mux.Router, session *mgo.Session){
	customRouter := r.PathPrefix(path).Subrouter()
	customRouter.HandleFunc("/getBooks", GetBooks(session)).Methods("GET")
	customRouter.HandleFunc("/insertBook", InsertBook(session)).Methods("POST")
	customRouter.HandleFunc("/getBook/{bookName}", GetBook(session)).Methods("GET")
}