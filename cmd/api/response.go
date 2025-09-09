package main

import "net/http"



func (app *application) homePage( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Home Page"))
}



func(app *application) profilePage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hey from profile page"))
}

