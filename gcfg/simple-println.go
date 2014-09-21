package main

import (
	"log"

	"code.google.com/p/gcfg"
)

type Config struct {
	Section struct {
		Name string
		Flg  bool
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf("err :%v\n", err)
	}
}

func main() {
	var cfg Config

	err := gcfg.ReadFileInto(&cfg, "myconfig.gcfg")
	check(err)

	println(cfg.Section.Name)
}
