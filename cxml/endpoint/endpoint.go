package endpoint

import (
	"github.com/mstrhakr/go-cxml/cxml/auth"
	"github.com/mstrhakr/go-cxml/cxml/credential"
	"github.com/mstrhakr/go-cxml/cxml/processor"
	"github.com/mstrhakr/go-cxml/cxml/serializer"
)

type Endpoint struct {
	serializer     *serializer.Serializer
	processor      *processor.Processor
	authenticator  auth.Authenticator
	credentialRepo credential.CredentialRepository
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
		serializer:     serializer.NewSerializer(),
		processor:      proc,
		authenticator:  authc,
		credentialRepo: repo,
	}
}

func (e *Endpoint) Process(input []byte) ([]byte, error) {
	doc, err := e.serializer.Deserialize(input)
	if err != nil {
		return nil, err
	}

	if err := e.authenticator.Authenticate(doc, e.credentialRepo); err != nil {
		return nil, err
	}

	out, err := e.processor.Process(doc)
	if err != nil {
		return nil, err
	}

	return e.serializer.Serialize(out)
}
