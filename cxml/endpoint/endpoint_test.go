package endpoint

import (
	"testing"

	"github.com/mstrhakr/go-cxml/cxml/auth"
	"github.com/mstrhakr/go-cxml/cxml/credential"
	"github.com/mstrhakr/go-cxml/cxml/handler"
	"github.com/mstrhakr/go-cxml/cxml/model"
	"github.com/mstrhakr/go-cxml/cxml/processor"
	"github.com/stretchr/testify/assert"
)

type basicOrderHandler struct{}

func (h *basicOrderHandler) Handle(req *model.CXML) (*model.CXML, error) {
	return &model.CXML{PayloadID: req.PayloadID, Response: &model.Response{Status: &model.Status{Code: "200", Text: "OK"}}}, nil
}

func (h *basicOrderHandler) Name() string { return "OrderRequest" }

func TestEndpoint_Process_Success(t *testing.T) {
	registry := handler.NewRegistry()
	registry.Register(&basicOrderHandler{})

	proc := processor.NewProcessor(registry)
	repo := credential.NewRegistry([]*model.Credential{{Domain: "D", Identity: "I", SharedSecret: "S"}})
	authc := auth.NewSimpleSharedSecretAuthenticator()

	ep := NewEndpoint(proc, authc, repo)

	input := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<cXML payloadID="abc" timestamp="2026-03-24T12:34:56" version="1.2.014">
  <Header>
    <Sender>
      <Credential domain="D">
        <Identity>I</Identity>
        <SharedSecret>S</SharedSecret>
      </Credential>
      <UserAgent>go-cxml</UserAgent>
    </Sender>
  </Header>
  <Request>
    <OrderRequest>
      <OrderRequestHeader orderID="PO-99" orderDate="2026-03-24"/>
    </OrderRequest>
  </Request>
</cXML>`)

	output, err := ep.Process(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "<Response>")
	assert.Contains(t, string(output), "200")
}

func TestEndpoint_Process_AuthFail(t *testing.T) {
	registry := handler.NewRegistry()
	registry.Register(&basicOrderHandler{})

	proc := processor.NewProcessor(registry)
	repo := credential.NewRegistry([]*model.Credential{{Domain: "D", Identity: "I", SharedSecret: "S"}})
	authc := auth.NewSimpleSharedSecretAuthenticator()

	ep := NewEndpoint(proc, authc, repo)

	input := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<cXML payloadID="abc" timestamp="2026-03-24T12:34:56" version="1.2.014">
  <Header>
    <Sender>
      <Credential domain="D">
        <Identity>I</Identity>
        <SharedSecret>WRONG</SharedSecret>
      </Credential>
    </Sender>
  </Header>
  <Request>
    <OrderRequest>
      <OrderRequestHeader orderID="PO-99" orderDate="2026-03-24"/>
    </OrderRequest>
  </Request>
</cXML>`)

	_, err := ep.Process(input)
	assert.Error(t, err)
}
