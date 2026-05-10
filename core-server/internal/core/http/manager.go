package http

import (
	"fmt"
	nethttp "net/http"
)

type manager struct {
	mux *nethttp.ServeMux
}

func New(routes ...Route) *manager {
	mux := nethttp.NewServeMux()
	for _, r := range routes {
		pattern := fmt.Sprintf("%s %s", r.Method, r.Path)

		mux.HandleFunc(pattern, r.Handler)
	}

	return &manager{
		mux: mux,
	}
}

func (m *manager) Start() error {
	return nethttp.ListenAndServe(":1337", m.mux)
}
