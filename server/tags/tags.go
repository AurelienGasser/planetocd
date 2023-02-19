package tags

import (
	"github.com/aureliengasser/planetocd/articles"
	"github.com/aureliengasser/planetocd/server/languages"
)

var Tags map[string][]string
var Articles map[string]map[string][]*articles.Article

func TagExists(lang string, tag string) bool {
	_, ok := Articles[lang][tag]
	return ok
}

func init() {
	if Articles != nil {
		return
	}

	Tags = make(map[string][]string, len(languages.SupportedLanguages))
	Articles = make(map[string]map[string][]*articles.Article, len(languages.SupportedLanguages))

	for lang, articlesByID := range articles.GetArticles() {
		Articles[lang] = make(map[string][]*articles.Article)
		for _, article := range articlesByID {
			for _, tag := range article.Tags {
				if _, ok := Articles[lang][tag]; !ok {
					Articles[lang][tag] = make([]*articles.Article, 0)
				}
				Articles[lang][tag] = append(Articles[lang][tag], article)
			}
		}
		Tags[lang] = make([]string, len(Articles[lang]))
		i := 0
		for tag := range Articles[lang] {
			Tags[lang][i] = tag
			i++
		}
	}
}
