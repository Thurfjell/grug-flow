// Package http is the core http package blah derp todo
package http

import (
	"log"
	"net/http"
	"time"
)

type HTTPManager struct {
	server *http.Server
	mux    *http.ServeMux
}

func NewHTTPManager(routes []Route) *HTTPManager {
	mux := http.NewServeMux()
	RegisterRoutes(mux, routes)

	return &HTTPManager{
		mux: mux,
		server: &http.Server{
			Addr:        "localhost:1337",
			Handler:     mux,
			IdleTimeout: 5 * time.Minute,
		},
	}
}

func (m *HTTPManager) Start() {
	if err := m.server.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func (m *HTTPManager) Stop() {
	if err := m.server.Close(); err != nil {
		log.Fatalf("server stop error: %v", err)
	}
}
