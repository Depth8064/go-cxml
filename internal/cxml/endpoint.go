package cxml

import (
	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/Depth8064/go-cxml/cxml/serializer"
)

// Endpoint is the primary entry point for simple cXML serialization.
type Endpoint struct {
	serializer *serializer.Serializer
}

func NewEndpoint() *Endpoint {
	return &Endpoint{serializer: serializer.NewSerializer()}
}

func (e *Endpoint) Serialize(doc *model.CXML) ([]byte, error) {
	return e.serializer.Serialize(doc)
}

func (e *Endpoint) Deserialize(data []byte) (*model.CXML, error) {
	return e.serializer.Deserialize(data)
}
