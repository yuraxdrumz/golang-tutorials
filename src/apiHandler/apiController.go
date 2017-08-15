package apiHandler

import (
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)




func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func GetBooks(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		session := s.Copy()
		defer session.Close()

		c := session.DB("store").C("books")

		var books []Book
		err := c.Find(bson.M{}).All(&books)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all books: ", err)
			return
		}
		respBody, err := json.MarshalIndent(books, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}

}

func InsertBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		session := s.Copy()
		defer session.Close()

		c := session.DB("store").C("books")

		var book Book
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			ErrorWithJSON(w, "Read Body error", http.StatusInternalServerError)
			log.Println("Failed get all books: ", err)
			return
		}
		jsonErr := json.Unmarshal(body, &book)
		if jsonErr != nil {
			ErrorWithJSON(w, "Json parse error", http.StatusInternalServerError)
			log.Println("Failed get all books: ", err)
			return
		}
		dbErr := c.Insert(book)
		if dbErr != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all books: ", err)
			return
		}
		respBody, err := json.MarshalIndent(book.Book1 + " was inserted", "", " ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}

}

func GetBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		session := s.Copy()
		defer session.Close()

		c := session.DB("store").C("books")

		var book Book
		err := c.Find(bson.M{"book1":vars["bookName"]}).One(&book)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all books:", err)
			return
		}
		respBody, err := json.MarshalIndent(book, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}

}