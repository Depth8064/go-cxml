package builder

import "github.com/Depth8064/go-cxml/cxml/model"

// StatusUpdateBuilder builds a cXML StatusUpdateRequest document.
type StatusUpdateBuilder struct {
	builder *Builder
}

func NewStatusUpdateBuilder() *StatusUpdateBuilder {
	return &StatusUpdateBuilder{builder: New()}
}

func (b *StatusUpdateBuilder) PayloadID(id string) *StatusUpdateBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *StatusUpdateBuilder) Timestamp(ts string) *StatusUpdateBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *StatusUpdateBuilder) Version(version string) *StatusUpdateBuilder {
	b.builder.Version(version)
	return b
}

func (b *StatusUpdateBuilder) From(party *model.Party) *StatusUpdateBuilder {
	b.builder.From(party)
	return b
}

func (b *StatusUpdateBuilder) To(party *model.Party) *StatusUpdateBuilder {
	b.builder.To(party)
	return b
}

func (b *StatusUpdateBuilder) Sender(sender *model.Sender) *StatusUpdateBuilder {
	b.builder.Sender(sender)
	return b
}

func (b *StatusUpdateBuilder) Request(req *model.StatusUpdateRequest) *StatusUpdateBuilder {
	b.builder.Request(&model.Request{StatusUpdateRequest: req})
	return b
}

func (b *StatusUpdateBuilder) Build() *model.CXML {
	return b.builder.Build()
}
