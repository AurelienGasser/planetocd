package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

//go:embed static
var staticFS embed.FS
var router *mux.Router
var isLocalEnvironment bool

// Listen ...
func Listen(port int, isLocal bool) {
	isLocalEnvironment = isLocal

	router = mux.NewRouter().
		Schemes("http", "https").
		Host(getHost(isLocal, port)).
		Subrouter()

	registerRoutes(router, port)
}

func getHost(isLocal bool, port int) string {
	host, ok := os.LookupEnv("PLANETOCD_HOST")
	if ok {
		return host
	}

	if isLocal {
		return fmt.Sprintf("localhost:%v", port)
	}
	return Host
}

func registerRoutes(router *mux.Router, port int) {
	router.Path("/").HandlerFunc(handleEnglishIndex).Name("index_en")
	router.Path("/robots.txt").HandlerFunc(handleRobots)
	router.Path("/sitemap.xml").HandlerFunc(handleSitemap).Name("sitemap")
	router.Path("/articles.json").HandlerFunc(handleArticlesJson).Name("articlesJson")
	router.PathPrefix("/static/").Handler(http.FileServer(http.FS(staticFS))).Name("static")

	s := router.PathPrefix("/{language}").Subrouter()
	s.HandleFunc("/tags", handleTags).Name("tags")
	s.HandleFunc("/about", handleAbout).Name("about")
	s.HandleFunc("/tag/{tag:[a-z-]+}", handleTag)
	s.HandleFunc("/tag/{tag:[a-z-]+}/", handleTag).Name("tag")
	s.HandleFunc("", handleArticles)
	s.HandleFunc("/", handleArticles).Name("articles")
	s.HandleFunc("/articles/{id:[0-9]+}/{slug}", handleArticle).Name("article")

	// Load articles
	ensureLoaded()

	// http.Error(w, "An internal error occurred", http.StatusInternalServerError)
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
	log.Println("Server started")
}

func getArticlesCanonicalURL(baseURL *url.URL, pageNumber int) *url.URL {
	res := *baseURL
	if pageNumber != 1 {
		res.RawQuery = fmt.Sprintf("page=%v", pageNumber)
	}
	return &res
}

func getPageNumber(query url.Values, maxPageNumber int) (int, bool) {
	pageNumberStr := query.Get("page")
	if pageNumberStr != "" {
		tmp, err := strconv.Atoi(pageNumberStr)
		if err == nil && tmp >= 1 && tmp <= maxPageNumber {
			return tmp, true
		} else {
			return 0, false
		}
	}
	return 1, true
}

func getViewModel(tmpl string, r *http.Request, canonicalURL *url.URL, title string, description string, socialImageURL *url.URL) *ViewModel {
	if socialImageURL == nil {
		var err error
		socialImageURL, err = router.Get("static").URL()
		if err != nil {
			panic(err)
		}
		socialImageURL.Path += "images/logo_social.webp"
	}
	lang := getLang(r)

	socialURL := *canonicalURL
	if socialURL.Host[:9] == "localhost" {
		socialURL.Host = "planetocd.org"
	}

	// _, err := r.Cookie(DismissBannerCookieName)
	// hasDismissBannerCookie := err == nil
	meta := &ViewModelMeta{
		TemplateName:          tmpl,
		Lang:                  lang,
		Title:                 title,
		Description:           description,
		CanonicalURL:          canonicalURL.String(),
		SocialURL:             socialURL.String(),
		RootURL:               getRootURL(lang).String(),
		SocialImageURL:        socialImageURL.String(),
		EnableGoogleAnalytics: !isLocalEnvironment,
		EnablePetitionBanner:  false,
	}

	return NewEmptyViewModel(Constants, meta)
}

func getLang(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["language"]
}
