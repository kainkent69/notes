package main

import (
	"fmt"
	"log"
	"main/middlewares"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
		err := r.Body.Close()
		if err != nil {
			log.Panic(err)
		}
	})

	router.HandleFunc("/echo/{msg}", func(w http.ResponseWriter, r *http.Request) {
		msg := r.PathValue("msg")
		str := fmt.Sprintf("You Said! msg: %v", msg)
		w.Header().Add(http.CanonicalHeaderKey("Status-code"), "201")
		w.Write([]byte(str))
		r.Body.Close()
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: middlewares.Log(router),
	}

	log.Printf("Server is running at port 8000")
	log.Fatal(server.ListenAndServe())

}
