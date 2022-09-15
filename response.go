package main

import (
	"encoding/json"
	"net/http"
)

type Response struct{}
type ResponseParams struct {
	w      http.ResponseWriter
	result any
}

func (response *Response) Ok(p ResponseParams) {
	response.Write(p.w, 200, p.result)
}

func (response *Response) Created(p ResponseParams) {
	response.Write(p.w, 201, p.result)
}

func (response *Response) NotFound(p ResponseParams) {
	response.Write(p.w, 404, p.result)
}

func (response *Response) Write(w http.ResponseWriter, statusCode int, result any) {
	w.WriteHeader(statusCode)

	if result != nil {
		json.NewEncoder(w).Encode(result)
	}
}
