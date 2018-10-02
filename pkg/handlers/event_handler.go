package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

// EventHandler :
type EventHandler struct {
	Database *sql.DB
}

// Track : collects a new event from post request
func (ec *EventHandler) Track(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "event created successuflly")
}
