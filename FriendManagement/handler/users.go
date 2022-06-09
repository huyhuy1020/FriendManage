package handler

import (
	"Assignment/models"
	"Assignment/service"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

func createUser(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FListRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrorBadRequest)
			return
		}
		response, err := service.CreateUser(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		//Send result response
		json.NewEncoder(w).Encode(response)
	}
}

func createFriendConnection(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FConnectionrequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrorBadRequest)
			return
		}
		response, err := service.CreateConnection(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		//Send result response
		json.NewEncoder(w).Encode(response)
	}
}

func getFriendList(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.CommonFriendRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrorBadRequest)
			return
		}
		response, err := service.GetCommonFriendList(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		//Send result response
		json.NewEncoder(w).Encode(response)
	}
}

func createSubscribe(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SubscriptionRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrorBadRequest)
			return
		}
		response, err := service.CreateSubscribe(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		//Send result response
		json.NewEncoder(w).Encode(response)

	}
}

func createBlockFriend(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.BlockRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrorBadRequest)
			return
		}
		response, err := service.CreateBlockFriend(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		//Send result response
		json.NewEncoder(w).Encode(response)
	}
}

func receiveFriendUpdate(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.EmailRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrorBadRequest)
			return
		}
		response, err := service.CreateUpdateReceive(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		//Send result response
		json.NewEncoder(w).Encode(response)
	}
	return nil
}

func getcommonFriend(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqq := &models.CommonFriendRequest{}
		if err := render.Bind(r, reqq); err != nil {
			render.Render(w, r, ErrorBadRequest)
			return
		}
		response, err := service.GetCommonFriendList(reqq)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		//Send result response
		json.NewEncoder(w).Encode(response)
	}
}
