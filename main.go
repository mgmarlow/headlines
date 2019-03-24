package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/joho/godotenv"
	"github.com/mgmarlow/headlines/articles"
	"github.com/mgmarlow/headlines/client"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment.")
	}

	// TODO: Send API requests after initializing termUI
	// for loading indicator
	top, err := articles.GetTopStories()
	if err != nil {
		log.Fatal(err)
	}

	if err := ui.Init(); err != nil {
		log.Fatal("failed to initialize termui")
	}
	defer ui.Close()

	// Prepare all widgets for dashboard
	widgets := []client.Widget{
		client.NewTopStoriesWidget(top),
	}

	// Initial Render
	for _, widget := range widgets {
		ui.Render(widget.GetElement())
	}

	// Event Handling
	for e := range ui.PollEvents() {
		if e.ID == "q" || e.ID == "<C-c>" {
			break
		}

		for _, widget := range widgets {
			widget.HandleEvents(e, ui.Render)
		}
	}
}
