package model

import "gopkg.in/mgo.v2/bson"

type Info struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Email         string        `json:"email" bson:"email"`
	IdAlgumaCoisa string        `json:"idAlgumaCoisa" bson:"idAlgumaCoisa"`
}
