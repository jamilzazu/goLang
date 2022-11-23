package main

import (
	"io"
	"net/http"
)

func requestGet() {
	request := createRequestGet()
	readRequestGet(request)
	closeRequestBody(request)
}

func createRequestGet() *http.Response {
	request, err := http.Get("https://google.com")
	messageError(err)
	return request
}

func readRequestGet(request *http.Response) {
	result, err := io.ReadAll(request.Body)
	messageError(err)
	println(string(result))
}

func closeRequestBody(request *http.Response) {
	defer request.Body.Close()
}

func messageError(message error) {
	if message != nil {
		panic(message)
	}
}

func main() {
	requestGet()
}
