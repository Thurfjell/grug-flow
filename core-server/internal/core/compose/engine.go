// Package compose turns a PageSpec into a rendered UI via layout + widgets
package compose

import (
	"core/internal/core"
	"fmt"
)

type Layout interface {
	Render(page PageSpec) (string, error)
}

type LayoutResolver interface {
	Get(name string) Layout
	GetNavItems() []core.NavItem
}

type WidgetSpec struct {
	Name string
	URL  string
}

type PageSpec struct {
	Title    string
	Layout   string
	Href     string
	Widgets  []WidgetSpec
	NavItems []core.NavItem
}

func Render(resolver LayoutResolver, page PageSpec) (string, error) {
	page.NavItems = resolver.GetNavItems()
	layout := resolver.Get(page.Layout)

	if layout == nil {
		return "", fmt.Errorf("unknown layout: %s", page.Layout)
	}

	return layout.Render(page)
}
