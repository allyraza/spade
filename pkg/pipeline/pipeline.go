package pipeline

import "net/http"

// Middleware :
type Middleware func(http.Handler) http.Handler

// Pipeline :
type Pipeline struct {
	middlewares []Middleware
}

// New :
func New(middlewares ...Middleware) *Pipeline {
	return &Pipeline{
		middlewares: middlewares,
	}
}

// Run : wrap handler with handlers in seq
func (p *Pipeline) Run(h http.Handler) http.Handler {
	for _, m := range p.middlewares {
		h = m(h)
	}

	return h
}

// RunFunc : wrap handler with handlers in seq
func (p *Pipeline) RunFunc(h http.HandlerFunc) http.Handler {
	return p.Run(h)
}
