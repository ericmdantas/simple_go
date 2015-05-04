package controller

import (
	"encoding/json"
	dao "github.com/ericmdantas/simple_go/dao"
	_ "github.com/ericmdantas/simple_go/model"
	router "github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func Cria(w http.ResponseWriter, req *http.Request, _ router.Params) {
	inf, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	var m map[string]interface{}

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(inf, &m)

	if err != nil {
		panic(err)
	}

	dao.Cria(m)

	w.WriteHeader(http.StatusCreated)
}
