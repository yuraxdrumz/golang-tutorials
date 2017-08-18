package apiHandler

import (
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"github.com/mssola/user_agent"
	"gopkg.in/mgo.v2/bson"
	"errors"
)


const (
	DBNAME = "newonetapadmin"
	COLLECTIONNAME = "tags"
	TAGID = "tagId"
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
	return func(w http.ResponseWriter, r *http.Request) {

		queries := r.URL.Query()
		parsedUserAgent := parseUserAgent(r.UserAgent())

		session := s.Copy()
		defer session.Close()

		tag, err := getTag(session, queries.Get(TAGID))
		if err != nil{
			ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}
		matchedGroup := tag.getMatchedGroup(&parsedUserAgent)
		respBody, err := json.MarshalIndent(matchedGroup, "", " ")
		if err != nil {
			ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
		return
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
	isId := bson.IsObjectIdHex(id)
	if isId {
		dataStore := s.DB(DBNAME).C(COLLECTIONNAME)
		err = dataStore.Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&tag)
	}else{
		err = errors.New("Id is not valid")
	}
	return
}

func (tag *Tag) getMatchedGroup(user *User) (group *Group){
	switch len(tag.Groups) {
	default:
		group = &tag.Groups[0]
	}
	return
}

func (tag *Tag) antiBot(user *User) (ab *AntiBot){
	switch tag.Protected {
	case "e":
	case "b":
		// check if bot
	case "n":
		// pass isOk: true to antibot struct
	default:
		// check if bot by default
	}
	// populate antibot and return it
	return
}
