/**
 * http://golang.org/pkg/text/template/
 *
 */

package main

//import "fmt"
import "os"
import "text/template"

func main() {

	type Inventory struct {
		Material string
		Count    uint
	}

	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")

	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)

	if err != nil {
		panic(err)
	}

}
