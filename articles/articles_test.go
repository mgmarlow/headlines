package articles

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestGetJson(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "ok",
			"totalResults": 2,
			"articles": [
				{
					"source": {
						"id": "the-new-york-times",
						"name": "The New York Times"
					},
					"author": null,
					"title": "Will the Mueller Report Be Made Public? 17 Answers to What May Come Next",
					"description": "Maggie Haberman, Michael Schmidt, Mark Mazzetti and more of our journalists explained what the submission of the full report means and what may come next.",
					"url": "https://www.nytimes.com/2019/03/22/us/mueller-questions-answers.html",
					"urlToImage": "https://static01.nyt.com/images/2019/04/22/us/politics/22readers-Mueller-QA/merlin_152505528_c5d01b21-111d-4932-b5c7-261e70b197da-facebookJumbo.jpg",
					"publishedAt": "2019-03-24T14:56:57.632Z",
					"content": "Mr. Barr appeared to be lowering expectations during his Senate confirmation hearings when he said he wanted to be as transparent as possible but might provide simply a summary to Congress.\r\nMaggie Haberman\r\nWe have no information that the full report will be… [+1846 chars]"
				},
				{
					"source": {
						"id": "the-new-york-times",
						"name": "The New York Times"
					},
					"author": null,
					"title": "Its Territory May Be Gone, but the U.S. Fight Against ISIS Is Far From Over",
					"description": "Thousands of Islamic State fighters are still at large in Iraq and Syria, rearming and regrouping. And the group pose threats elsewhere, in Afghanistan, West Africa and the Philippines.",
					"url": "https://www.nytimes.com/2019/03/24/us/politics/us-isis-fight.html",
					"urlToImage": "https://static01.nyt.com/images/2019/03/14/us/politics/24dc-isis1/00dc-isis1-facebookJumbo.jpg",
					"publishedAt": "2019-03-24T10:06:27.63Z",
					"content": "Separate estimates, including one by the United Nations in February, put the groups strength even higher. James F. Jeffrey, the American special envoy for Syria, said this month that there are 15,000 to 20,000 armed Islamic State fighters in Iraq and Syria, a… [+1380 chars]"
				}
			]
		}`)
	}))

	defer ts.Close()

	top := ArticlesResult{}
	err := GetJson(ts.URL, &top)
	if err != nil {
		t.Fatal(err)
	}

	if (top.Status != "ok") {
		t.Fatal("actual status should be 'ok'")
	}

	expected := ArticlesResult{
		Status: "ok",
		TotalResults : 2,
		Articles: []Article{
			Article{
				Source: Source{
					ID: "the-new-york-times",
					Name: "The New York Times",
				},
				Author: "",
				Title: "Will the Mueller Report Be Made Public? 17 Answers to What May Come Next",
				Description: "Maggie Haberman, Michael Schmidt, Mark Mazzetti and more of our journalists explained what the submission of the full report means and what may come next.",
				URL: "https://www.nytimes.com/2019/03/22/us/mueller-questions-answers.html",
				URLToImage: "https://static01.nyt.com/images/2019/04/22/us/politics/22readers-Mueller-QA/merlin_152505528_c5d01b21-111d-4932-b5c7-261e70b197da-facebookJumbo.jpg",
				PublishedAt: "2019-03-24T14:56:57.632Z",
				Content: "Mr. Barr appeared to be lowering expectations during his Senate confirmation hearings when he said he wanted to be as transparent as possible but might provide simply a summary to Congress.\r\nMaggie Haberman\r\nWe have no information that the full report will be… [+1846 chars]",
			},
			Article{
				Source: Source{
					ID: "the-new-york-times",
					Name: "The New York Times",
				},
				Author: "",
				Title: "Its Territory May Be Gone, but the U.S. Fight Against ISIS Is Far From Over",
				Description: "Thousands of Islamic State fighters are still at large in Iraq and Syria, rearming and regrouping. And the group pose threats elsewhere, in Afghanistan, West Africa and the Philippines.",
				URL: "https://www.nytimes.com/2019/03/24/us/politics/us-isis-fight.html",
				URLToImage: "https://static01.nyt.com/images/2019/03/14/us/politics/24dc-isis1/00dc-isis1-facebookJumbo.jpg",
				PublishedAt: "2019-03-24T10:06:27.63Z",
				Content: "Separate estimates, including one by the United Nations in February, put the groups strength even higher. James F. Jeffrey, the American special envoy for Syria, said this month that there are 15,000 to 20,000 armed Islamic State fighters in Iraq and Syria, a… [+1380 chars]",
			},
		},
	}

	if (!cmp.Equal(expected, top)) {
		t.Fatal("JSON result should be marshalled into ArticlesResult{}")
	}
}
