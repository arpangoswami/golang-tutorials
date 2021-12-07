package main

import (
	"net/http"
)

//Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home-page.html")
}

//About is the about route handler
func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about-page.html")
}
