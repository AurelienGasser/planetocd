package viewModel

import (
	"html/template"
	"net/url"
)

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
	URL       *url.URL
	Title     string
	HTMLShort template.HTML
	Tags      []string
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
