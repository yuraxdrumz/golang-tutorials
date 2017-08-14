package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"apiHandler"
	"gopkg.in/mgo.v2"
)

const (
	MONGOADDR = "localhost"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("%s %s not found\n", r.Method, r.URL)))
}

func main(){
	s, err := mgo.Dial(MONGOADDR)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to db: " + MONGOADDR)
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	apiHandler.RegisterRouter("/api", router, s)
	log.Printf("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}