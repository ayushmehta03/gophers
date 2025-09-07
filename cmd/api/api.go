package main

import (
	"net/http"
	"time"
)


type application struct{
	config config

}


type config struct{
	addr string
}


func (app *application) mount() *http.ServeMux{
	mux:=http.NewServeMux()

	mux.HandleFunc("GET /health",app.healthSttaus)

	mux.HandleFunc("GET /wealth", app.wealthSttaus)
	return mux
}


func (app *application) run(mux *http.ServeMux) error{
	serv:=&http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second*30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Second*60,

	}

	return serv.ListenAndServe()
}
