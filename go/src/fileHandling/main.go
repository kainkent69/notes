package main

import (
	"io"
	"log"
)

func main() {
	log.Println("Reading A file")
	fs.ReadFile("./hello.txt")
}
