package client

import (
	"log"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mgmarlow/headlines/api"
)

// TopStories consists of two widgets, an article link and an article description.
type TopStories struct {
	selectedArticle int
	stories         api.ArticlesResult
	description     *widgets.Paragraph
	link            *widgets.Paragraph
}

// TopStoriesComponent constructs the "Top Stories" widgets.
func TopStoriesComponent() *TopStories {
	topStories, err := api.GetTopStories()
	if err != nil {
		log.Fatal(err)
		return &TopStories{}
	}

	widget := buildDescription(topStories.Articles[0])
	widget.Title = buildTitle(0, topStories.TotalResults)

	return &TopStories{
		selectedArticle: 0,
		stories:         topStories,
		description:     widget,
		link:            buildLink(topStories.Articles[0].URL),
	}
}

// Render renders the link and description widget components using ui.Render.
func (t *TopStories) Render(renderer func(...ui.Drawable)) {
	renderer(t.description, t.link)
}

// HandleEvents provides keyboard event handling. ui.Render is passed
// in as the second argument to ensure re-renders only occur when the
// element is changed.
func (t *TopStories) HandleEvents(event ui.Event, renderer func(...ui.Drawable)) {
	switch event.ID {
	case "<Tab>":
		t.selectedArticle++
		if t.selectedArticle >= t.stories.TotalResults {
			t.selectedArticle = 0
		}
		t.link = buildLink(t.stories.Articles[t.selectedArticle].URL)
		t.description = buildDescription(t.stories.Articles[t.selectedArticle])
		t.description.Title = buildTitle(t.selectedArticle, t.stories.TotalResults)
		t.Render(renderer)
		break
	}
}

func buildDescription(article api.Article) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "\n" + article.Title + "\n\n" + article.Description + "\n\nSample:\n" + article.Content
	p.SetRect(0, 2, 100, 20)
	return p
}

func buildTitle(selected, total int) string {
	var title string

	for i := 0; i < total; i++ {
		if i == selected {
			title = title + "[" + strconv.Itoa(i+1) + "] "
		} else {
			title = title + " " + strconv.Itoa(i+1) + "  "
		}
	}

	return title
}

func buildLink(url string) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = url
	p.SetRect(0, 0, 100, 1)
	p.Border = false
	return p
}
