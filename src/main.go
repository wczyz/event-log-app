package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wczyz/event-log-app/src/db"
	"github.com/wczyz/event-log-app/src/handlers"
)

func main() {
	db := db.Init()
	h := handlers.New(db)
	router := mux.NewRouter()

	router.HandleFunc("/create", h.CreateEvent).Methods(http.MethodPost)
	router.HandleFunc("/filter", h.FilterEvents).Methods(http.MethodGet)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
