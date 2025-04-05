package server

import (
	"net/http"
	"sort"
	"strings"

	"github.com/aureliengasser/planetocd/server/tags"
)

func NewTag(tag_ string, articles *Articles) *tag {
	return &tag{
		Tag:      tag_,
		Articles: articles,
	}
}

type tag struct {
	Tag      string
	Articles *Articles
}

func handleTags(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL := mustGetURL("tags", lang)
	title := Translate(lang, "Tags") + " - " + SiteName

	tags := tags.GetAllTags(lang, allArticles)
	sort.Slice(tags, func(i, j int) bool {
		return strings.ToLower(TranslateTag(lang, tags[i])) < strings.ToLower(TranslateTag(lang, tags[j]))
	})

	vm := getViewModel("tags", r, canonicalURL, title, "", nil)
	vm.Body = tags

	RenderTemplate(w, vm)
}
