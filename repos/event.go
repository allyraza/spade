package repos

import (
	"github.com/allyraza/spade/models"
)

// EventRepo : maps events to database rows
type EventRepo struct{}

// FindByID : finds an event by id
func (er *EventRepo) FindByID(id string) *models.Event {
	return &models.Event{ID: "100"}
}
