package port

type Router interface {
	Handle(Request) error
}
