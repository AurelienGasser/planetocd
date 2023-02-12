package cache

import (
	"html/template"
	"net/url"

	"github.com/aureliengasser/planetocd/articles"
	"github.com/aureliengasser/planetocd/server/viewModel"
)

type Article struct {
	*articles.Article
	Slug         string
	URL          *url.URL
	Illustration *ArticleIllustration
	Pages        []*ArticlePage
	Tags         []string
	Translators  []string
}

func (a *Article) GetPages() []viewModel.PaginationPage {
	res := make([]viewModel.PaginationPage, len(a.Pages))
	for i, v := range a.Pages {
		res[i] = viewModel.PaginationPage(v)
	}
	return res
}

type ArticlePage struct {
	PageNumber int
	HTML       template.HTML
	HTMLShort  template.HTML
	URL        *url.URL
}

func (p *ArticlePage) GetPageNumber() int {
	return p.PageNumber
}

func (p *ArticlePage) GetURL() *url.URL {
	return p.URL
}
