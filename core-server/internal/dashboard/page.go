// Package dashboard page
package dashboard

import "core/internal/core/compose"

func DashboardPage() compose.PageSpec {
	return compose.PageSpec{
		Title:  "Dashboard",
		Layout: "title-grid",
		Widgets: []compose.WidgetSpec{
			{Name: "test_1", URL: "/widgets/test/"},
			{Name: "test_2", URL: "/widgets/test/"},
			{Name: "test_3", URL: "/widgets/test/"},
		},
	}
}
