package dashboard

import (
	"log"
	"net/http"

	"core/internal/core/compose"
	corehttp "core/internal/core/http"
)

type Handler struct {
	Resolver compose.LayoutResolver
	Renderer compose.WidgetRenderer
}

func (h *Handler) Routes() []corehttp.Route {
	return []corehttp.Route{
		{Method: "GET", Path: "/dashboard", Handler: h.GetDashboard},
	}
}

func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	page := DashboardPage()

	html, err := compose.Render(h.Resolver, h.Renderer, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	_, err = w.Write([]byte(html))
	if err != nil {
		log.Printf("failed to get dashboard: %v", err)
	}
}
