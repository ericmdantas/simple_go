package controller

import (
	"encoding/json"
	dao "github.com/ericmdantas/simple_go/dao"
	model "github.com/ericmdantas/simple_go/model"
	router "github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func Cria(w http.ResponseWriter, req *http.Request, _ router.Params) {
	inf, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	info := model.Info{}

	err = json.Unmarshal(inf, &info)

	if err != nil {
		panic(err)
	}

	go dao.Cria(info)

	w.WriteHeader(http.StatusCreated)
}
