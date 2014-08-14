package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {

	//register the function
	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	const templateText = `
Input:    {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{printf "%q" . | title}} 
`

	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
