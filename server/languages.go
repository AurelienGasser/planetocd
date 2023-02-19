package server

import (
	"net/http"

	"github.com/aureliengasser/planetocd/server/languages"
	"golang.org/x/text/language"
)

func getLanguage(r *http.Request) string {
	acceptLanguage := r.Header["Accept-Language"]
	if len(acceptLanguage) == 0 {
		return ""
	}
	for _, lang := range acceptLanguage {
		tag, _ := language.MatchStrings(languages.LanguageMatcher, lang)
		base, _ := tag.Base()
		return languages.LangBaseToLang[base]
	}
	return ""
}
