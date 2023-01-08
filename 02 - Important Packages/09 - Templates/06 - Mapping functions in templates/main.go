package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Client struct {
	Name string
	Age  int
}

func ToUpper(text string) string {
	return strings.ToUpper(text)
}

func createClient() []Client {
	return []Client{{"Jamil", 34}, {"Cristiano", 30}, {"Flavio", 40}, {"Gustavo", 34}}
}

func newTemplate(templateName string) *template.Template {
	return template.New(templateName)
}

func newClientTemplateMust() *template.Template {
	nameTemplate := "content.html"
	newTemplate := newTemplate(nameTemplate)
	newTemplate.Funcs(template.FuncMap{"ToUpper": ToUpper})
	return template.Must(newTemplate.ParseFiles(templatesHtml()...))
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

func templatesHtml() []string {
	return []string{"header.html", "content.html", "footer.html"}
}

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	serverClient()
}
