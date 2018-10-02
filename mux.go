package spade

import (
	"fmt"
	"log"
	"net/http"

	"github.com/allyraza/spade/pkg/handlers"
	"github.com/allyraza/spade/pkg/pipeline"
)

func (s *Spade) initMux() {

	eventHandler := &handlers.EventHandler{Database: s.Database}

	// pipeline for api
	api := pipeline.New(
		Logger,
		CheckPost,
	)

	// pipeline of middlewares for web handlers
	web := pipeline.New(
		Logger,
	)

	s.Mux.Handle("/track", api.RunFunc(eventHandler.Track))

	s.Mux.Handle("/web", web.RunFunc(eventHandler.Track))
}

// Logger : logs request method and path
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// CheckPost : allows post request method only
func CheckPost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Printf("checkPost: invalid request method (%v)\n", r.Method)
			fmt.Fprint(w, "bad request\n")
			return
		}

		next.ServeHTTP(w, r)
	})
}
