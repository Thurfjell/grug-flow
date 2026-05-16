package compose

import "core/internal/core"

type Resolver struct {
	layouts  map[string]Layout
	navItems []core.NavItem
}

func NewResolver(layouts map[string]Layout) *Resolver {
	return &Resolver{
		layouts: layouts,
	}
}

func (r *Resolver) Get(name string) Layout {
	return r.layouts[name]
}

func (r *Resolver) SetNavItems(items []core.NavItem) {
	r.navItems = items
}

func (r *Resolver) GetNavItems() []core.NavItem {
	return r.navItems
}
