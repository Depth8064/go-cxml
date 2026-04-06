package handler

import "github.com/Depth8064/go-cxml/cxml/model"

type Handler interface {
	// Handle processes a request payload and returns a response or error.
	Handle(req *model.CXML) (*model.CXML, error)

	// Name identifies supported payload type (e.g., "OrderRequest").
	Name() string
}

// Registry keeps a collection of handlers by payload name.
type Registry struct {
	handlers map[string]Handler
}

func NewRegistry() *Registry {
	return &Registry{handlers: map[string]Handler{}}
}

func (r *Registry) Register(h Handler) {
	if h == nil {
		return
	}
	r.handlers[h.Name()] = h
}

func (r *Registry) Get(name string) (Handler, bool) {
	h, ok := r.handlers[name]
	return h, ok
}
