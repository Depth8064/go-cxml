package auth

import (
	"errors"

	"github.com/Depth8064/go-cxml/cxml/credential"
	"github.com/Depth8064/go-cxml/cxml/model"
)

type Authenticator interface {
	Authenticate(header *model.CXML, repo credential.CredentialRepository) error
}

type SimpleSharedSecretAuthenticator struct{}

func NewSimpleSharedSecretAuthenticator() *SimpleSharedSecretAuthenticator {
	return &SimpleSharedSecretAuthenticator{}
}

func (a *SimpleSharedSecretAuthenticator) Authenticate(c *model.CXML, repo credential.CredentialRepository) error {
	if repo == nil || repo.Count() == 0 {
		// No credentials configured means open access (legacy behavior when SharedSecret is blank)
		return nil
	}
	if c == nil || c.Sender == nil || c.Sender.Credential == nil {
		return errors.New("auth: missing sender credential")
	}
	cred := c.Sender.Credential
	if !repo.Validate(cred) {
		return errors.New("auth: invalid shared secret")
	}
	return nil
}
