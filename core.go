package headlines

import (
	"encoding/json"
	"net/http"
)

type Source struct {
	Id string
	Name string
}

type Article struct {
	Source Source
	Author string
	Title string
	Description string
	URL string
	URLToImage string
	PublishedAt string
	Content string
}

type ArticlesResult struct {
	Status string
	TotalResults int
	Articles []Article
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
