package main

import (
	"net/http"
	"text/template"
)

type Client struct {
	Name string
	Age  int
}

func createClient() []Client {
	return []Client{{"Jamil", 34}, {"Cristiano", 30}, {"Flavio", 40}, {"Gustavo", 34}}
}

func newTemplate(templateName string) *template.Template {
	return template.New(templateName)
}

func newClientTemplateMust() *template.Template {
	nameTemplate := "clientTemplate.html"
	newTemplate := newTemplate(nameTemplate)
	return template.Must(newTemplate.ParseFiles(nameTemplate))
}

func serverClient() {
	http.HandleFunc("/", handlerClient)
	newClientServer()
}

func newClientServer() {
	http.ListenAndServe(":8282", nil)
}

func handlerClient(response http.ResponseWriter, _ *http.Request) {
	err := newClientTemplateMust().Execute(response, createClient())
	messageError(err)

}
func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	serverClient()
}
