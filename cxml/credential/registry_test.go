package credential

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
)

func TestRegistry_Count(t *testing.T) {
	r := NewRegistry([]*model.Credential{{Domain: "d", Identity: "i", SharedSecret: "s"}, nil})
	if got, want := r.Count(), 2; got != want {
		t.Fatalf("unexpected count: got %d want %d", got, want)
	}
}

func TestRegistry_FindAndValidate(t *testing.T) {
	match := &model.Credential{Domain: "d", Identity: "i", SharedSecret: "s"}
	r := NewRegistry([]*model.Credential{nil, match})

	found, ok := r.Find("d", "i", "s")
	if !ok {
		t.Fatal("expected credential to be found")
	}
	if found != match {
		t.Fatal("expected found credential pointer to match")
	}

	if !r.Validate(&model.Credential{Domain: "d", Identity: "i", SharedSecret: "s"}) {
		t.Fatal("expected credential to validate")
	}
	if r.Validate(&model.Credential{Domain: "d", Identity: "x", SharedSecret: "s"}) {
		t.Fatal("expected credential to fail validation")
	}
	if r.Validate(nil) {
		t.Fatal("expected nil credential to fail validation")
	}

	missing, ok := r.Find("x", "i", "s")
	if ok {
		t.Fatal("did not expect credential match")
	}
	if missing != nil {
		t.Fatal("expected missing credential to be nil")
	}
}
