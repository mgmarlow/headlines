package articles

import (
	"os"
)

func GetTopStories() (ArticlesResult, error) {
	articles := ArticlesResult{}
	url := constructURL("https://newsapi.org/v2/top-headlines")

	err := GetJson(url, &articles)
	if err != nil {
		return ArticlesResult{}, err
	}

	return articles, nil
}

func constructURL(base string) string {
	return base + "?sources=the-new-york-times&apiKey=" + os.Getenv("NEWS_API_KEY")
}