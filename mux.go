package spade

import (
	"fmt"
	"log"
	"net/http"

	"github.com/allyraza/spade/handlers"
	"github.com/allyraza/spade/pkg/pipeline"
)

func (s *Spade) initMux() {

	eventHandler := &handlers.EventHandler{Database: s.Database}

	apiPipeline := pipeline.New(logger, checkPost)
	webPipeline := pipeline.New(logger)

	s.Mux.HandleFunc("/track", apiPipeline.Run(eventHandler.Track))
	s.Mux.HandleFunc("/web", webPipeline.Run(eventHandler.Track))
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func checkPost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Printf("checkPost: invalid request method (%v)\n", r.Method)
			fmt.Fprint(w, "bad request")
			return
		}
		next.ServeHTTP(w, r)
	}
}
