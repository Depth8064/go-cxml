package auth

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
)

type repoStub struct {
	count    int
	validate bool
}

func (r *repoStub) Validate(*model.Credential) bool { return r.validate }

func (r *repoStub) Find(domain, identity, sharedSecret string) (*model.Credential, bool) {
	return nil, false
}

func (r *repoStub) Count() int { return r.count }

func TestNewSimpleSharedSecretAuthenticator(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	assert.NotNil(t, a)
}

func TestAuthenticate_OpenAccessWhenNoRepo(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{}, nil)
	assert.NoError(t, err)
}

func TestAuthenticate_OpenAccessWhenRepoEmpty(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{}, &repoStub{count: 0})
	assert.NoError(t, err)
}

func TestAuthenticate_MissingCredential(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{}, &repoStub{count: 1, validate: true})
	assert.EqualError(t, err, "auth: missing sender credential")
}

func TestAuthenticate_InvalidCredential(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{
		Sender: &model.Sender{Credential: &model.Credential{Domain: "d", Identity: "i", SharedSecret: "s"}},
	}, &repoStub{count: 1, validate: false})
	assert.EqualError(t, err, "auth: invalid shared secret")
}

func TestAuthenticate_ValidCredential(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{
		Sender: &model.Sender{Credential: &model.Credential{Domain: "d", Identity: "i", SharedSecret: "s"}},
	}, &repoStub{count: 1, validate: true})
	assert.NoError(t, err)
}
