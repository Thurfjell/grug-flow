// Package compose turns a PageSpec into a rendered UI via layout + widgets
package compose

import (
	"fmt"
	"html/template"
	"log"
)

type Layout interface {
	Render(page PageSpec, widgets []template.HTML) (string, error)
}

type LayoutResolver interface {
	Get(name string) Layout
}

type WidgetRenderer interface {
	Render(spec WidgetSpec) (template.HTML, error)
}

type WidgetSpec struct {
	Name   string
	Params map[string]string
	URL    string
}

type PageSpec struct {
	Title   string
	Layout  string
	Widgets []WidgetSpec
}

func Render(resolver LayoutResolver, renderer WidgetRenderer, page PageSpec) (string, error) {
	layout := resolver.Get(page.Layout)

	if layout == nil {
		return "", fmt.Errorf("unknown layout: %s", page.Layout)
	}

	if len(page.Widgets) == 0 {
		return layout.Render(page, nil)
	}

	widgets := make([]template.HTML, 0, len(page.Widgets))

	for _, w := range page.Widgets {
		if w.Name == "" {
			log.Println("skipping widget with no name")
			continue
		}
		html, err := renderer.Render(w)
		if err != nil {
			return "", fmt.Errorf("render widget %q: %w", w.Name, err)
		}
		widgets = append(widgets, html)
	}
	return layout.Render(page, widgets)
}
