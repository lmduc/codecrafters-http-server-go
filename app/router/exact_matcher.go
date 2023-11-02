package router

type ExactMatcher struct {
	path string
}

func (e *ExactMatcher) Match(path string) bool {
	return e.path == path
}

func NewExactMatcher(path string) *ExactMatcher {
	return &ExactMatcher{path}
}
