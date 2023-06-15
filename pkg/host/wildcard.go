package host

import (
	"ReconDB/pkg"
	"regexp"
	"strings"
)

// WildCardRegex used for validating wildcard inputs with WildCardPattern Regex.
func WildCardRegex(query string) bool {
	regex := regexp.MustCompile(pkg.WildCardPattern)
	if regex.MatchString(query) {
		return true
	}
	return false
}

func MatchWildcard(pattern, text string) bool {
	pattern = "^" + pattern + "$"
	pattern = strings.ReplaceAll(pattern, "*", ".*")

	matched, _ := regexp.MatchString(pattern, text)
	return matched
}
