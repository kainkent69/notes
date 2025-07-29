package main

import (
	"io/fs" 
	"log"
)

func main() {
	log.Println("Reading A file")
	fs.ReadFile("./hello.txt",2)
	
}
