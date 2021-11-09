package cache

type Articles struct {
	Lang  string
	Pages []*ArticlesPage
}

type ArticlesPage struct {
	PageNumber int
	Articles   []*Article
}
