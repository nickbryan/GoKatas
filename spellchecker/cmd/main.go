package main

import (
	"log"

	"github.com/nickbryan/GoKatas/spellchecker"
)

const baseURL = "http://codekata.com/"

func main() {
	if app, err := spellchecker.NewApp(baseURL); err != nil {
		log.Fatalln(err)
	} else {
		log.Fatalln(app.Run())
	}
}
