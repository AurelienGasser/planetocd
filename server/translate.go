package server

import (
	"embed"
	"fmt"

	"github.com/aureliengasser/planetocd/server/languages"
	"gopkg.in/yaml.v2"
)

//go:embed translations
var translationsFS embed.FS
var translations = make(map[string]map[string]string)

// Translate ...
func Translate(lang string, key string, amounts ...int) string {
	if len(amounts) == 0 || amounts[0] <= 1 {
		return translations[lang][key]
	}
	res := translations[lang][fmt.Sprintf("%v (plural)", key)]
	if res != "" {
		return res
	}
	return translations[lang][key]
}

// TranslateTag ...
func TranslateTag(lang string, tag string) string {
	return translations[lang][fmt.Sprintf("tag_%v", tag)]
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
	for _, lang := range languages.SupportedLanguages {
		translations[lang] = loadTranslations(lang)
	}
}
