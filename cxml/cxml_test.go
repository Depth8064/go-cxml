package cxml

import (
	"strings"
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
)

func TestEndpoint_SerializeDeserialize(t *testing.T) {
	e := NewEndpoint()

	doc := &model.CXML{
		PayloadID: "p-1",
		Request: &model.Request{
			OrderRequest: &model.OrderRequest{},
		},
	}

	encoded, err := e.Serialize(doc)
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}
	if !strings.Contains(string(encoded), "OrderRequest") {
		t.Fatal("expected serialized output to contain OrderRequest")
	}

	decoded, err := e.Deserialize(encoded)
	if err != nil {
		t.Fatalf("deserialize failed: %v", err)
	}
	if got, want := decoded.PayloadID, "p-1"; got != want {
		t.Fatalf("unexpected payload id: got %q want %q", got, want)
	}
}

func TestEndpoint_SerializeNil(t *testing.T) {
	e := NewEndpoint()

	encoded, err := e.Serialize(nil)
	if err == nil {
		t.Fatal("expected serialize error for nil input")
	}
	if encoded != nil {
		t.Fatal("expected nil encoded output on error")
	}
}

func TestEndpoint_DeserializeEmpty(t *testing.T) {
	e := NewEndpoint()

	doc, err := e.Deserialize(nil)
	if err == nil {
		t.Fatal("expected deserialize error for empty input")
	}
	if doc != nil {
		t.Fatal("expected nil document on error")
	}
}
