package pipeline

import "net/http"

// PipelineFunc :
type Func func(http.HandlerFunc) http.HandlerFunc

// Pipeline :
type Pipeline struct {
	handlers []Func
}

// New :
func New(handlers ...Func) *Pipeline {
	return &Pipeline{
		handlers: handlers,
	}
}

// Run : wrap handler with handlers in seq
func (p *Pipeline) Run(h http.HandlerFunc) http.HandlerFunc {

	for _, hf := range p.handlers {
		h = hf(h)
	}

	return h
}
