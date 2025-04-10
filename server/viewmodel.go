package server

import (
	"fmt"
	"html/template"
	"net/url"
	"regexp"
	"strings"
)

type ViewModel struct {
	Constants map[string]interface{}
	Meta      *ViewModelMeta
	Body      interface{}
}

type ViewModelMeta struct {
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
	EnablePetitionBanner  bool
}

func NewEmptyViewModel(constants map[string]interface{}, meta *ViewModelMeta) *ViewModel {
	return &ViewModel{
		Constants: constants,
		Meta:      meta,
	}
}

func NewViewModel(constants map[string]interface{}, meta *ViewModelMeta, body interface{}) *ViewModel {
	return &ViewModel{
		Constants: constants,
		Meta:      meta,
		Body:      body,
	}
}

func (p *ViewModel) NewViewModel(body interface{}) *ViewModel {
	return NewViewModel(p.Constants, p.Meta, body)
}

// T translates an input key using the Page's lang code
func (p *ViewModel) T(key string) string {
	return Translate(p.Meta.Lang, key)
}

// T translates an input tag using the Page's lang code
func (p *ViewModel) TranslateTag(tag string) string {
	return TranslateTag(p.Meta.Lang, tag)
}

// URL adds the language prefix to an URL path
func (p *ViewModel) URL(path string) string {
	return "/" + p.Meta.Lang + path
}

// MustGetURL returns the URL for a name, and panics if not found
func (p *ViewModel) MustGetURL(name string) *url.URL {
	return mustGetURL(name, p.Meta.Lang)
}

func (p *ViewModel) Replace(s template.HTML, old, new string, n int) template.HTML {
	return template.HTML(strings.Replace(string(s), old, new, n))
}

func (p *ViewModel) ReplaceEmail(s string) template.HTML {
	re := regexp.MustCompile(`(.*)\[(.*)\]\(#email#\)(.*)`)
	return template.HTML(re.ReplaceAll([]byte(s), []byte("$1<a href=\"mailto:"+template.HTMLEscapeString(Email)+"\">$2</a>$3")))
}

func (p *ViewModel) ReplaceURLPattern(s string, needle string, url string) template.HTML {
	re := regexp.MustCompile(`(.*)\[([^\]]*)\]\(#` + needle + `#\)(.*)`)
	return template.HTML(re.ReplaceAll([]byte(s), []byte("$1<a href=\""+template.HTMLEscapeString(url)+"\">$2</a>$3")))
}

func (p *ViewModel) ReplaceURLPatternTemplate(s template.HTML, needle string, url string) template.HTML {
	return p.ReplaceURLPattern(string(s), needle, url)
}

func (p *ViewModel) Tag(tag string) template.HTML {
	return template.HTML(fmt.Sprintf("<a href=\"%v\"><span class=\"uk-label uk-label-success\">%v</span></a>", p.TagURL(tag), p.TranslateTag(tag)))
}

func (p *ViewModel) TagURL(tag string) *url.URL {
	return mustGetTagURL(p.Meta.Lang, tag)
}

func (p *ViewModel) UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}
