package spade

import (
	"database/sql"
	"net/http"
)

// Config :
type Config struct {
	Address string
}

// Spade :
type Spade struct {
	Address  string
	Mux      *http.ServeMux
	Database *sql.DB
}

// New : creates a new spade
func New(config *Config) *Spade {
	s := &Spade{
		Address: config.Address,
		Mux:     http.NewServeMux(),
	}

	// register routes
	s.initMux()

	return s
}
