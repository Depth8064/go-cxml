package document

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRegistry_SaveAndGet(t *testing.T) {
	reg := NewInMemoryRegistry()
	c := &model.CXML{PayloadID: "x1"}
	reg.Save("x1", c)

	got, ok := reg.Get("x1")
	assert.True(t, ok)
	assert.Equal(t, c, got)
}

func TestInMemoryRegistry_GetUnknown(t *testing.T) {
	reg := NewInMemoryRegistry()
	_, ok := reg.Get("missing")
	assert.False(t, ok)
}
