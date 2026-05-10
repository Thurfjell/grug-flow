// Package titlecontent create title content grid layout template
package titlecontent

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

func (l *layout) Render(page compose.PageSpec, widgets []string) (string, error) {
	var buf bytes.Buffer

	data := struct {
		Page    compose.PageSpec
		Widgets []string
	}{
		Page:    page,
		Widgets: widgets,
	}

	err := l.tmpl.ExecuteTemplate(&buf, "title-content.html", data)
	if err != nil {
		return "", fmt.Errorf("render: %w", err)
	}

	return buf.String(), nil
}

//go:embed template/*
var templateFS embed.FS

func New() (compose.Layout, error) {
	tmpl, err := template.ParseFS(templateFS, "template/title-content.html")
	if err != nil {
		return nil, err
	}

	return &layout{
		tmpl: tmpl,
	}, nil
}
