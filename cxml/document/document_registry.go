package document

import "github.com/Depth8064/go-cxml/cxml/model"

type DocumentRegistry interface {
	Save(payloadID string, doc *model.CXML)
	Get(payloadID string) (*model.CXML, bool)
}

type InMemoryRegistry struct {
	store map[string]*model.CXML
}

func NewInMemoryRegistry() *InMemoryRegistry {
	return &InMemoryRegistry{store: map[string]*model.CXML{}}
}

func (r *InMemoryRegistry) Save(payloadID string, doc *model.CXML) {
	if r.store == nil {
		r.store = map[string]*model.CXML{}
	}
	r.store[payloadID] = doc
}

func (r *InMemoryRegistry) Get(payloadID string) (*model.CXML, bool) {
	if r.store == nil {
		return nil, false
	}
	doc, ok := r.store[payloadID]
	return doc, ok
}
