package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

//go:embed all:templates
var templateFS embed.FS
var templates = loadTemplates()

// RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, p *ViewModel) {
	err := templates[p.Meta.TemplateName].ExecuteTemplate(w, "layout", p)
	if err != nil {
		fmt.Printf("Error serving page %v\n", err)
	}
}

func loadTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)
	templates["index_en"] = loadTemplate("index_en.html")
	templates["articles"] = loadTemplate("articles.html")
	templates["tag"] = loadTemplate("tag.html")
	templates["article"] = loadTemplate("article.html")
	templates["about_fr"] = loadTemplate("about_fr.html")
	templates["about_es"] = loadTemplate("about_es.html")
	templates["about_zh"] = loadTemplate("about_zh.html")
	return templates
}

var defaultfuncs = map[string]interface{}{}

func loadTemplate(filename string) *template.Template {
	templatesPath := "templates/"
	partialsPath := "templates/partials/*"
	tpl := template.New(filename).Funcs(defaultfuncs)
	partialsTpl := template.Must(tpl.ParseFS(templateFS, partialsPath))
	return template.Must(partialsTpl.ParseFS(templateFS, path.Join(templatesPath, filename)))
}
