package auth

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
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
	if a == nil {
		t.Fatal("expected authenticator instance")
	}
}

func TestAuthenticate_OpenAccessWhenNoRepo(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{}, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestAuthenticate_OpenAccessWhenRepoEmpty(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{}, &repoStub{count: 0})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestAuthenticate_MissingCredential(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{}, &repoStub{count: 1, validate: true})
	if err == nil {
		t.Fatal("expected error")
	}
	if got, want := err.Error(), "auth: missing sender credential"; got != want {
		t.Fatalf("unexpected error: got %q want %q", got, want)
	}
}

func TestAuthenticate_InvalidCredential(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{
		Sender: &model.Sender{Credential: &model.Credential{Domain: "d", Identity: "i", SharedSecret: "s"}},
	}, &repoStub{count: 1, validate: false})
	if err == nil {
		t.Fatal("expected error")
	}
	if got, want := err.Error(), "auth: invalid shared secret"; got != want {
		t.Fatalf("unexpected error: got %q want %q", got, want)
	}
}

func TestAuthenticate_ValidCredential(t *testing.T) {
	a := NewSimpleSharedSecretAuthenticator()
	err := a.Authenticate(&model.CXML{
		Sender: &model.Sender{Credential: &model.Credential{Domain: "d", Identity: "i", SharedSecret: "s"}},
	}, &repoStub{count: 1, validate: true})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
