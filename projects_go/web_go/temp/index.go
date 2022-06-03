package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name    string
	Hobbies []string
}

const data = `{{$name := .Name}}
{{range .Hobbies}}{{$name}} Loves {{.}}
{{end}}`

func main() {
	// name := Person{"Niha"}
	person := Person{
		Name:    "Niha",
		Hobbies: []string{"Watching Movies", "Eating Icecream"},
	}

	template1 := template.New("Portfolio")

	// template1, err := template1.Parse("Hello World! I'm {{.Name}}.")
	template1, err := template1.Parse(data)

	if err != nil {
		log.Fatal(err)
		return
	}

	// template1.Execute(os.Stdout, name)
	template1.Execute(os.Stdout, person)
}
