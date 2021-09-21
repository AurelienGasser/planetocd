package server

import (
	"html/template"
	"net/url"
)

type articleSummary struct {
	URL       *url.URL
	Title     string
	HTMLShort template.HTML
}
