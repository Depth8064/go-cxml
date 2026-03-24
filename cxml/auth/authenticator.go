package auth

import (
	"errors"

	"github.com/mstrhakr/go-cxml/cxml/credential"
	"github.com/mstrhakr/go-cxml/cxml/model"
)

type Authenticator interface {
	Authenticate(header *model.CXML, repo credential.CredentialRepository) error
}

type SimpleSharedSecretAuthenticator struct{}

func NewSimpleSharedSecretAuthenticator() *SimpleSharedSecretAuthenticator {
	return &SimpleSharedSecretAuthenticator{}
}

func (a *SimpleSharedSecretAuthenticator) Authenticate(c *model.CXML, repo credential.CredentialRepository) error {
	if c == nil || c.Sender == nil || c.Sender.Credential == nil {
		return errors.New("auth: missing sender credential")
	}
	cred := c.Sender.Credential
	if !repo.Validate(cred) {
		return errors.New("auth: invalid shared secret")
	}
	return nil
}
