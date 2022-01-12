package main

import (
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title string
	Description string
}

const tmpl = `Note - Title: {{.Title}}, Description: {{.Description}}`

func main() {
	// Create an instance of Note structure
	note := Note{"text/template", "Template generates textual output"}

	/*
	// Create a new template with a name
	t := template.New("note")

	// Parse some content and generate a template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	*/

	// OR
	// Create a new template with a name and parse a text template
	t, err := template.New("note").Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// Applies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}