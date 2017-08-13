package apiHandler

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)


func HandleId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	log.Println(vars["id"])
	w.Write([]byte("This sub router works only on /api"))
}