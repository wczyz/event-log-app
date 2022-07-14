package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/wczyz/event-log-app/src/models"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		log.Fatalln(err)
	}

	if result := h.db.Create(&event); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(result.Error)
		return
	}

	w.WriteHeader(http.StatusCreated)
    log.Println("Created event")
}

func (h handler) FilterEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	eventType := r.URL.Query().Get("type")

	var queries []string

	if start != "" && end != "" {
		queries = append(queries, fmt.Sprintf("time BETWEEN '%s' AND '%s'", start, end))
	}

	if eventType != "" {
		queries = append(queries, fmt.Sprintf("type = %s", eventType))
	}

	if result := h.db.Where(strings.Join(queries, " AND ")).Find(&events); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(result.Error)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
