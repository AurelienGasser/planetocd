package server

import (
	"net/url"
	"strconv"
)

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

func mustGetLikeArticleURL(lang string, id int, slug string) *url.URL {
	res, err := router.Get("likeArticle").URL(
		"language", lang,
		"id", strconv.Itoa(id),
		"slug", slug)
	if err != nil {
		panic(err)
	}
	return res
}

func mustGetUpdateArticleLikeURL(lang string, articleID int, slug string, likeID int) *url.URL {
	res, err := router.Get("updateArticleLike").URL(
		"language", lang,
		"id", strconv.Itoa(articleID),
		"likeID", strconv.Itoa(likeID),
		"slug", slug)
	if err != nil {
		panic(err)
	}
	return res
}

func mustGetTagURL(lang string, tag string) *url.URL {
	res, err := router.Get("tag").URL(
		"language", lang,
		"tag", tag)
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
