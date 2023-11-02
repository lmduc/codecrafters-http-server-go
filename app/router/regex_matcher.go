package router

import "regexp"

type RegexMatcher struct {
	r *regexp.Regexp
}

func (r *RegexMatcher) Match(path string) bool {
	return r.r.MatchString(path)
}

func NewRegexMatcher(re string) *RegexMatcher {
	return &RegexMatcher{regexp.MustCompile(re)}
}
