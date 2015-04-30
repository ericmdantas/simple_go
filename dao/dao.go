package dao

import (
	model "github.com/ericmdantas/simple_go/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL        string = "localhost"
	DB         string = "simple_go"
	Collection string = "info"
)

func Cria(info model.Info) {
	session, err := mgo.Dial(URL)

	if err != nil {
		panic(err)
	}

	defer session.Close()

	info.Id = bson.NewObjectId()

	coll := session.DB(DB).C(Collection)

	err = coll.Insert(info)

	if err != nil {
		panic(err)
	}
}
