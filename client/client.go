package client

import (
	"strconv"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mgmarlow/headlines/articles"
)

func BuildList(top articles.ArticlesResult) *widgets.List {
	l := widgets.NewList()
	l.Title = "Top Stories"
	l.Rows = extractTitles(top.Articles)
	l.WrapText = true
	l.SetRect(0, 0, 100, 10)
	return l
}

func extractTitles(articles []articles.Article) []string {
	var titles []string
	for i, article := range articles {
		titles = append(titles, "[" + strconv.Itoa(i+1) + "] " + article.Title)
	}
	return titles
}