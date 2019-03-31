package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/joho/godotenv"
	"github.com/mgmarlow/headlines/client"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment.")
	}

	if err := ui.Init(); err != nil {
		log.Fatal("failed to initialize termui")
	}
	defer ui.Close()

	client.LoadingComponent().Render(ui.Render)

	// Prepare all widgets for dashboard
	components := []client.Component{
		client.TopStoriesComponent(),
	}

	ui.Clear()

	// Initial Render
	for _, c := range components {
		c.Render(ui.Render)
	}

	// Event Handling
	for e := range ui.PollEvents() {
		if e.ID == "q" || e.ID == "<C-c>" {
			break
		}

		for _, c := range components {
			c.HandleEvents(e, ui.Render)
		}
	}
}
