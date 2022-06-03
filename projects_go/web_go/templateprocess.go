package main

import (
	// "fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	// fmt.Println(os.Getwd())
	template1, err := template.ParseGlob("templates/*.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	template1.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
}
