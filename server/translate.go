package server

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

//go:embed translations
var translationsFS embed.FS
var translations = make(map[string]map[string]string)

// Translate ...
func Translate(lang string, key string) string {
	return translations[lang][key]
}

func loadTranslations(lang string) map[string]string {
	var trans map[string]string
	yamlFile, err := translationsFS.ReadFile("translations/" + lang + ".yaml")
	if err != nil {
		panic(fmt.Sprintf("Error opening translations file: %v ", err))
	}
	err = yaml.Unmarshal(yamlFile, &trans)
	if err != nil {
		panic(fmt.Sprintf("Error loading translations: %v ", err))
	}

	return trans
}

func init() {
	for _, lang := range SupportedLanguages {
		translations[lang] = loadTranslations(lang)
	}
}
