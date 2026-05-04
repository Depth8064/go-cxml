package builder

import "github.com/Depth8064/go-cxml/cxml/model"

// ConfirmationRequestBuilder builds a cXML ConfirmationRequest document.
type ConfirmationRequestBuilder struct {
	builder *Builder
}

func NewConfirmationRequestBuilder() *ConfirmationRequestBuilder {
	return &ConfirmationRequestBuilder{builder: New()}
}

func (b *ConfirmationRequestBuilder) PayloadID(id string) *ConfirmationRequestBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *ConfirmationRequestBuilder) Timestamp(ts string) *ConfirmationRequestBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *ConfirmationRequestBuilder) Version(version string) *ConfirmationRequestBuilder {
	b.builder.Version(version)
	return b
}

func (b *ConfirmationRequestBuilder) From(party *model.Party) *ConfirmationRequestBuilder {
	b.builder.From(party)
	return b
}

func (b *ConfirmationRequestBuilder) To(party *model.Party) *ConfirmationRequestBuilder {
	b.builder.To(party)
	return b
}

func (b *ConfirmationRequestBuilder) Sender(sender *model.Sender) *ConfirmationRequestBuilder {
	b.builder.Sender(sender)
	return b
}

func (b *ConfirmationRequestBuilder) Request(cr *model.ConfirmationRequest) *ConfirmationRequestBuilder {
	b.builder.Request(&model.Request{ConfirmationRequest: cr})
	return b
}

func (b *ConfirmationRequestBuilder) Build() *model.CXML {
	return b.builder.Build()
}
