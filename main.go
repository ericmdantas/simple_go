package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	routes "github.com/ericmdantas/simple_go/routes"
	router "github.com/julienschmidt/httprouter"
	"runtime"
)

const PORT string = ":3333"

func main() {
	r := router.New()
	n := negroni.Classic()

	runtime.GOMAXPROCS(runtime.NumCPU())

	routes.Init(r)

	n.UseHandler(r)

	fmt.Printf("Rodando em: %v\n", PORT)

	n.Run(PORT)
}
