package endpoint

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/auth"
	"github.com/Depth8064/go-cxml/cxml/credential"
	"github.com/Depth8064/go-cxml/cxml/document"
	"github.com/Depth8064/go-cxml/cxml/handler"
	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/Depth8064/go-cxml/cxml/processor"
	"github.com/Depth8064/go-cxml/cxml/validation"
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
<!DOCTYPE cXML SYSTEM "http://xml.cxml.org/schemas/cXML/1.2.014/cXML.dtd">
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
<!DOCTYPE cXML SYSTEM "http://xml.cxml.org/schemas/cXML/1.2.014/cXML.dtd">
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

	output, err := ep.Process(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "<Response>")
	assert.Contains(t, string(output), "401")
}

func TestEndpoint_Process_DTDFail(t *testing.T) {
	registry := handler.NewRegistry()
	registry.Register(&basicOrderHandler{})

	proc := processor.NewProcessor(registry)
	repo := credential.NewRegistry([]*model.Credential{{Domain: "D", Identity: "I", SharedSecret: "S"}})
	authc := auth.NewSimpleSharedSecretAuthenticator()

	ep := NewEndpoint(proc, authc, repo)

	input := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<cXML payloadID="abc"></cXML>`)

	output, err := ep.Process(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "<Response>")
	assert.Contains(t, string(output), "400")
}

func TestEndpoint_NewEndpoint_Defaults(t *testing.T) {
	ep := NewEndpoint(nil, nil, nil)
	assert.NotNil(t, ep)

	input := []byte(`<?xml version="1.0"?><cXML payloadID="abc"></cXML>`)
	output, err := ep.Process(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "<Response>")
	assert.Contains(t, string(output), "400")
}

func TestEndpoint_SettersAndDeserializeFailurePath(t *testing.T) {
	registry := handler.NewRegistry()
	registry.Register(&basicOrderHandler{})
	ep := NewEndpoint(processor.NewProcessor(registry), auth.NewSimpleSharedSecretAuthenticator(), credential.NewRegistry(nil))

	ep.SetDTDValidator(nil)
	ep.SetDocumentRegistry(document.NewInMemoryRegistry())
	ep.SetCredentialRepository(nil)
	ep.SetCredentialRepository(credential.NewRegistry(nil))

	output, err := ep.Process([]byte("not-xml"))
	assert.NoError(t, err)
	assert.Contains(t, string(output), "<Response>")
	assert.Contains(t, string(output), "400")
}

func TestEndpoint_SetDTDValidatorWithInstance(t *testing.T) {
	ep := NewEndpoint(nil, nil, nil)
	ep.SetDTDValidator(validation.NewDTDValidator())

	input := []byte(`<?xml version="1.0"?><cXML payloadID="abc"></cXML>`)
	output, err := ep.Process(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "400")
}

func TestEndpoint_Process_ProcessorFailureReturns500(t *testing.T) {
	ep := NewEndpoint(nil, nil, nil)
	ep.SetDTDValidator(nil)

	input := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<cXML payloadID="abc" timestamp="2026-03-24T12:34:56" version="1.2.014">
  <Request>
    <OrderRequest>
      <OrderRequestHeader orderID="PO-99" orderDate="2026-03-24"/>
    </OrderRequest>
  </Request>
</cXML>`)

	output, err := ep.Process(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "<Response>")
	assert.Contains(t, string(output), "500")
}
