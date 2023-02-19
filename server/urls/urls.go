package urls

import (
	"regexp"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

// Slugify ...
func Slugify(s string) string {
	s = unidecode.Unidecode(s)
	s = strings.ToLower(s)
	var re = regexp.MustCompile("[^a-z0-9-_]+")
	s = re.ReplaceAllLiteralString(s, "-")
	re = regexp.MustCompile("-{2,}")
	s = re.ReplaceAllLiteralString(s, "-")
	re = regexp.MustCompile("(-|-)$")
	s = re.ReplaceAllLiteralString(s, "")
	return s
}
