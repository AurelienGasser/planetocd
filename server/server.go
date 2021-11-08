package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"

	"github.com/aureliengasser/planetocd/server/viewModel"
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
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
	log.Println("Server started")
}

func handleEnglishIndex(w http.ResponseWriter, r *http.Request) {
	// lang := getLanguage(r)
	// if lang != "" {
	// 	url := mustGetURL("articles", lang)
	// 	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
	// }
	canonicalURL := mustGetURL("index_en", "")

	title := SiteName + " - Knowledge base about Obsessive Compulsive Disorder (OCD)"

	p := getViewModel("index_en", r, canonicalURL, title, "", nil)
	p.Meta.DisableHeaderLinks = true
	RenderTemplate(w, p)
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL := mustGetURL("articles", lang)

	title := SiteName + " - " + Translate(lang, "Articles_about_Obsessive_Compusive_Disorder")
	description := ""

	p := getViewModel("articles", r, canonicalURL, title, description, nil)

	all, err := getArticles(lang)
	if err != nil {
		http.NotFound(w, r)
		fmt.Println(err)
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
			HTMLShort: article.Pages[0].HTMLShort,
			URL:       mustGetArticleURL(article.Lang, article.ID, article.Slug),
		}
		i++
	}

	p.Body = summaries
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

	canonicalURL, _ := router.Get("article").URL("language", lang, "id", idStr, "slug", article.Slug)

	title := article.Title + " - " + SiteName
	description := ""
	vm := getViewModel("article", r, canonicalURL, title, description, article.ImageURL)
	pageNumber := 1
	pageNumberStr := r.URL.Query().Get("page")
	if pageNumberStr != "" {
		tmp, err := strconv.Atoi(pageNumberStr)
		if err == nil && tmp >= 1 && tmp <= len(article.Pages) {
			pageNumber = tmp
		}
	}

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

func getViewModel(tmpl string, r *http.Request, canonicalURL *url.URL, title string, description string, socialImageURL *url.URL) *ViewModel {
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

	return &ViewModel{
		Constants: Constants,
		Meta: &ViewModelMeta{
			TemplateName:          tmpl,
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
