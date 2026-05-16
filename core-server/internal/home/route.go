package home

import (
	"log"
	"net/http"

	"core/internal/core"
	"core/internal/core/compose"
)

type Handler struct {
	Resolver compose.LayoutResolver
}

func (h *Handler) Routes() []core.Route {
	return []core.Route{
		{
			Method: "",
			Path:   "/",
			Nav: &core.NavItem{
				Label: "Home",
			},
			Handler: h.GetHome,
		},
	}
}

func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) {
	page := HomePage()

	html, err := compose.Render(h.Resolver, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	_, err = w.Write([]byte(html))
	if err != nil {
		log.Printf("failed to get home: %v", err)
	}
}
