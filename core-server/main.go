package main

import (
	"log"

	"core/internal/core/compose"
	"core/internal/core/compose/layout/titlegrid"
	"core/internal/core/http"
	"core/internal/dashboard"
)

func main() {
	gridtemplate, err := titlegrid.New()
	if err != nil {
		log.Fatalf("gridtemplate: %v", err)
	}

	resolver := compose.NewResolver(map[string]compose.Layout{
		"title-grid": gridtemplate,
	})

	renderer := compose.StaticWidgetRenderer{}

	dashboardHandler := dashboard.Handler{Resolver: resolver, Renderer: renderer}

	httpManager := http.New(dashboardHandler.Routes()...)

	if err := httpManager.Start(); err != nil {
		log.Fatalf("http manager start: %v", err)
	}
}
