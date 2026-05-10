package http

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	Middlewares []Middleware
}

type RouteGroup struct {
	Prefix      string
	Middlewares []Middleware
	Routes      []Route
}

func chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}

func RegisterRoutes(mux *http.ServeMux, routes []Route) {
	for _, r := range routes {
		handler := chain(r.Handler, r.Middlewares...)

		mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.Path), func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})
	}
}
