// Package titlegrid create title grid layout template
package titlegrid

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"core/internal/core/compose"
)

type layout struct {
	tmpl *template.Template
}

func (l *layout) Render(page compose.PageSpec, widgets []template.HTML) (string, error) {
	var buf bytes.Buffer
	data := struct {
		Page    compose.PageSpec
		Widgets []template.HTML
	}{
		Page:    page,
		Widgets: widgets,
	}

	err := l.tmpl.ExecuteTemplate(&buf, "title-grid.html", data)
	if err != nil {
		return "", fmt.Errorf("render: %w", err)
	}

	return buf.String(), err
}

//go:embed template/*.html
var templateFS embed.FS

func New() (*layout, error) {
	tmpl, err := template.ParseFS(templateFS, "template/title-grid.html")
	if err != nil {
		return nil, err
	}

	return &layout{
		tmpl,
	}, nil
}
