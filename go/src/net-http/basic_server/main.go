package main

import (
	"main/data"
	"main/router"
	"net/http"
)

var Router http.ServeMux

func main() {
	Router := http.NewServeMux()
	db := data.DB()
	router.Init(db, Router)
}
