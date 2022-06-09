package handler

import (
	"Assignment/database"
	"Assignment/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var dbIns database.Database

func users(router chi.Router) {
	st := service.Storage{Db: dbIns}
	router.Post("/register", createUser(st))
	router.Post("/friendconnection", createFriendConnection(st))
	router.Post("/friendlist", getFriendList(st))
	router.Post("/commonFriend", getcommonFriend(st))
	router.Post("/subscribeFriend", createSubscribe(st))
	router.Post("/blockFriend", createBlockFriend(st))
	router.Post("/receiveUpdate", receiveFriendUpdate(st))
}

func notfoundhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(400)
	err := render.Render(w, r, ErrorNotFound)
	if err != nil {
		return
	}
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(405)
	err := render.Render(w, r, ErrorMethodNotAllowed)
	if err != nil {
		return
	}
}

//Create router(New Handler)
func NewHandler(db database.Database) *chi.Mux {
	router := chi.NewRouter()
	dbIns = db
	router.MethodNotAllowed(MethodNotAllowed)
	router.NotFound(notfoundhandler)
	router.Route("/api", users)
	return router
}
