package handler

import (
	"github.com/go-chi/render"
	"net/http"

	_ "github.com/go-chi/render"
)

type ResponseError struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
}

var (
	ErrorMethodNotAllowed = &ResponseError{StatusCode: 405, Message: "Method not allowed"}
	ErrorNotFound         = &ResponseError{StatusCode: 404, Message: "Resource not found"}
	ErrorBadRequest       = &ResponseError{StatusCode: 400, Message: "Bad Request"}
)

func (e *ResponseError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ServerErrorRenderer(err error) *ResponseError {
	return &ResponseError{
		StatusCode: 500,
		Message:    err.Error(),
	}
}

func ServerSuccessRenderer(err error) *ResponseError {
	return &ResponseError{
		StatusCode: 201,
		Message:    err.Error(),
	}
}
