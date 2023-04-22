package check

import (
	"ReconDB/middlewares"
	"regexp"
)

func WildCardRegex(query string) bool {
	regex := regexp.MustCompile(middlewares.WildCardPattern)
	if regex.MatchString(query) {
		return true
	}
	return false
}
