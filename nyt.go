package headlines

import (
	"os"
	"net/http"
	"encoding/json"
)

func GetTopStories() (ArticlesResult, error) {
	articles := ArticlesResult{}
	url := constructURL("https://newsapi.org/v2/top-headlines")

	err := getJson(url, &articles)
	if err != nil {
		return ArticlesResult{}, err
	}

	return articles, nil
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func constructURL(base string) string {
	return base + "?sources=the-new-york-times&apiKey=" + os.Getenv("NEWS_API_KEY")
}