package client

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Loading provides a simple loading indicator before initialization is finished.
type Loading struct {
	component *widgets.Paragraph
}

// LoadingComponent creates a new paragraph widget.
func LoadingComponent() *Loading {
	loading := widgets.NewParagraph()
	loading.Text = "Loading..."
	loading.SetRect(0, 0, 10, 10)
	loading.Border = false

	return &Loading{
		component: loading,
	}
}

// Render renders the Loading component with ui.Render.
func (l *Loading) Render(renderer func(...ui.Drawable)) {
	renderer(l.component)
}
