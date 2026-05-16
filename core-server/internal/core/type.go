package core

import (
	"net/http"
)

type NavItem struct {
	Label string
	Href  string
}

type Middleware func(http.Handler) http.Handler

type Route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	Middlewares []Middleware // Not used atm, but just for an idea. Could be used for anything. Route specific MW.
	Nav         *NavItem
}
