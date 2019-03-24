package client

import (
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mgmarlow/headlines/articles"
)

type topStoriesWidget struct {
	element *widgets.List
}

func NewTopStoriesWidget(topStories articles.ArticlesResult) *topStoriesWidget {
	return &topStoriesWidget{
		element: buildList(topStories),
	}
}

// GetElement allows element access for Widget interface
func (w *topStoriesWidget) GetElement() ui.Drawable {
	return w.element
}

// HandleEvents provides keyboard event handling. ui.Render is passed
// in as the second argument to ensure re-renders only occur when the
// element is changed.
func (w *topStoriesWidget) HandleEvents(event ui.Event, renderer func(...ui.Drawable)) {
	switch event.ID {
	case "j", "<Down>":
		w.element.ScrollDown()
		renderer(w.element)
		break
	case "k", "<Up>":
		w.element.ScrollUp()
		renderer(w.element)
		break
	}
}

func buildList(top articles.ArticlesResult) *widgets.List {
	l := widgets.NewList()
	l.Title = "Top Stories (NYT)"
	l.Rows = extractTitles(top.Articles)
	l.WrapText = true
	l.SetRect(0, 0, 100, 10)
	return l
}

func extractTitles(articles []articles.Article) []string {
	var titles []string
	for i, article := range articles {
		titles = append(titles, "["+strconv.Itoa(i+1)+"] "+article.Title)
		titles = append(titles, "    - "+article.URL)
	}
	return titles
}
