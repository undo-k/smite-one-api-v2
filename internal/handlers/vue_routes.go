package handlers

import (
	"github.com/go-chi/chi"
	"net/http"
)

func vueRoutes(router *chi.Mux) {
	webDir := http.FileServer(http.Dir("./web/app/dist/"))
	router.Handle("/*", http.StripPrefix("/", webDir))
}
