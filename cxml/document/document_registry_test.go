package document

import (
	"fmt"
	"sync"
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

func TestInMemoryRegistry_ConcurrentSaveAndGet(t *testing.T) {
	reg := NewInMemoryRegistry()

	const workers = 64
	const writesPerWorker = 200

	var wg sync.WaitGroup
	wg.Add(workers)

	for w := 0; w < workers; w++ {
		go func(worker int) {
			defer wg.Done()

			for i := 0; i < writesPerWorker; i++ {
				id := fmt.Sprintf("p-%d-%d", worker, i)
				doc := &model.CXML{PayloadID: id}
				reg.Save(id, doc)

				if got, ok := reg.Get(id); !ok || got == nil || got.PayloadID != id {
					t.Fatalf("expected to retrieve payload %s", id)
				}
			}
		}(w)
	}

	wg.Wait()

	// Spot-check a few known keys from different workers.
	for _, id := range []string{"p-0-0", "p-1-199", "p-63-42"} {
		if got, ok := reg.Get(id); !ok || got == nil || got.PayloadID != id {
			t.Fatalf("expected payload %s to exist after concurrent writes", id)
		}
	}
}
