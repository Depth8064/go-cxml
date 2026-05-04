package cxml

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.Contains(t, string(encoded), "OrderRequest")

	decoded, err := e.Deserialize(encoded)
	assert.NoError(t, err)
	assert.Equal(t, "p-1", decoded.PayloadID)
}

func TestEndpoint_SerializeNil(t *testing.T) {
	e := NewEndpoint()

	encoded, err := e.Serialize(nil)
	assert.Error(t, err)
	assert.Nil(t, encoded)
}

func TestEndpoint_DeserializeEmpty(t *testing.T) {
	e := NewEndpoint()

	doc, err := e.Deserialize(nil)
	assert.Error(t, err)
	assert.Nil(t, doc)
}
