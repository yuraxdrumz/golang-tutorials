package apiHandler

import (
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"github.com/mssola/user_agent"
	"gopkg.in/mgo.v2/bson"
)




func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errMessage := ErrorMessage{message}
	answer,jsonErr := json.Marshal(errMessage)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}
	w.Write(answer)
}

func ResponseWithJSON(w http.ResponseWriter, message []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(message)
}


func GetHtmlCode(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){

		queries := r.URL.Query()
		parsedUserAgent := parseUserAgent(r.UserAgent())

		session := s.Copy()
		defer session.Close()

		tag, err := getTag(session, queries.Get("tagId"))
		if err != nil{
			ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
		}
		matchedGroup := tag.getMatchedGroup(parsedUserAgent)
		respBody, err := json.MarshalIndent(matchedGroup.HtmlCode, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, 200)
		//dataStore := session.DB("store").C("tags")
	}
}

func parseUserAgent(userAgent string) User {
	ua := user_agent.New(userAgent)
	engineName, engineVersion := ua.Engine()
	browserName, browserVersion := ua.Browser()
	user := User{
		Bot: ua.Bot(),
		Mobile: ua.Mobile(),
		Platform : ua.Platform(),
		Os: ua.OS(),
		Engine: Engine{
			Name:engineName,
			Version:engineVersion,
		},
		Browser: Browser{
			Name:browserName,
			Version:browserVersion,
		},
	}
	return user
}

func getTag(s *mgo.Session,id string) (tag *Tag, err error){
	dataStore := s.DB("newonetapadmin").C("tags")
	err = dataStore.Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&tag)
	return
}

func (tag *Tag) getMatchedGroup(user User) (group Group){
	switch len(tag.Groups) {
	case 1:
		group = tag.Groups[0]
	default:
		group = tag.Groups[0]
	}
	return
}


//func GetBooks(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
//	return func(w http.ResponseWriter, r *http.Request){
//		session := s.Copy()
//		defer session.Close()
//
//		c := session.DB("store").C("books")
//
//		var books []Book
//		err := c.Find(bson.M{}).All(&books)
//		if err != nil {
//			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
//			log.Println("Failed get all books: ", err)
//			return
//		}
//		respBody, err := json.MarshalIndent(books, "", " ")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		ResponseWithJSON(w, respBody, http.StatusOK)
//	}
//
//}
//func InsertBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
//	return func(w http.ResponseWriter, r *http.Request){
//		session := s.Copy()
//		defer session.Close()
//
//		c := session.DB("store").C("books")
//
//		var book Book
//		body, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			ErrorWithJSON(w, "Read Body error", http.StatusInternalServerError)
//			log.Println("Failed get all books: ", err)
//			return
//		}
//		jsonErr := json.Unmarshal(body, &book)
//		if jsonErr != nil {
//			ErrorWithJSON(w, "Json parse error", http.StatusInternalServerError)
//			log.Println("Failed get all books: ", err)
//			return
//		}
//		dbErr := c.Insert(book)
//		if dbErr != nil {
//			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
//			log.Println("Failed get all books: ", err)
//			return
//		}
//		respBody, err := json.MarshalIndent(book.Book1 + " was inserted", "", " ")
//		if err != nil {
//			log.Fatal(err)
//		}
//		ResponseWithJSON(w, respBody, http.StatusOK)
//	}
//
//}
//func GetBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
//	return func(w http.ResponseWriter, r *http.Request){
//		vars := mux.Vars(r)
//		session := s.Copy()
//		defer session.Close()
//
//		c := session.DB("store").C("books")
//
//		var book Book
//		err := c.Find(bson.M{"book1":vars["bookName"]}).One(&book)
//		if err != nil {
//			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
//			log.Println("Failed get all books:", err)
//			return
//		}
//		respBody, err := json.MarshalIndent(book, "", " ")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		ResponseWithJSON(w, respBody, http.StatusOK)
//	}
//
//}