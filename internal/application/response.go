package application

import (
	"encoding/json"
	"net/http"
)

type Response struct{}
type ResponseParams struct {
	W      http.ResponseWriter
	Result any
}

func (response *Response) Ok(p ResponseParams) {
	response.Write(p.W, 200, p.Result)
}

func (response *Response) Created(p ResponseParams) {
	response.Write(p.W, 201, p.Result)
}

func (response *Response) NotFound(p ResponseParams) {
	response.Write(p.W, 404, p.Result)
}

func (response *Response) BadRequest(p ResponseParams) {
	response.Write(p.W, 403, p.Result)
}

func (response *Response) ServerError(p ResponseParams) {
	response.Write(p.W, 500, p.Result)
}

func (response *Response) Write(w http.ResponseWriter, statusCode int, result any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if result != nil {
		json.NewEncoder(w).Encode(result)
	}
}
