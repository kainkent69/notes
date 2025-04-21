package main

import (
	"log"
	"main/data"
	"main/middlewares"
	"main/router"
	"net/http"
)

var Router http.ServeMux

func main() {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	Router := http.NewServeMux()
	db := data.DB()
	router.Init(db, Router)
	// create a new server
	server := http.Server{
		Addr:    ":8000",
		Handler: middlewares.Log(Router),
	}

	log.Printf("Server is running at port 8000")
	log.Fatal("server fail ", server.ListenAndServe())

}
