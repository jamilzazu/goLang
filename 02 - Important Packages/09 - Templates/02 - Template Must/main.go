package main

import (
	"os"
	"text/template"
)

type Client struct {
	Name string
	Age  int
}

func createClient() Client {
	return Client{"Jamil", 34}
}

func newTemplate(templateName string) *template.Template {
	return template.New(templateName)
}

func newClientTemplateMust() *template.Template {
	newTemplate := newTemplate("ClientTemplate")
	return template.Must(newTemplate.Parse("Client: {{.Name}} - Age: {{.Age}}"))
}

func executeClientTemplate() {
	err := newClientTemplateMust().Execute(os.Stdout, createClient())
	messageError(err)
}

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	executeClientTemplate()
}
