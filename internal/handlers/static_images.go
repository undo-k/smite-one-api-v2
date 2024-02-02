package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/undo-k/smite-one-api-v2/internal/models"
	"net/http"
)

func (m *Repository) StaticImages(w http.ResponseWriter, r *http.Request) {
	var gods []models.God
	for _, v := range m.App.GodCache {
		gods = append(gods, v)
	}
	godBytes, _ := json.MarshalIndent(gods, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(godBytes)
	if err != nil {
		panic(err)
	}
}

func staticImages(router *chi.Mux) {
	staticFiles := http.FileServer(http.Dir("./web/static/images/"))
	router.Handle("/static/images/*", http.StripPrefix("/static/images/", staticFiles))
}
