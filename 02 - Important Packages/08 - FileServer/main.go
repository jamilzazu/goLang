package main

import (
	"log"
	"net/http"
)

func createFileServer() {
	fileServer := http.FileServer(http.Dir("./public"))
	serverMux := http.NewServeMux()
	serverMux.Handle("/", fileServer)
	serverMux.HandleFunc("/account", handlerAccount)
	log.Fatal(http.ListenAndServe(":8080", serverMux))
}

func handlerAccount(response http.ResponseWriter, _ *http.Request) {
	response.Write([]byte("Account"))
}

func main() {
	createFileServer()
}
