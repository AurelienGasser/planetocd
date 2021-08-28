package server

import (
	"html/template"
	"net/url"
	"regexp"
	"strings"
)

type page struct {
	Constants map[string]interface{}
	Body      interface{}
	Meta      *pageMeta
}

type pageMeta struct {
	TemplateName          string
	Lang                  string
	Description           string
	CanonicalURL          string
	SocialURL             string
	Title                 string
	RootURL               string
	SocialImageURL        string
	EnableGoogleAnalytics bool
	DisableHeaderLinks    bool
}

// T translates an input key using the Page's lang code
func (p *page) T(key string) string {
	return Translate(p.Meta.Lang, key)
}

// URL adds the language prefix to an URL path
func (p *page) URL(path string) string {
	return "/" + p.Meta.Lang + path
}

// MustGetURL returns the URL for a name, and panics if not found
func (p *page) MustGetURL(name string) *url.URL {
	return mustGetURL(name, p.Meta.Lang)
}

func (p *page) ReplaceURL(s, old, url string) template.HTML {
	return template.HTML(strings.Replace(s, old, "<a href=\""+template.HTMLEscapeString(url)+"\">"+old+"</a>", 1))
}

func (p *page) ReplaceURLTemplate(s template.HTML, old, url string) template.HTML {
	return p.ReplaceURL(string(s), old, url)
}

func (p *page) ReplaceEmail(s string) template.HTML {
	re := regexp.MustCompile(`(.*)\[(.*)\]\(#email#\)(.*)`)
	return template.HTML(re.ReplaceAll([]byte(s), []byte("$1<a href=\"mailto:"+template.HTMLEscapeString(Email)+"\">$2</a>$3")))
}
