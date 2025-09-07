package main

import "net/http"


func (app *application) healthSttaus(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Helath status is ok"))
	
}


func (app *application) wealthSttaus(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Wealth status is not ok"))
	
}
