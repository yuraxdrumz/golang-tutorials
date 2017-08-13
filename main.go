package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"apiHandler"
	"gopkg.in/mgo.v2"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("%s %s not found\n", r.Method, r.URL)))
}

func main(){
	session, err := mgo.Dial("localhost/golang-test")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("store").C("books")
	book := map[string]string{
		"book1":"blabla",
	}
	err = c.Insert(book)
	log.Println(err)
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	apiHandler.RegisterRouter("/api", router)
	log.Printf("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}