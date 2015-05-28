package controller

import (
	"encoding/json"
	dao "github.com/ericmdantas/simple_go/dao"
	router "github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

func Cria(w http.ResponseWriter, req *http.Request, _ router.Params) {
	inf, err := ioutil.ReadAll(req.Body)

	var m map[string]interface{}

	if err != nil {
		req.Body.Close()
		log.Println(err)
		return
	}

	err = json.Unmarshal(inf, &m)

	if err != nil {
		req.Body.Close()
		log.Println(err)
		return
	}

	dao.Cria(m)

	req.Body.Close()

	w.WriteHeader(http.StatusCreated)
}
