package check

import (
	"ReconDB/middlewares"
	"regexp"
)

// WildCardRegex used for validating wildcard inputs with WildCardPattern Regex.
func WildCardRegex(query string) bool {
	regex := regexp.MustCompile(middlewares.WildCardPattern)
	if regex.MatchString(query) {
		return true
	}
	return false
}
