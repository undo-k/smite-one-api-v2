package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/undo-k/smite-one-api-v2/internal/config"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func Handler(router *chi.Mux, a *config.AppConfig) {
	Repo = NewRepo(a)

	router.Use(chimiddle.StripSlashes)

	router.Get("/api/v2/gods", Repo.GetGods)
	router.Get("/api/v2/gods/{godId}", Repo.GetGodById)

	staticImages(router)
	vueRoutes(router)

}
