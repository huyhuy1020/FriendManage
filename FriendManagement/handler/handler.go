package handler

import (
	"Assignemnt/database"
	"github.com/go-chi/chi"
)

var dbIns database.database

//Create router(New Handler)

func NewHandler(db database.database) {
	router := chi.NewRouter()
	dbIns = db
	router.Method
}
