package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v1"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		C int
		D []int
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func main() {
	t := T{}

	// yamlの情報を構造体に置き換える
	err := yaml.Unmarshal([]byte(data), &t)
	check(err)
	fmt.Printf("--- t:\n%v\n\n", t)

	// 構造体のメンバにアクセスする事も可能
	println(t.A)

	d, err := yaml.Marshal(&t)
	check(err)
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	check(err)
	fmt.Printf("--- t :\n%v\n\n", m)

	// マップにして要素にアクセスする事も可能
	for key, _ := range m {
		fmt.Printf("%v\n", m[key])
	}

	d, err = yaml.Marshal(&m)
	check(err)
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

}
