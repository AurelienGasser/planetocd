package cache

type Articles struct {
	Lang  string
	Pages []*ArticlesPage
}

func NewArticles(lang string, pages []*ArticlesPage) *Articles {
	return &Articles{
		Lang:  lang,
		Pages: pages,
	}
}
