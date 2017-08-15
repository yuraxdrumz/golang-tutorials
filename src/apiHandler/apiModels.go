package apiHandler

import "gopkg.in/mgo.v2/bson"

type Book struct {
	Book1    string   `json:"bookName"`
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
