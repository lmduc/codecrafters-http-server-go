package port

type PathMatcher interface {
	Match(string) bool
}
