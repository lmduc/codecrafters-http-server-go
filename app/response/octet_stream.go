package response

type OctetStream struct {
	*Response
}

func NewOctetStream() *OctetStream {
	return &OctetStream{NewResponse("application/octet-stream")}
}
