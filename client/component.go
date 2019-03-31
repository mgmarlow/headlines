package client

import (
	ui "github.com/gizak/termui/v3"
)

// Component is the top-level interface for all widgets
// that will be added to the headlines dashboard.
type Component interface {
	Render(renderer func(...ui.Drawable))
	HandleEvents(event ui.Event, renderer func(...ui.Drawable))
}
