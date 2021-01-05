package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

var router *mux.Router
var isLocalEnvironment bool

// Listen ...
func Listen(port int, isLocal bool) {
	isLocalEnvironment = isLocal
	router = mux.NewRouter()

	if isLocal {
		router = router.
			Schemes("http").
			Host(fmt.Sprintf("localhost:%v", port)).
			Subrouter()
	} else {
		router = router.
			Schemes("https", "http").
			Host(Host).
			Subrouter()
	}

	router.Path("/").HandlerFunc(handleEnglishIndex).Name("index_en")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))).Name("static")

	s := router.PathPrefix("/{language}").Subrouter()
	s.HandleFunc("/about", handleAbout).Name("about")
	s.HandleFunc("", handleArticles)
	s.HandleFunc("/", handleArticles).Name("articles")
	s.HandleFunc("/articles/{id:[0-9]+}/{slug}", handleArticle).Name("article")

	// http.Error(w, "An internal error occurred", http.StatusInternalServerError)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func handleEnglishIndex(w http.ResponseWriter, r *http.Request) {
	// lang := getLanguage(r)
	// if lang != "" {
	// 	url := mustGetURL("articles", lang)
	// 	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
	// }
	canonicalURL := mustGetURL("index_en", "")

	title := SiteName + " - Knowledge base about Obsessive Compulsive Disorder (OCD)"

	p := getPage(w, r, canonicalURL, title, "", nil)
	p.Meta.DisableHeaderLinks = true
	RenderTemplate(w, "index_en", p)
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL := mustGetURL("articles", lang)

	title := SiteName + " - " + Translate(lang, "Articles_about_Obsessive_Compusive_Disorder")
	description := ""

	p := getPage(w, r, canonicalURL, title, description, nil)
	all, err := getArticles(lang)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	sorted := make([]*article, len(all))
	i := 0
	for _, article := range all {
		sorted[i] = article
		i++
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Article.PublishedDate.After(sorted[j].Article.PublishedDate) })

	summaries := make([]articleSummary, len(all))
	for i, article := range sorted {
		summaries[i] = articleSummary{
			Title:     article.Title,
			HTMLShort: article.HTMLShort,
			URL:       mustGetArticleURL(article),
		}
		i++
	}

	p.Body = summaries
	RenderTemplate(w, "articles", p)
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	canonicalURL := mustGetURL("articles", lang)
	article, err := getArticle(lang, id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	title := article.Title + " - " + SiteName
	description := ""
	p := getPage(w, r, canonicalURL, title, description, article.ImageURL)
	p.Body = article
	RenderTemplate(w, "article", p)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL := mustGetURL("about", lang)
	title := Translate(lang, "About") + " - " + SiteName

	p := getPage(w, r, canonicalURL, title, "", nil)
	RenderTemplate(w, "about", p)
}

func getPage(w http.ResponseWriter, r *http.Request, canonicalURL *url.URL, title string, description string, socialImageURL *url.URL) *page {
	if socialImageURL == nil {
		var err error
		socialImageURL, err = router.Get("static").URL()
		if err != nil {
			panic(err)
		}
		socialImageURL.Path += "images/logo_social.png"
	}
	lang := getLang(r)

	socialURL := *canonicalURL
	if socialURL.Host[:9] == "localhost" {
		socialURL.Host = "planetocd.org"
	}

	return &page{
		Constants: Constants,
		Meta: &pageMeta{
			Lang:                  lang,
			Title:                 title,
			Description:           description,
			CanonicalURL:          canonicalURL.String(),
			SocialURL:             socialURL.String(),
			RootURL:               getRootURL(lang).String(),
			SocialImageURL:        socialImageURL.String(),
			EnableGoogleAnalytics: !isLocalEnvironment,
		},
	}
}

func getLang(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["language"]
}
