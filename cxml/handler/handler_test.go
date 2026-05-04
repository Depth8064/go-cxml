package handler

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
)

type handlerStub struct {
	name string
}

func (h *handlerStub) Handle(req *model.CXML) (*model.CXML, error) {
	return req, nil
}

func (h *handlerStub) Name() string { return h.name }

func TestRegistry_RegisterAndGet(t *testing.T) {
	r := NewRegistry()

	h := &handlerStub{name: "OrderRequest"}
	r.Register(h)

	got, ok := r.Get("OrderRequest")
	if !ok {
		t.Fatal("expected registered handler to exist")
	}
	if got != h {
		t.Fatal("expected returned handler to match registered handler")
	}

	_, ok = r.Get("Missing")
	if ok {
		t.Fatal("did not expect missing handler to be found")
	}
}

func TestRegistry_RegisterNil(t *testing.T) {
	r := NewRegistry()
	r.Register(nil)
	_, ok := r.Get("anything")
	if ok {
		t.Fatal("did not expect nil registration to add handler")
	}
}
