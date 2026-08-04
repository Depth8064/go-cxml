package document

import (
	"sync"

	"github.com/Depth8064/go-cxml/cxml/model"
)

type DocumentRegistry interface {
	Save(payloadID string, doc *model.CXML)
	Get(payloadID string) (*model.CXML, bool)
}

type InMemoryRegistry struct {
	mu    sync.RWMutex
	store map[string]*model.CXML
}

func NewInMemoryRegistry() *InMemoryRegistry {
	return &InMemoryRegistry{store: map[string]*model.CXML{}}
}

func (r *InMemoryRegistry) Save(payloadID string, doc *model.CXML) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.store == nil {
		r.store = map[string]*model.CXML{}
	}
	r.store[payloadID] = doc
}

func (r *InMemoryRegistry) Get(payloadID string) (*model.CXML, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.store == nil {
		return nil, false
	}
	doc, ok := r.store[payloadID]
	return doc, ok
}
