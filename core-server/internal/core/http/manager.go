package http

import (
	"context"
	"core/internal/core"
	"fmt"
	"log"
	nethttp "net/http"
	"time"
)

type manager struct {
	server   *nethttp.Server
	Addr     string
	NavItems []core.NavItem
}

func New(groups ...[]core.Route) *manager {

	mux := nethttp.NewServeMux()
	var navItems []core.NavItem
	for _, group := range groups {
		for _, r := range group {
			if r.Nav != nil {
				navItems = append(navItems, core.NavItem{
					Label: r.Nav.Label,
					Href:  r.Path,
				})
			}

			// If Method is empty, just use Path (matches all methods in Go 1.22+)
			pattern := r.Path
			if r.Method != "" {
				pattern = fmt.Sprintf("%s %s", r.Method, r.Path)
			}
			log.Println("Path", r.Path)

			mux.HandleFunc(pattern, r.Handler)
		}
	}

	// mux.HandleFunc("/widgets/nav/", navHandler(navItems))

	s := &nethttp.Server{
		Addr:        ":1337",
		IdleTimeout: 5 * time.Minute,
		Handler:     mux,
	}

	return &manager{
		server:   s,
		Addr:     s.Addr,
		NavItems: navItems,
	}
}

// func navHandler(items []NavItem) nethttp.HandlerFunc {
// 	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
// 		current := r.Header.Get("Hx-Current-Url")
// 		w.Header().Set("Content-Type", "text/html")
// 		fmt.Fprint(w, renderNav(items, current))
// 	}
// }

// func renderNav(items []NavItem, currentURL string) string {
// 	links := ""
// 	for _, item := range items {
// 		active := strings.Contains(currentURL, item.Href)
// 		activeClass := "text-zinc-400 hover:text-zinc-200 hover:bg-zinc-800"
// 		if active {
// 			activeClass = "text-indigo-400 bg-indigo-950"
// 		}
// 		links += fmt.Sprintf(
// 			`<a href="%s" class="px-3 py-1.5 rounded-lg text-sm %s">%s</a>`,
// 			item.Href, activeClass, item.Label,
// 		)
// 	}

// 	return fmt.Sprintf(`
// <nav class="border-b border-zinc-800 bg-zinc-950">
//   <div class="max-w-7xl mx-auto px-4 flex items-center justify-between h-14">
//     <span class="text-sm font-mono text-zinc-400">grug<span class="text-indigo-400">flow</span></span>
//     <div class="hidden md:flex items-center gap-1">%s</div>
//     <button onclick="document.getElementById('mobile-nav').classList.toggle('hidden')"
//       class="md:hidden text-zinc-400 border border-zinc-800 rounded-lg p-1.5">
//       <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
//         <line x1="3" y1="6" x2="21" y2="6"/><line x1="3" y1="12" x2="21" y2="12"/><line x1="3" y1="18" x2="21" y2="18"/>
//       </svg>
//     </button>
//   </div>
//   <div id="mobile-nav" class="hidden md:hidden border-t border-zinc-800 px-4 py-2 flex flex-col gap-1">%s</div>
// </nav>`, links, links)
// }

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
