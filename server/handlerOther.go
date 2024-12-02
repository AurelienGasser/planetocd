package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aureliengasser/planetocd/articles"
	"github.com/aureliengasser/planetocd/server/tags"
	"github.com/snabb/sitemap"
)

func handleArticlesJson(w http.ResponseWriter, r *http.Request) {
	res := make([]articles.ArticleMetadata, 0)
	ids := make(map[int]bool)

	for _, articles := range allArticles {
		for id, article := range articles {
			if !ids[id] {
				res = append(res, article.OriginalMetadata)
				ids[id] = true
			}
		}
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.Write(jsonBytes)
}

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
	fmt.Println("Handling index")
	canonicalURL := mustGetURL("index_en", "")

	title := SiteName + " - Knowledge base about Obsessive Compulsive Disorder (OCD)"

	p := getViewModel("index_en", r, canonicalURL, title, "", nil)
	p.Meta.DisableHeaderLinks = true
	RenderTemplate(w, p)
}
