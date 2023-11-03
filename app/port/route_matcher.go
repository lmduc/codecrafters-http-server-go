package port

type RouteMatcher interface {
	Match(Request) bool
}
