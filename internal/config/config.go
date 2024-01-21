package config

import (
	"github.com/undo-k/smite-one-api-v2/internal/models"
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	GodCache      map[string]models.God
	InfoLog       *log.Logger
	InProduction  bool
}
