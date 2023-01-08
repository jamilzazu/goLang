package main

import (
	"log"
	"net/http"
	"time"
)

func handler(response http.ResponseWriter, request *http.Request) {
	newContext := request.Context()
	log.Println("Request started")
	defer log.Println("Request completed")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request successfully processed")
		response.Write([]byte("Request successfully processed"))
	case <-newContext.Done():
		log.Println("Request canceled by the customer")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
