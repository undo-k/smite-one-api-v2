package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/undo-k/smite-one-api-v2/internal/cache"
	"github.com/undo-k/smite-one-api-v2/internal/config"
	"github.com/undo-k/smite-one-api-v2/internal/handlers"
	"net/http"
	"os"
)

var app config.AppConfig

func main() {
	log.SetReportCaller(true)

	if os.Getenv("DEBUG_MODE") == "False" {
		app.InProduction = true
	} else {
		app.InProduction = false
	}

	app.UseCache = true

	if app.UseCache {
		godCache, err := cache.CreateGodCache()
		if err != nil {
			log.Error("Failed to create god cache:")
			log.Error(err)
		}
		app.GodCache = godCache
	}

	var router *chi.Mux = chi.NewRouter()
	var corsHandler http.Handler
	if app.InProduction {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"https://smite-one-production.up.railway.app/"},
			AllowedMethods:   []string{http.MethodGet},
			AllowCredentials: false,
		})
		corsHandler = c.Handler(router)
	} else {
		corsHandler = cors.Default().Handler(router)
	}

	handlers.Handler(router, &app)

	fmt.Println("Listening and Learning on localhost:8080")

	log.Error(http.ListenAndServe(":8080", corsHandler))
}
