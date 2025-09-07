package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)


type application struct{
	config config

}


type config struct{
	addr string
}


func (app *application) mount() *chi.Mux{
	r:=chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router){
	r.Get("/health",app.healthSttaus)
	r.Get("/wealth",app.wealthSttaus)

	})
	
	return r
}


func (app *application) run(mux http.Handler) error{
	serv:=&http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second*30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Second*60,

	}

	return serv.ListenAndServe()
}
