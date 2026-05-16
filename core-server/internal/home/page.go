package home

import "core/internal/core/compose"

func HomePage() compose.PageSpec {
	return compose.PageSpec{
		Title:  "Home",
		Layout: "title-content",
		Href:   "/",
		Widgets: []compose.WidgetSpec{
			{Name: "home", URL: "/widgets/home/"},
		},
	}
}
