package handler

import (
	"Assignment/database"
	"Assignment/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

var dbIns database.Database

//Create router(New Handler)

func NewHandler(db database.Database) {
	router := chi.NewRouter()
	dbIns = db
	router.MethodNotAllowed(MethodNotAllowed)
	router.NotFound(notfoundhandler)
	router.Route("/api", users)
	return router
}

func users(router chi.Router) {
	st := service.Storage{Db: dbIns}
	router.Post("/register", Cre)
}

func notfoundhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrorNotFound)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrorMethodNotAllowed)
}
