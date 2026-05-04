package builder

import "github.com/Depth8064/go-cxml/cxml/model"

// ReceivingAdviceBuilder builds a cXML ReceivingAdviceRequest document.
type ReceivingAdviceBuilder struct {
	builder *Builder
}

func NewReceivingAdviceBuilder() *ReceivingAdviceBuilder {
	return &ReceivingAdviceBuilder{builder: New()}
}

func (b *ReceivingAdviceBuilder) PayloadID(id string) *ReceivingAdviceBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *ReceivingAdviceBuilder) Timestamp(ts string) *ReceivingAdviceBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *ReceivingAdviceBuilder) Version(version string) *ReceivingAdviceBuilder {
	b.builder.Version(version)
	return b
}

func (b *ReceivingAdviceBuilder) From(party *model.Party) *ReceivingAdviceBuilder {
	b.builder.From(party)
	return b
}

func (b *ReceivingAdviceBuilder) To(party *model.Party) *ReceivingAdviceBuilder {
	b.builder.To(party)
	return b
}

func (b *ReceivingAdviceBuilder) Sender(sender *model.Sender) *ReceivingAdviceBuilder {
	b.builder.Sender(sender)
	return b
}

func (b *ReceivingAdviceBuilder) Request(req *model.ReceivingAdviceRequest) *ReceivingAdviceBuilder {
	b.builder.Request(&model.Request{ReceivingAdviceRequest: req})
	return b
}

func (b *ReceivingAdviceBuilder) Build() *model.CXML {
	return b.builder.Build()
}
