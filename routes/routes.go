package routes

import (
	ctrl "github.com/ericmdantas/simple_go/controllers"
	router "github.com/julienschmidt/httprouter"
)

func Init(router *router.Router) {
	router.POST("/api/alguma-coisa", ctrl.Cria)
}
