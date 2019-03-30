package client

import (
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mgmarlow/headlines/api"
)

type topStoriesWidget struct {
	selectedArticle int
	stories         api.ArticlesResult
	list            *widgets.Paragraph
	description     *widgets.Paragraph
}

// NewTopStoriesWidget constructs the "Top Stories" widget
func NewTopStoriesWidget(topStories api.ArticlesResult) *topStoriesWidget {
	description := buildWidget(topStories.Articles[0])
	description.Title = buildTitle(0, topStories.TotalResults)

	return &topStoriesWidget{
		selectedArticle: 0,
		stories:         topStories,
		description:     description,
	}
}

// Render renders all widget using ui.Render
func (w *topStoriesWidget) Render(renderer func(...ui.Drawable)) {
	renderer(w.description)
}

// HandleEvents provides keyboard event handling. ui.Render is passed
// in as the second argument to ensure re-renders only occur when the
// element is changed.
func (w *topStoriesWidget) HandleEvents(event ui.Event, renderer func(...ui.Drawable)) {
	switch event.ID {
	case "<Tab>":
		w.selectedArticle++
		if w.selectedArticle >= w.stories.TotalResults {
			w.selectedArticle = 0
		}
		w.description = buildWidget(w.stories.Articles[w.selectedArticle])
		w.description.Title = buildTitle(w.selectedArticle, w.stories.TotalResults)
		w.Render(renderer)
		break
	}
}

func buildWidget(article api.Article) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = article.Title + "\n\n" + article.Description + "\n\n" + article.URL
	p.SetRect(0, 0, 100, 10)
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
