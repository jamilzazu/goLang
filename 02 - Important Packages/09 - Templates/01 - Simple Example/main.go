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

func executeClientTemplate() {
	newTemplate := newTemplate("ClientTemplate")
	newTemplate, err := newTemplate.Parse("Client: {{.Name}} - Age: {{.Age}}")
	messageError(err)
	err = newTemplate.Execute(os.Stdout, createClient())
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
