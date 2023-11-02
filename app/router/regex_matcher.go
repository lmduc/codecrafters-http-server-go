package router

import (
	"fmt"
	"regexp"
)

type RegexMatcher struct {
	r *regexp.Regexp
}

func (r *RegexMatcher) Match(path string) bool {
	return r.r.MatchString(path)
}

func (r *RegexMatcher) FindMatch(path string) string {
	fmt.Println("Inside regex: ", path)
	return r.r.FindString(path)
}

func NewRegexMatcher(re string) *RegexMatcher {
	fmt.Println("re is: ", re)
	return &RegexMatcher{regexp.MustCompile(re)}
}
