package document

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
)

func TestInMemoryRegistry_SaveAndGet(t *testing.T) {
	reg := NewInMemoryRegistry()
	c := &model.CXML{PayloadID: "x1"}
	reg.Save("x1", c)

	got, ok := reg.Get("x1")
	if !ok {
		t.Fatal("expected saved document to be found")
	}
	if got != c {
		t.Fatal("expected returned document pointer to match saved one")
	}
}

func TestInMemoryRegistry_GetUnknown(t *testing.T) {
	reg := NewInMemoryRegistry()
	_, ok := reg.Get("missing")
	if ok {
		t.Fatal("did not expect missing document to be found")
	}
}

func TestInMemoryRegistry_SaveInitializesNilStore(t *testing.T) {
	reg := &InMemoryRegistry{}
	c := &model.CXML{PayloadID: "x2"}
	reg.Save("x2", c)

	got, ok := reg.Get("x2")
	if !ok {
		t.Fatal("expected saved document to be found")
	}
	if got != c {
		t.Fatal("expected returned document pointer to match saved one")
	}
}

func TestInMemoryRegistry_GetOnNilStore(t *testing.T) {
	reg := &InMemoryRegistry{}
	got, ok := reg.Get("missing")
	if ok {
		t.Fatal("did not expect missing document to be found")
	}
	if got != nil {
		t.Fatal("expected nil document when store is uninitialized")
	}
}
