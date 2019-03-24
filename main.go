package headlines

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment.")
	}

	articles, err := GetTopStories()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(articles.Status)
}
