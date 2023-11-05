package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"mysrvr/app/controller"
)

func Init() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/login", controller.LoginPOST)
	r.Get("/logout", controller.LogoutGET)
	r.Post("/register", controller.RegisterPOST)
	
	return r
}
