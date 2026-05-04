package handler

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
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
	assert.True(t, ok)
	assert.Equal(t, h, got)

	_, ok = r.Get("Missing")
	assert.False(t, ok)
}

func TestRegistry_RegisterNil(t *testing.T) {
	r := NewRegistry()
	r.Register(nil)
	_, ok := r.Get("anything")
	assert.False(t, ok)
}
