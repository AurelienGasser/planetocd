package server

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

// Slugify ...
func Slugify(s string) string {
	s = unidecode.Unidecode(s)
	s = strings.ToLower(s)
	var re = regexp.MustCompile("[^a-z0-9-_]+")
	s = re.ReplaceAllLiteralString(s, "-")
	re = regexp.MustCompile("-{2,}")
	s = re.ReplaceAllLiteralString(s, "-")
	re = regexp.MustCompile("(-|-)$")
	s = re.ReplaceAllLiteralString(s, "")
	return s
}

func mustGetURL(name string, lang string) *url.URL {
	res, err := router.Get(name).URL("language", lang)
	if err != nil {
		panic(err)
	}
	return res
}

func mustGetArticleURL(lang string, id int, slug string) *url.URL {
	res, err := router.Get("article").URL(
		"language", lang,
		"id", strconv.Itoa(id),
		"slug", slug)
	if err != nil {
		panic(err)
	}
	return res
}

func mustGetArticlePageURL(lang string, id int, slug string, page int) *url.URL {
	res := mustGetArticleURL(lang, id, slug)
	q := res.Query()
	q.Set("page", strconv.Itoa(page))
	res.RawQuery = q.Encode()
	return res
}

func getRootURL(lang string) *url.URL {
	rootURL, err := router.Get("articles").URL("language", lang)
	if err != nil {
		rootURL = &url.URL{}
	}
	return rootURL
}
