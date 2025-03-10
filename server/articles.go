package server

import (
	"fmt"
	"html/template"
	"math"
	"sort"

	"github.com/aureliengasser/planetocd/articles"
	"github.com/aureliengasser/planetocd/server/cache"
	"github.com/aureliengasser/planetocd/server/languages"
	"github.com/aureliengasser/planetocd/server/tags"
	"github.com/aureliengasser/planetocd/server/urls"
	"github.com/aureliengasser/planetocd/server/viewModel"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

var articlesPageSize int = 4
var allArticles map[string]map[int]*cache.Article
var allArticlesPaginated map[string]*cache.Articles
var allArticlesPaginatedByTag map[string]map[string]*cache.Articles
var articlesMap map[string]map[int]*cache.Article

func getArticles(lang string) (map[int]*cache.Article, error) {
	ensureLoaded()
	res, ok := allArticles[lang]
	if !ok {
		return nil, fmt.Errorf("unknown lang \"%v\"", lang)
	}
	return res, nil
}

func getArticle(lang string, id int) (*cache.Article, error) {
	byLang, err := getArticles(lang)
	if err != nil {
		return nil, err
	}
	article, ok := byLang[id]
	if !ok {
		return nil, fmt.Errorf("unknown article id %v for lang %v", id, lang)
	}
	return article, nil
}

func ensureLoaded() {
	if allArticles != nil {
		return
	}

	articlesMap = getArticlesMap()
	allArticles = getAllArticles()
	allArticlesPaginated = getAllArticlesPaginated()

	loadAllArticlesPaginatedByTag()
}

func getArticlesMap() map[string]map[int]*cache.Article {
	src := articles.GetArticles()
	res := make(map[string]map[int]*cache.Article, 0)
	for _, lang := range languages.SupportedLanguages {
		res[lang] = make(map[int]*cache.Article)
		for id, article := range src[lang] {
			res[lang][id] = newArticle(article)
		}
	}
	return res
}

func getAllArticles() map[string]map[int]*cache.Article {
	res := make(map[string]map[int]*cache.Article)
	all := articles.GetArticles()
	for lang, byLang := range all {
		res[lang] = make(map[int]*cache.Article)
		for id, article := range byLang {
			res[lang][id] = newArticle(article)
		}
	}
	return res
}

func loadAllArticlesPaginatedByTag() {
	allArticlesPaginatedByTag = make(map[string]map[string]*cache.Articles)

	for lang, articlesByTag := range tags.Articles {
		allArticlesPaginatedByTag[lang] = make(map[string]*cache.Articles)
		for tag, src := range articlesByTag {
			sorted := make([]*articles.Article, len(src))
			copy(sorted, src)
			sortArticles(sorted)
			numPages := int(math.Ceil(float64(len(sorted)) / float64(articlesPageSize)))
			pages := make([]*cache.ArticlesPage, numPages)
			for i := 0; i < numPages; i++ {
				page := sorted[i*articlesPageSize : int(math.Min(float64(len(sorted)), float64((i+1)*articlesPageSize)))]

				pages[i] = &cache.ArticlesPage{
					PageNumber: i + 1,
					Articles:   convertSlice(lang, page),
				}
			}

			allArticlesPaginatedByTag[lang][tag] = cache.NewArticles(lang, pages)
		}
	}
}

func getAllArticlesPaginated() map[string]*cache.Articles {
	res := make(map[string]*cache.Articles)

	for lang, articles := range allArticles {
		sorted := getSortedArticles(articles)
		numPages := int(math.Ceil(float64(len(sorted)) / float64(articlesPageSize)))
		pages := make([]*cache.ArticlesPage, numPages)

		for i := 0; i < numPages; i++ {
			pages[i] = &cache.ArticlesPage{
				PageNumber: i + 1,
				Articles:   sorted[i*articlesPageSize : int(math.Min(float64(len(sorted)), float64((i+1)*articlesPageSize)))],
			}
		}

		res[lang] = &cache.Articles{
			Lang:  lang,
			Pages: pages,
		}
	}

	return res
}

// getSortedArticles gets a list of articles sorted by descending published date
func getSortedArticles(all map[int]*cache.Article) []*cache.Article {
	res := make([]*cache.Article, len(all))
	i := 0
	for _, article := range all {
		res[i] = article
		i++
	}
	sortCacheArticles(res)
	return res
}

func convertSlice(lang string, src []*articles.Article) []*cache.Article {
	res := make([]*cache.Article, len(src))
	for i := range src {
		res[i] = articlesMap[lang][src[i].ID]
	}
	return res
}

func sortArticles(a []*articles.Article) {
	sort.Slice(a, func(i, j int) bool {
		return a[i].PublishedDate.After(a[j].PublishedDate)
	})
}

func sortCacheArticles(a []*cache.Article) {
	sort.Slice(a, func(i, j int) bool {
		return a[i].PublishedDate.After(a[j].PublishedDate)
	})
}

func newArticle(a *articles.Article) *cache.Article {
	slug := urls.Slugify(a.Title)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	pages := make([]*cache.ArticlePage, len(a.MarkdownPages))

	for i, p := range a.MarkdownPages {
		htmlBytes := markdown.ToHTML([]byte(p), nil, renderer)
		html := string(htmlBytes)
		htmlShort := getHTMLShort(html)
		ap := cache.ArticlePage{
			PageNumber: i + 1,
			HTML:       template.HTML(html),
			HTMLShort:  template.HTML(htmlShort),
			URL:        mustGetArticlePageURL(a.Lang, a.ID, slug, i+1),
		}
		pages[i] = &ap
	}

	translators := make([]string, 0)
	hasHumanTranslator := false
	for _, translator := range a.Translators {
		if !translator.IsBot {
			translators = append(translators, translator.Name)
			hasHumanTranslator = true
		}
	}

	res := &cache.Article{
		Article:            a,
		Pages:              pages,
		Translators:        translators,
		HasHumanTranslator: hasHumanTranslator,
		Slug:               slug,
		Tags:               a.Tags,
		URL:                mustGetArticleURL(a.Lang, a.ID, slug),
	}

	if a.Image != "" {
		staticURL, err := router.Get("static").URL()
		if err != nil {
			panic(err)
		}
		staticURL.Path += fmt.Sprintf("images/illustrations/%v", a.Image)
		res.Illustration = viewModel.NewArticleIllustration(staticURL)
	}
	return res
}

func getHTMLShort(HTML string) string {
	endTags := []string{"</p>", "<br/>", "<br />", "</li>", "<br>", "</ul>", "</ol>", "</blockquote>"}
	length := len(HTML)
	var i int
	for _, endTag := range endTags {
		i = 500
		for ; i+len(endTag) < length; i++ {
			if HTML[i:i+len(endTag)] == endTag {
				return HTML[:i+len(endTag)]
			}
		}
	}
	return HTML
}
