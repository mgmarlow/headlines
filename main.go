package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/mgmarlow/headlines/articles"
	"github.com/mgmarlow/headlines/client"
	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment.")
	}

	top, err := articles.GetTopStories()
	if err != nil {
		log.Fatal(err)
	}

 	if err := ui.Init(); err != nil {
		log.Fatal("failed to initialize termui: %v", err)
	} 
	defer ui.Close()

	list := client.BuildList(top)
	ui.Render(list)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
