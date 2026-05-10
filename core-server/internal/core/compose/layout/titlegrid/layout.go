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

func (l *layout) Render(page compose.PageSpec) (string, error) {
	var buf bytes.Buffer

	err := l.tmpl.ExecuteTemplate(&buf, "layout.html", page)
	if err != nil {
		return "", fmt.Errorf("render title-grid: %w", err)
	}

	return buf.String(), err
}

//go:embed template/*.html
var templateFS embed.FS

func New() (*layout, error) {
	tmpl, err := template.ParseFS(templateFS, "template/layout.html")
	if err != nil {
		return nil, err
	}

	return &layout{
		tmpl,
	}, nil
}
