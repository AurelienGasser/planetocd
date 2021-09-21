package server

import (
	"fmt"
	"html/template"
	"net/url"

	"github.com/aureliengasser/planetocd/articles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

type article struct {
	*articles.Article
	Slug     string
	URL      *url.URL
	ImageURL *url.URL
	Pages    []articlePage
}

type articlePage struct {
	PageNumber int
	HTML       template.HTML
	HTMLShort  template.HTML
	URL        *url.URL
}

func newArticle(a *articles.Article) *article {
	slug := Slugify(a.Title)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	pages := make([]articlePage, len(a.MarkdownPages))

	for i, p := range a.MarkdownPages {
		htmlBytes := markdown.ToHTML([]byte(p), nil, renderer)
		html := string(htmlBytes)
		htmlShort := getHTMLShort(html)
		ap := articlePage{
			PageNumber: i + 1,
			HTML:       template.HTML(html),
			HTMLShort:  template.HTML(htmlShort),
			URL:        mustGetArticlePageURL(a.Lang, a.ID, slug, i+1),
		}
		pages[i] = ap
	}

	res := &article{
		Article: a,
		Pages:   pages,
		Slug:    slug,
		URL:     mustGetArticleURL(a.Lang, a.ID, slug),
	}

	if a.Image != "" {
		staticURL, err := router.Get("static").URL()
		if err != nil {
			panic(err)
		}
		staticURL.Path += fmt.Sprintf("/images/illustrations/%v", a.Image)
		res.ImageURL = staticURL
	}
	return res
}

func getHTMLShort(HTML string) string {
	endTag := "</p>"
	length := len(HTML)
	i := 500
	for ; i+len(endTag) < length; i++ {
		if HTML[i:i+len(endTag)] == endTag {
			break
		}
	}
	return HTML[:i+len(endTag)]
}
