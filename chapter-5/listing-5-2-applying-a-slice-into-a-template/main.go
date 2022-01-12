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

const tmpl = `Notes are:
{{range .}}Title: {{.Title}}, Description: {{.Description}}
{{end}}
`

func main() {
	// Create slice of Note objects
	notes := []Note{
		{"text/template", "Template generates textual output"},
		{"html/template", "Template generates HTML output"},
	}

	// Create a new template with a name and parse some content
	// to generate a template
	t, err := template.New("note").Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// Applies a parsed template to the slice of Note objects
	if err := t.Execute(os.Stdout, notes); err != nil {
		log.Fatal("Parse: ", err)
		return
	}
}