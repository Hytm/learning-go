package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter() //Call the router
	var message = "Listening on 3000"
	log.Printf("%s",
		message,
	)
	log.Fatal(http.ListenAndServe(":3000", router))
}
