package server

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	Host string
	Port int
}

func Start(config ServerConfig, r *chi.Mux) {
	http.ListenAndServe(config.Host + ":" + fmt.Sprint(config.Port), r)
}
