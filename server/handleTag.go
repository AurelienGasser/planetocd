package server

import (
	"net/http"

	"github.com/aureliengasser/planetocd/server/cache"
	"github.com/aureliengasser/planetocd/server/tags"
	"github.com/gorilla/mux"
)

func handleTag(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	vars := mux.Vars(r)
	tag := vars["tag"]

	var pages *cache.Articles = nil
	var ok bool

	if ok = tags.TagExists(lang, tag); !ok {
		http.NotFound(w, r)
		return
	}

	tagName := TranslateTag(lang, tag)
	title := tagName + " - " + Translate(lang, "Articles_about_Obsessive_Compulsive_Disorder")
	description := tagName

	if pages, ok = allArticlesPaginatedByTag[lang][tag]; !ok {
		http.NotFound(w, r)
		return
	}

	var pageNumber int

	if pageNumber, ok = getPageNumber(r.URL.Query(), len(pages.Pages)); !ok {
		http.NotFound(w, r)
		return
	}

	baseURL := mustGetTagURL(lang, tag)
	p := getViewModel("tag", r, getArticlesCanonicalURL(baseURL, pageNumber), title, description, nil)
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

	p.Body = NewTag(tag, NewArticles(pageIndex, pageVms, articles))
	RenderTemplate(w, p)
}
