package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

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
	templates["article"] = loadTemplate("article.html")
	templates["about_fr"] = loadTemplate("about_fr.html")
	templates["about_es"] = loadTemplate("about_es.html")
	templates["about_zh"] = loadTemplate("about_zh.html")
	return templates
}

func loadTemplate(filename string) *template.Template {
	templatesPath := "server/templates"
	partialsPath := "server/templates/partials"

	return template.Must(template.Must(
		template.ParseGlob(partialsPath + "/*")).
		ParseFiles(path.Join(templatesPath, "/"+filename)))
}
