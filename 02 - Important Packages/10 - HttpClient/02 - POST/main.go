package main

import (
	"bytes"
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
	c := http.Client{Timeout: time.Second}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "jamil"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	messageError(err)

	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
