package myserver

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		println(e)
		log.Fatal(e)
	}
}

func checkAndExit(e error) {
	if e != nil {
		println(e)
		log.Fatal(e)
		os.Exit(1)
	}
}
