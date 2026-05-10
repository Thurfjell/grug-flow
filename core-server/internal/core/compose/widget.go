package compose

import (
	"fmt"
	"html/template"
)

type StaticWidgetRenderer struct{}

func (r StaticWidgetRenderer) Render(spec WidgetSpec) (template.HTML, error) {
	html := fmt.Sprintf(`
		<div>
			<h2>%s</h2>
			<p>Hello from demo widget</p>
		</div>
		`, spec.Name)

	return template.HTML(html), nil
}
