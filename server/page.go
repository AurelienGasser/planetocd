package server

import (
	"html/template"
	"net/url"
	"regexp"
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

func (p *page) ReplaceEmail(s string) template.HTML {
	re := regexp.MustCompile(`(.*)\[(.*)\]\(#email#\)(.*)`)
	return template.HTML(re.ReplaceAll([]byte(s), []byte("$1<a href=\"mailto:"+template.HTMLEscapeString(Email)+"\">$2</a>$3")))
}

func (p *page) ReplaceURLPattern(s string, needle string, url string) template.HTML {
	re := regexp.MustCompile(`(.*)\[([^\]]*)\]\(#` + needle + `#\)(.*)`)
	return template.HTML(re.ReplaceAll([]byte(s), []byte("$1<a href=\""+template.HTMLEscapeString(url)+"\">$2</a>$3")))
}

func (p *page) ReplaceURLPatternTemplate(s template.HTML, needle string, url string) template.HTML {
	return p.ReplaceURLPattern(string(s), needle, url)
}
