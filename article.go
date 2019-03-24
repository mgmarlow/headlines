package headlines

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
