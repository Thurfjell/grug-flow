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

func (l *layout) Render(page compose.PageSpec) (string, error) {
	var buf bytes.Buffer

	err := l.tmpl.ExecuteTemplate(&buf, "title-content.html", page)
	if err != nil {
		return "", fmt.Errorf("render title-content: %w", err)
	}

	return buf.String(), nil
}

//go:embed template/*
var templateFS embed.FS

func New() (*layout, error) {
	tmpl, err := template.ParseFS(templateFS, "template/title-content.html")
	if err != nil {
		return nil, err
	}

	return &layout{
		tmpl: tmpl,
	}, nil
}
