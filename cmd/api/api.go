package main

import (
	"gophers/internals/store"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct{
	config config
	store store.Storage

}

type config struct{
	addr string
}


func (app*application) mount() *chi.Mux{
		r:=chi.NewRouter()
		r.Use(middleware.Recoverer)
		r.Use(middleware.Logger)

		r.Route("/v1", func(r chi.Router){
			r.Get("/home",app.homePage)
			r.Get("/profile",app.profilePage)
		})
		return r


}

func (app *application) run (mux http.Handler) error{

	serv:=&http.Server{
		Addr: app.config.addr,
		Handler: mux,
		ReadTimeout: time.Second*30,
		WriteTimeout: time.Second*10,
		IdleTimeout: time.Second*60,
		
	}

	return serv.ListenAndServe();

	
}