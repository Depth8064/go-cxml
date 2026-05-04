package builder

import "github.com/Depth8064/go-cxml/cxml/model"

// ProfileRequestBuilder builds a cXML ProfileRequest document.
type ProfileRequestBuilder struct {
	builder *Builder
}

func NewProfileRequestBuilder() *ProfileRequestBuilder {
	return &ProfileRequestBuilder{builder: New()}
}

func (b *ProfileRequestBuilder) PayloadID(id string) *ProfileRequestBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *ProfileRequestBuilder) Timestamp(ts string) *ProfileRequestBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *ProfileRequestBuilder) Version(version string) *ProfileRequestBuilder {
	b.builder.Version(version)
	return b
}

func (b *ProfileRequestBuilder) From(party *model.Party) *ProfileRequestBuilder {
	b.builder.From(party)
	return b
}

func (b *ProfileRequestBuilder) To(party *model.Party) *ProfileRequestBuilder {
	b.builder.To(party)
	return b
}

func (b *ProfileRequestBuilder) Sender(sender *model.Sender) *ProfileRequestBuilder {
	b.builder.Sender(sender)
	return b
}

func (b *ProfileRequestBuilder) Request(req *model.ProfileRequest) *ProfileRequestBuilder {
	b.builder.Request(&model.Request{ProfileRequest: req})
	return b
}

func (b *ProfileRequestBuilder) Build() *model.CXML {
	return b.builder.Build()
}
