package credential

import "github.com/mstrhakr/go-cxml/cxml/model"

type CredentialRepository interface {
	Validate(cred *model.Credential) bool
	Find(domain, identity, sharedSecret string) (*model.Credential, bool)
}

type Registry struct {
	entries []*model.Credential
}

func NewRegistry(entries []*model.Credential) *Registry {
	return &Registry{entries: entries}
}

func (r *Registry) Validate(cred *model.Credential) bool {
	if cred == nil {
		return false
	}
	_, ok := r.Find(cred.Domain, cred.Identity, cred.SharedSecret)
	return ok
}

func (r *Registry) Find(domain, identity, sharedSecret string) (*model.Credential, bool) {
	for _, entry := range r.entries {
		if entry == nil {
			continue
		}
		if entry.Domain == domain && entry.Identity == identity && entry.SharedSecret == sharedSecret {
			return entry, true
		}
	}
	return nil, false
}
