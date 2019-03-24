package main

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"github.com/mgmarlow/headlines/articles"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment.")
	}

	top, err := articles.GetTopStories()
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range top.Articles {
		fmt.Println(article.Description)
	}
}
