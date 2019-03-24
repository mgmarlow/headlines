package client

import (
	ui "github.com/gizak/termui/v3"
)

// Widget is the top-level interface for all widgets
// that will be added to the headlines dashboard.
type Widget interface {
	GetElement() ui.Drawable
	HandleEvents(event ui.Event, renderer func(...ui.Drawable))
}
