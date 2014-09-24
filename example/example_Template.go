/**
 * http://golang.org/pkg/text/template/#example_Template
 *
 */

package main

import "log"
import "os"
import "text/template"

func main() {

	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.{{else}}
It is s shame you couldn't make it to the wedding.{{end}}
{{with .Gift}}Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rondney", "", false},
	}

	f, err_f := os.Create("sample.txt")
	t := template.Must(template.New("letter").Parse(letter))
	if err_f != nil {
		log.Println("executing template:", err_f)
	}

	for _, r := range recipients {
		//err := t.Execute(os.Stdout, r)
		err := t.Execute(f, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
