package apiHandler

import "gopkg.in/mgo.v2/bson"

type Engine struct {
	Name string
	Version string
}

type Browser struct {
	Name string
	Version string
}

type Ct struct {
	Value string `json:"name"`
}

type Country struct {
	Value string `json:"name"`
}

type OperatingSystem struct {
	Value string `json:"name"`
}

type Device struct {
	Value string `json:"name"`
}

type Group struct{
	Name string `json:"name"`
	HtmlCode string `json:"htmlCode" bson:"htmlCode"`
	Cts []Ct `json:"cts"`
	Countries []Country `json:"countries"`
	Oss []OperatingSystem `json:"oss"`
	Devices []Device `json:"devices"`
}

type Tag struct {
	ID bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
	Groups []Group `json:"groups"`
	Owner string `json:"owner"`
	Secure bool `json:"secure"`
	Protected string `json:"protected"`
}

type User struct {
	Mobile bool
	Bot bool
	Platform string
	Os string
	Engine Engine
	Browser Browser
}

type AntiBot struct {
	isOk bool
	isBot bool
}


type ErrorMessage struct {
	Err string `json:"error"`
}

func (e *ErrorMessage) Error() string {
	return e.Err
}