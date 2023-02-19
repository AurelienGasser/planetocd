package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/aureliengasser/planetocd/server/cache"
	"github.com/aureliengasser/planetocd/server/tags"
	"github.com/aureliengasser/planetocd/server/viewModel"
	"github.com/gorilla/mux"
	"github.com/snabb/sitemap"
)

func handleRobots(w http.ResponseWriter, r *http.Request) {
	lines := make([]string, 0)
	lines = append(lines, "User-agent: *")
	lines = append(lines, "Allow: /")
	url, err := router.Get("sitemap").URL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	lines = append(lines, fmt.Sprintf("Sitemap: %v", url))

	w.WriteHeader(http.StatusOK)
	for _, line := range lines {
		w.Write([]byte(line + "\n"))
	}
}

func handleSitemap(w http.ResponseWriter, r *http.Request) {
	urls := make([]*url.URL, 0)
	urls = append(urls, mustGetURL("index_en", ""))

	for lang, articles := range allArticlesPaginated {
		baseURL := mustGetURL("articles", lang)
		for _, page := range articles.Pages {
			urls = append(urls, getArticlesCanonicalURL(baseURL, page.PageNumber))
			for _, article := range page.Articles {
				articleBaseURL := mustGetArticleURL(lang, article.ID, article.Slug)
				for _, articlePage := range article.Pages {
					articleURL := getArticlesCanonicalURL(articleBaseURL, articlePage.PageNumber)
					urls = append(urls, articleURL)
				}
			}
		}
		urls = append(urls, mustGetURL("about", lang))
	}
	for lang, tags := range tags.Tags {
		for _, tag := range tags {
			urls = append(urls, mustGetTagURL(lang, tag))
		}
	}
	s := sitemap.New()

	for _, url := range urls {
		s.Add(&sitemap.URL{Loc: url.String()})
	}
	s.WriteTo(w)
}

func handleEnglishIndex(w http.ResponseWriter, r *http.Request) {
	// lang := getLanguage(r)
	// if lang != "" {
	// 	url := mustGetURL("articles", lang)
	// 	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
	// }
	fmt.Println("Handing index")
	canonicalURL := mustGetURL("index_en", "")

	title := SiteName + " - Knowledge base about Obsessive Compulsive Disorder (OCD)"

	p := getViewModel("index_en", r, canonicalURL, title, "", nil)
	p.Meta.DisableHeaderLinks = true
	RenderTemplate(w, p)
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)

	title := SiteName + " - " + Translate(lang, "Articles_about_Obsessive_Compusive_Disorder")
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

	articles := make([]*viewModel.ArticlesArticle, len(page.Articles))
	for i, article := range page.Articles {
		articles[i] = &viewModel.ArticlesArticle{
			Title:        article.Title,
			HTMLShort:    article.Pages[0].HTMLShort,
			URL:          mustGetArticleURL(article.Lang, article.ID, article.Slug),
			Tags:         article.Tags,
			Illustration: article.Illustration,
		}
	}

	pageVms := make([]*viewModel.ArticlesPage, len(pages.Pages))
	for i := range pages.Pages {
		pageVms[i] = &viewModel.ArticlesPage{
			PageNumber: i + 1,
			URL:        getArticlesCanonicalURL(baseURL, i+1),
		}
	}

	p.Body = viewModel.NewArticles(pageIndex, pageVms, articles)
	RenderTemplate(w, p)
}

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
	title := tagName + " - " + Translate(lang, "Articles_about_Obsessive_Compusive_Disorder")
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

	articles := make([]*viewModel.ArticlesArticle, len(page.Articles))
	for i, article := range page.Articles {
		articles[i] = &viewModel.ArticlesArticle{
			Title:        article.Title,
			HTMLShort:    article.Pages[0].HTMLShort,
			URL:          mustGetArticleURL(article.Lang, article.ID, article.Slug),
			Tags:         article.Tags,
			Illustration: article.Illustration,
		}
	}

	pageVms := make([]*viewModel.ArticlesPage, len(pages.Pages))
	for i := range pages.Pages {
		pageVms[i] = &viewModel.ArticlesPage{
			PageNumber: i + 1,
			URL:        getArticlesCanonicalURL(baseURL, i+1),
		}
	}

	p.Body = viewModel.NewTag(tag, viewModel.NewArticles(pageIndex, pageVms, articles))
	RenderTemplate(w, p)
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	article, err := getArticle(lang, id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	title := article.Title + " - " + SiteName
	description := ""
	pageNumber := 1
	pageNumberStr := r.URL.Query().Get("page")
	if pageNumberStr != "" {
		tmp, err := strconv.Atoi(pageNumberStr)
		if err == nil && tmp >= 1 && tmp <= len(article.Pages) {
			pageNumber = tmp
		}
	}

	canonicalURL, _ := router.Get("article").URL("language", lang, "id", idStr, "slug", article.Slug)
	if pageNumber > 1 {
		canonicalURL.RawQuery = fmt.Sprintf("page=%v", pageNumber)
	}

	vm := getViewModel("article", r, canonicalURL, title, description, article.Illustration.Md())
	suggestions, err := getArticleSuggestions(article)

	if err != nil {
		log.Printf("Error getting suggestions for article %v in lang %v: %v\n", article.ID, lang, err)
	}

	articleVM := articleViewModel{
		Article:          article,
		CurrentPageIndex: pageNumber - 1,
		Pagination:       viewModel.GetPagination(article, pageNumber),
		Suggestions:      suggestions,
	}
	vm.Body = articleVM
	RenderTemplate(w, vm)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL := mustGetURL("about", lang)
	title := Translate(lang, "About") + " - " + SiteName

	p := getViewModel(fmt.Sprintf("about_%v", lang), r, canonicalURL, title, "", nil)
	RenderTemplate(w, p)
}
