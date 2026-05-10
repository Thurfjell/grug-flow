package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	dashboardHandler := dashboard.Handler{Resolver: resolver}

	httpManager := http.New(dashboardHandler.Routes()...)

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- httpManager.Start()
	}()

	log.Printf("core server started at - %s", httpManager.Addr)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		log.Fatalf("server failed: %v", err)

	case <-stop:
		log.Println("Signal received, initiating teardown...")
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := httpManager.Stop(shutdownCtx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v", err)
	}

	log.Println("Server exited cleanly.")

}
