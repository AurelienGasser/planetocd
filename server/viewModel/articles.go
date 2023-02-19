package viewModel

import (
	"html/template"
	"net/url"
)

// var ArticlesByTag map[string]map[string]*Articles

// func GetArticlesViewModelByTag(lang string, tag string) *Articles {
// 	ensureLoaded()
// }

// func ensureLoaded() {
// 	if ArticlesByTag != nil {
// 		return
// 	}
// 	ArticlesByTag = make(map[string]map[string]*Articles, len(tags.Articles))
// 	for lang, articlesByTag := range tags.Articles {
// 		ArticlesByTag[lang] = make(map[string]*Articles, len(articlesByTag))
// 		for tag, articles := range articlesByTag {
// 			ArticlesByTag[lang][tag] = newArticle()
// 		}
// 	}
// }

func NewArticles(currentPageIndex int, pages []*ArticlesPage, articles []*ArticlesArticle) *Articles {
	res := &Articles{
		CurrentPageIndex: currentPageIndex,
		Pages:            pages,
		Articles:         articles,
	}
	res.Pagination = GetPagination(res, currentPageIndex+1)
	return res
}

type Articles struct {
	CurrentPageIndex int
	Pagination       *Pagination
	Pages            []*ArticlesPage
	Articles         []*ArticlesArticle
}

type ArticlesPage struct {
	URL        *url.URL
	PageNumber int
}

type ArticlesArticle struct {
	URL          *url.URL
	Title        string
	HTMLShort    template.HTML
	Tags         []string
	Illustration *ArticleIllustration
}

func (al *Articles) GetPages() []PaginationPage {
	res := make([]PaginationPage, len(al.Pages))
	for i, v := range al.Pages {
		res[i] = PaginationPage(v)
	}
	return res
}

func (alp *ArticlesPage) GetURL() *url.URL {
	return alp.URL
}

func (alp *ArticlesPage) GetPageNumber() int {
	return alp.PageNumber
}
