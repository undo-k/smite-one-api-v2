package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func vueRoutes(router *chi.Mux) {
	webDir := http.FileServer(http.Dir("./web/app/dist/"))
	router.Handle("/*", http.StripPrefix("/", webDir))
}
