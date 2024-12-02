package server

import (
	"fmt"
	"net/http"
)

func handleAbout(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL := mustGetURL("about", lang)
	title := Translate(lang, "About") + " - " + SiteName

	p := getViewModel(fmt.Sprintf("about_%v", lang), r, canonicalURL, title, "", nil)
	RenderTemplate(w, p)
}
