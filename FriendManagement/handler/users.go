package handler

import (
	"Assignment/models"
	"Assignment/service"
	"github.com/go-chi/render"
	"net/http"
)

func createUser(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.FListRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrorBadRequest)
			return
		}
		response, err := service.CreateUser(req)

	}
}
