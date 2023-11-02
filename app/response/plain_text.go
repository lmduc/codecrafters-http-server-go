package response

type PlainTextResponse struct {
	*Response
}

func NewPlainTextResponse() *PlainTextResponse {
	return &PlainTextResponse{NewResponse("plain/text")}
}
