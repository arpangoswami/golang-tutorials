package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(rw http.ResponseWriter, templateTitle string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templateTitle)
	err := parsedTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("Error parsing template ", err.Error())
		return
	}

}
