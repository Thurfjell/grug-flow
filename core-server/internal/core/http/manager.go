package http

import (
	"context"
	"fmt"
	"log"
	nethttp "net/http"
	"time"
)

type manager struct {
	server *nethttp.Server
	Addr   string
}

func New(groups ...[]Route) *manager {

	mux := nethttp.NewServeMux()
	for _, group := range groups {
		for _, r := range group {
			// If Method is empty, just use Path (matches all methods in Go 1.22+)
			pattern := r.Path
			if r.Method != "" {
				pattern = fmt.Sprintf("%s %s", r.Method, r.Path)
			}
			log.Println("Path", r.Path)

			mux.HandleFunc(pattern, r.Handler)
		}
	}

	s := &nethttp.Server{
		Addr:        ":1337",
		IdleTimeout: 5 * time.Minute,
		Handler:     mux,
	}

	return &manager{
		server: s,
		Addr:   s.Addr,
	}
}

func (m *manager) Start() error {
	err := m.server.ListenAndServe()

	if err == nethttp.ErrServerClosed {
		return nil
	}

	return err
}

func (m *manager) Stop(ctx context.Context) error {
	return m.server.Shutdown(ctx)
}
