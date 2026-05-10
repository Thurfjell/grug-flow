package http

import (
	"context"
	"fmt"
	nethttp "net/http"
	"time"
)

type manager struct {
	server *nethttp.Server
	Addr   string
}

func New(routes ...Route) *manager {
	mux := nethttp.NewServeMux()
	for _, r := range routes {
		pattern := fmt.Sprintf("%s %s", r.Method, r.Path)

		mux.HandleFunc(pattern, r.Handler)
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
