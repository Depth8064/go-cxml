package endpoint

import (
	"github.com/Depth8064/go-cxml/cxml/auth"
	"github.com/Depth8064/go-cxml/cxml/builder"
	"github.com/Depth8064/go-cxml/cxml/credential"
	"github.com/Depth8064/go-cxml/cxml/document"
	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/Depth8064/go-cxml/cxml/processor"
	"github.com/Depth8064/go-cxml/cxml/serializer"
	"github.com/Depth8064/go-cxml/cxml/validation"
)

type Endpoint struct {
	serializer       *serializer.Serializer
	processor        *processor.Processor
	authenticator    auth.Authenticator
	credentialRepo   credential.CredentialRepository
	dtdValidator     *validation.DTDValidator
	documentRegistry document.DocumentRegistry
}

func NewEndpoint(proc *processor.Processor, authc auth.Authenticator, repo credential.CredentialRepository) *Endpoint {
	if proc == nil {
		proc = processor.NewProcessor(nil)
	}
	if authc == nil {
		authc = auth.NewSimpleSharedSecretAuthenticator()
	}
	if repo == nil {
		repo = credential.NewRegistry(nil)
	}
	return &Endpoint{
		serializer:       serializer.NewSerializer(),
		processor:        proc,
		authenticator:    authc,
		credentialRepo:   repo,
		dtdValidator:     validation.NewDTDValidator(),
		documentRegistry: document.NewInMemoryRegistry(),
	}
}

func (e *Endpoint) SetDTDValidator(v *validation.DTDValidator) {
	e.dtdValidator = v
}

func (e *Endpoint) SetDocumentRegistry(r document.DocumentRegistry) {
	e.documentRegistry = r
}

func (e *Endpoint) SetCredentialRepository(r credential.CredentialRepository) {
	if r != nil {
		e.credentialRepo = r
	}
}

func (e *Endpoint) Process(input []byte) ([]byte, error) {
	if e.dtdValidator != nil {
		if err := e.dtdValidator.Validate(input); err != nil {
			return e.serializeErrorResponse("400", err.Error())
		}
	}

	doc, err := e.serializer.Deserialize(input)
	if err != nil {
		return e.serializeErrorResponse("400", err.Error())
	}

	if err := e.authenticator.Authenticate(doc, e.credentialRepo); err != nil {
		return e.serializeErrorResponse("401", err.Error())
	}

	if e.documentRegistry != nil && doc != nil && doc.PayloadID != "" {
		e.documentRegistry.Save(doc.PayloadID, doc)
	}

	out, err := e.processor.Process(doc)
	if err != nil {
		return e.serializeErrorResponse("500", err.Error())
	}

	if e.documentRegistry != nil && out != nil && out.PayloadID != "" {
		e.documentRegistry.Save(out.PayloadID, out)
	}

	return e.serializer.Serialize(out)
}

func (e *Endpoint) serializeErrorResponse(code, message string) ([]byte, error) {
	errorDoc := builder.New().
		PayloadID("error").
		Version("1.2.014").
		Response(&model.Response{Status: &model.Status{Code: code, Text: message}}).
		Build()

	return e.serializer.Serialize(errorDoc)
}
