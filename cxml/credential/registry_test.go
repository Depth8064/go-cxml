package credential

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
)

func TestRegistry_Count(t *testing.T) {
	r := NewRegistry([]*model.Credential{{Domain: "d", Identity: "i", SharedSecret: "s"}, nil})
	assert.Equal(t, 2, r.Count())
}

func TestRegistry_FindAndValidate(t *testing.T) {
	match := &model.Credential{Domain: "d", Identity: "i", SharedSecret: "s"}
	r := NewRegistry([]*model.Credential{nil, match})

	found, ok := r.Find("d", "i", "s")
	assert.True(t, ok)
	assert.Equal(t, match, found)

	assert.True(t, r.Validate(&model.Credential{Domain: "d", Identity: "i", SharedSecret: "s"}))
	assert.False(t, r.Validate(&model.Credential{Domain: "d", Identity: "x", SharedSecret: "s"}))
	assert.False(t, r.Validate(nil))

	missing, ok := r.Find("x", "i", "s")
	assert.False(t, ok)
	assert.Nil(t, missing)
}
