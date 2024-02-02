package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/undo-k/smite-one-api-v2/internal/api"
	"net/http"
)

func (m *Repository) GetGodById(w http.ResponseWriter, r *http.Request) {
	godId := chi.URLParam(r, "godId")
	god, ok := m.App.GodCache[godId]
	if ok {
		godBytes, _ := json.MarshalIndent(god, "", "\t")

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(godBytes)
		if err != nil {
			panic(err)
		}
	} else {
		api.RequestErrorHandler(w, errors.New("god not found"))
	}

}
