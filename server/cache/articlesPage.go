package cache

type ArticlesPage struct {
	PageNumber int
	Articles   []*Article
}

func NewArticlesPage(pageNumber int, articles []*Article) *ArticlesPage {
	return &ArticlesPage{
		PageNumber: pageNumber,
		Articles:   articles,
	}
}
