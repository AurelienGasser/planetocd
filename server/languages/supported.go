package languages

import "golang.org/x/text/language"

var SupportedLanguages []string
var LangBaseToLang map[language.Base]string
var LanguageMatcher language.Matcher

func init() {
	frBase, _ := language.French.Base()
	esBase, _ := language.Spanish.Base()
	// zhBase, _ := language.Chinese.Base()

	SupportedLanguages = []string{
		"fr",
		"es",
		// "zh",
	}

	LangBaseToLang = map[language.Base]string{
		frBase: "fr",
		esBase: "es",
		// zhBase: "zh",
	}

	matcherLanguages := [...]language.Tag{
		language.English,
		language.French,
		language.Spanish,
		// language.Chinese,
	}

	LanguageMatcher = language.NewMatcher(matcherLanguages[:])
}
