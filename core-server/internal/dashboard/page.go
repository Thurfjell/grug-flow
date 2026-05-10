// Package dashboard page
package dashboard

import "core/internal/core/compose"

func DashboardPage() compose.PageSpec {
	return compose.PageSpec{
		Title:  "Dashboard",
		Layout: "title-grid",
		Widgets: []compose.WidgetSpec{
			{Name: "todos", URL: "/widgets/todos/"},
			{Name: "todos_form", URL: "/widgets/todos_form/"},
			{Name: "test_3", URL: "/widgets/test/"},
		},
	}
}
