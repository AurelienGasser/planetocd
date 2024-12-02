package server

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/aureliengasser/planetocd/server/cache"
	"github.com/aureliengasser/planetocd/server/viewModel"
)

func NewArticles(currentPageIndex int, pages []*ArticlesPage, articles []*ArticlesArticle) *Articles {
	res := &Articles{
		CurrentPageIndex: currentPageIndex,
		Pages:            pages,
		Articles:         articles,
	}
	res.Pagination = viewModel.GetPagination(res, currentPageIndex+1)
	return res
}

type Articles struct {
	CurrentPageIndex int
	Pagination       *viewModel.Pagination
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
	Illustration *viewModel.ArticleIllustration
}

func (al *Articles) GetPages() []viewModel.PaginationPage {
	res := make([]viewModel.PaginationPage, len(al.Pages))
	for i, v := range al.Pages {
		res[i] = viewModel.PaginationPage(v)
	}
	return res
}

func (alp *ArticlesPage) GetURL() *url.URL {
	return alp.URL
}

func (alp *ArticlesPage) GetPageNumber() int {
	return alp.PageNumber
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)

	title := SiteName + " - " + Translate(lang, "Articles_about_Obsessive_Compulsive_Disorder")
	description := ""

	var pages *cache.Articles = nil
	var ok bool

	if pages, ok = allArticlesPaginated[lang]; !ok {
		http.NotFound(w, r)
		return
	}

	var pageNumber int

	if pageNumber, ok = getPageNumber(r.URL.Query(), len(pages.Pages)); !ok {
		http.NotFound(w, r)
		return
	}

	baseURL := mustGetURL("articles", lang)
	p := getViewModel("articles", r, getArticlesCanonicalURL(baseURL, pageNumber), title, description, nil)
	pageIndex := pageNumber - 1
	page := pages.Pages[pageIndex]

	articles := make([]*ArticlesArticle, len(page.Articles))
	for i, article := range page.Articles {
		articles[i] = &ArticlesArticle{
			Title:        article.Title,
			HTMLShort:    article.Pages[0].HTMLShort,
			URL:          mustGetArticleURL(article.Lang, article.ID, article.Slug),
			Tags:         article.Tags,
			Illustration: article.Illustration,
		}
	}

	pageVms := make([]*ArticlesPage, len(pages.Pages))
	for i := range pages.Pages {
		pageVms[i] = &ArticlesPage{
			PageNumber: i + 1,
			URL:        getArticlesCanonicalURL(baseURL, i+1),
		}
	}
	p.Body = NewArticles(pageIndex, pageVms, articles)
	RenderTemplate(w, p)
}
