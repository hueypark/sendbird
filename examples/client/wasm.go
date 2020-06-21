// +build wasm

package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	f, err := newFrontend()
	if err != nil {
		log.Fatalln(err)
	}

	app.Route("/", f)
	app.Run()
}
