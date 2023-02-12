package cache

import (
	"net/url"
	"regexp"
)

var regexArticleIllustration = regexp.MustCompile(`(\.[^\.]+)$`)

type ArticleIllustration struct {
	URLWithoutSize *url.URL
}

func NewArticleIllustration(urlWithoutSize *url.URL) *ArticleIllustration {
	return &ArticleIllustration{
		URLWithoutSize: urlWithoutSize,
	}
}

func (i *ArticleIllustration) Sm() *url.URL {
	if i == nil {
		return nil
	}
	return i.withSize("sm")
}

func (i *ArticleIllustration) Md() *url.URL {
	if i == nil {
		return nil
	}
	return i.withSize("md")
}

func (i *ArticleIllustration) Lg() *url.URL {
	if i == nil {
		return nil
	}
	return i.withSize("lg")
}

func (i *ArticleIllustration) withSize(suffix string) *url.URL {
	withSize := *i.URLWithoutSize
	withSize.Path = regexArticleIllustration.ReplaceAllString(withSize.Path, "_"+suffix+"$1")
	return &withSize
}
