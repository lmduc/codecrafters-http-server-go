package response

type TextPlainResponse struct {
	*Response
}

func NewTextPlainResponse() *TextPlainResponse {
	return &TextPlainResponse{NewResponse("text/plain")}
}
