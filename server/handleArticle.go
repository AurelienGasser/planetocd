package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aureliengasser/planetocd/server/likes"
	"github.com/aureliengasser/planetocd/server/viewModel"
	"github.com/gorilla/mux"
)

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

	likeURL := mustGetLikeArticleURL(lang, id, article.Slug)
	updateLikeURL := mustGetUpdateArticleLikeURL(lang, id, article.Slug, -1)

	vm := getViewModel("article", r, canonicalURL, title, description, article.Illustration.Md())
	suggestions, err := getArticleSuggestions(article)

	if err != nil {
		log.Printf("Error getting suggestions for article %v in lang %v: %v\n", article.ID, lang, err)
	}

	likes, err := likes.Get(id)
	if err != nil {
		log.Printf("Error getting likes for article %v in lang %v: %v\n", article.ID, lang, err)
	}

	articleVM := articleViewModel{
		Article:          article,
		CurrentPageIndex: pageNumber - 1,
		Pagination:       viewModel.GetPagination(article, pageNumber),
		Suggestions:      suggestions,
		LikeURL:          likeURL,
		UpdateLikeURL:    updateLikeURL,
		TestURL:          mustGetURL("articles", lang),
		Likes:            likes,
	}
	vm.Body = &articleVM
	RenderTemplate(w, vm)
}
