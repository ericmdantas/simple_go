package dao

import (
	_ "github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func limpaDB() {
	s, _ := mgo.Dial(URL)

	_, _ = s.DB(DB).C(Collection).RemoveAll(bson.M{})

	s.Close()
}

func TestMain(m *testing.M) {
	limpaDB()

	m.Run()
}

var infoTests = []struct {
	info map[string]interface{}
}{
	{
		info: map[string]interface{}{"a": true},
	},
}

func Test_Cria(t *testing.T) {
	for _, _t := range infoTests {
		Cria(_t.info)
	}
}
