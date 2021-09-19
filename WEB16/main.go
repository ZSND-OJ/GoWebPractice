package main

import (
	"log"
	"net/http"

	"gopractice/goblock/goweb/web16/app"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeaHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
