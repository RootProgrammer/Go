package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	template1, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}

	template1.ExecuteTemplate(os.Stdout, "template.gohtml", "Hello World!")	
}