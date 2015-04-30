package main

import (
	"fmt"
	routes "github.com/ericmdantas/simple_go/routes"
	router "github.com/julienschmidt/httprouter"
	"net/http"
)

const PORT string = ":3333"

func main() {
	r := router.New()

	routes.Init(r)

	fmt.Printf("Rodando em: %v\n", PORT)

	http.ListenAndServe(PORT, r)
}
