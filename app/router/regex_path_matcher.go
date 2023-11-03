package router

import (
	"regexp"

	"github.com/codecrafters-io/http-server-starter-go/app/port"
)

type RegexPathMatcher struct {
	verb string
	r    *regexp.Regexp
}

func (r *RegexPathMatcher) Match(req port.Request) bool {
	return r.r.MatchString(req.Path())
}

func (r *RegexPathMatcher) FindMatches(path string) []string {
	return r.r.FindStringSubmatch(path)
}

func NewRegexPathMatcher(verb string, re string) *RegexPathMatcher {
	return &RegexPathMatcher{verb, regexp.MustCompile(re)}
}
