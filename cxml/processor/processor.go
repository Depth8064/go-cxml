package processor

import (
	"errors"

	"github.com/Depth8064/go-cxml/cxml/handler"
	"github.com/Depth8064/go-cxml/cxml/model"
)

type Processor struct {
	registry *handler.Registry
}

func NewProcessor(reg *handler.Registry) *Processor {
	if reg == nil {
		reg = handler.NewRegistry()
	}
	return &Processor{registry: reg}
}

func (p *Processor) Process(doc *model.CXML) (*model.CXML, error) {
	if doc == nil {
		return nil, errors.New("cxml: nil document")
	}

	var payloadName string
	if doc.IsRequest() {
		payloadName = doc.Request.PayloadType()
	} else if doc.IsResponse() {
		payloadName = "Response"
	} else if doc.IsMessage() {
		payloadName = "Message"
	}

	if payloadName == "" {
		return nil, errors.New("cxml: unsupported payload type")
	}

	h, ok := p.registry.Get(payloadName)
	if !ok {
		return nil, errors.New("cxml: no handler registered")
	}

	return h.Handle(doc)
}
