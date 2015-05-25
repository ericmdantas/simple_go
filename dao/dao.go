package dao

import (
	_ "github.com/ericmdantas/simple_go/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL        string = "localhost"
	DB         string = "simple_go"
	Collection string = "info"
)

func Cria(info map[string]interface{}) {
	session, err := mgo.Dial(URL)

	session.SetSafe(nil) // fire & forget

	if err != nil {
		session.Close()
		panic(err)
	}

	info["_id"] = bson.NewObjectId()
	session.DB(DB).C(Collection).Insert(info)
	session.Close()
}
