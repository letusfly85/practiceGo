/**
 * http://stackoverflow.com/questions/17211027/go-with-parseglob-how-to-render-more-than-two-templates-in-golang
 *
 *
 *
 */

package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {

	pattern := filepath.Join("template", "*.tmpl")
	tmpl := template.Must(template.ParseGlob(pattern))

	err := tmpl.ExecuteTemplate(os.Stdout, "main", nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
