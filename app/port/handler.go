package port

type Handler interface {
	Handle(Request) (Response, error)
}
