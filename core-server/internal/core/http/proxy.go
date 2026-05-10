package http

import (
	"fmt"
	nethttp "net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ProxyManager struct {
	proxy  *WidgetProxyHandler
	routes []Route
}

func NewProxyManager(req WidgetRegistry) *ProxyManager {
	return &ProxyManager{
		proxy: &WidgetProxyHandler{Registry: req},
	}
}

func (pm *ProxyManager) Add(name, target string) *ProxyManager {
	pm.proxy.Registry.Add(name, target)

	path := fmt.Sprintf("/widgets/%s/", name)
	pm.routes = append(pm.routes, Route{
		Method:  "GET",
		Path:    path,
		Handler: pm.proxy.ServeHTTP,
	})

	return pm
}

func (pm *ProxyManager) Routes() []Route {
	return pm.routes
}

type WidgetRegistry interface {
	Get(name string) (string, bool)
	Add(name, value string) bool
}

type WidgetProxyHandler struct {
	Registry WidgetRegistry
	Client   *nethttp.Client
}

type MemRegistry struct {
	Widgets map[string]string
}

func NewMemRegistry() *MemRegistry {
	return &MemRegistry{Widgets: make(map[string]string)}
}

func (m *MemRegistry) Get(name string) (string, bool) {
	url, ok := m.Widgets[name]
	return url, ok
}

func (m *MemRegistry) Add(name, value string) bool {
	m.Widgets[name] = value
	return true
}

func (h *WidgetProxyHandler) ServeHTTP(w nethttp.ResponseWriter, r *nethttp.Request) {
	trimmed := strings.TrimPrefix(r.URL.Path, "/widgets/")
	parts := strings.Split(trimmed, "/")
	name := parts[0]

	rawTarget, ok := h.Registry.Get(name)
	if !ok {
		nethttp.Error(w, "widget not found", nethttp.StatusNotFound)
		return
	}

	targetURL, err := url.Parse(rawTarget)
	if err != nil {
		nethttp.Error(w, "invalid widget url", 500)
		return
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(targetURL)
			pr.Out.URL.Path = "/" + strings.Join(parts[1:], "/")
		},
	}

	proxy.ServeHTTP(w, r)
}
