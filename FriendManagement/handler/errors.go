package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type ResponedError struct {
	StatusCode int    `json:"statuscode`
	Message    string `json:"message"`
}

var (
	ErrorMethodNotAllowed = &ResponedError{StatusCode: 405, Message: "Method not allowed"}
	ErrorNotFound         = &ResponedError{StatusCode: 404, Message: "Resource not found"}
	ErrorBadRequest       = &ResponedError{StatusCode: 400, Message: "Bad Request"}
)

func (e *ResponedError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ServerErrorRenderer(err error) *ResponedError {
	return &ResponedError{
		StatusCode: 500,
		Message:    err.Error(),
	}
}

func ServerSuccessRenderer(err error) *ResponedError {
	return &ResponedError{
		StatusCode: 201,
		Message:    err.Error(),
	}
}
