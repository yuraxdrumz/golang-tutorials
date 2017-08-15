package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"apiHandler"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"fmt"
)

const (
	MONGOADDR = "localhost"
)


func NotFoundHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	err := apiHandler.ErrorMessage{Err:fmt.Sprintf("%s %s not found", r.Method, r.URL)}
	answer,_ := json.Marshal(err)
	w.Write(answer)
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