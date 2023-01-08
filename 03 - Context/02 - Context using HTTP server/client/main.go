package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	newContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(newContext, "GET", "http://localhost:8080", nil)
	messageError(err)

	response, err := http.DefaultClient.Do(request)
	messageError(err)
	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}
