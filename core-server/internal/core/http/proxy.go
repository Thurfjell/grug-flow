package http

import (
	"core/internal/core"
	"fmt"
	nethttp "net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"sync"
)

type ProxyManager struct {
	proxy  *WidgetProxyHandler
	routes []core.Route
}

func NewProxyManager(req WidgetRegistry) *ProxyManager {
	return &ProxyManager{
		proxy: &WidgetProxyHandler{Registry: req},
	}
}

func (pm *ProxyManager) Add(name, target string) *ProxyManager {
	pm.proxy.Registry.Add(name, target)

	path := fmt.Sprintf("/widgets/%s/", name)
	pm.routes = append(pm.routes, core.Route{
		Method:  "",
		Path:    path,
		Handler: pm.proxy.ServeHTTP,
	})

	return pm
}

func (pm *ProxyManager) Routes() []core.Route {
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
	mu      sync.RWMutex
}

func NewMemRegistry() *MemRegistry {
	return &MemRegistry{Widgets: make(map[string]string)}
}

func (m *MemRegistry) Get(name string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	url, ok := m.Widgets[name]
	return url, ok
}

func (m *MemRegistry) Add(name, value string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
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
			// Get the "extra" stuff after /widgets/name/
			subPath := strings.Join(parts[1:], "/")

			// Join the Registry path (/todos) with the subpath (form)
			// This ensures /widgets/test/form -> /todos/form
			pr.Out.URL.Path = path.Join(targetURL.Path, subPath)

			// Important: path.Join strips trailing slashes,
			// if your Elysia route needs it, add it back:
			if strings.HasSuffix(r.URL.Path, "/") && !strings.HasSuffix(pr.Out.URL.Path, "/") {
				pr.Out.URL.Path += "/"
			}
		},
	}

	proxy.ServeHTTP(w, r)
}
