package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "localhost:3333")
	if err != nil {
		log.Fatal(err)
	}

	println(c, "GET /")
	io.Copy(os.Stdout, c)

	c.Close()
}
