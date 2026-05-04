package builder

import "github.com/Depth8064/go-cxml/cxml/model"

// PunchOutSetupBuilder builds a cXML PunchOutSetupRequest document.
type PunchOutSetupBuilder struct {
	builder *Builder
}

func NewPunchOutSetupBuilder() *PunchOutSetupBuilder {
	return &PunchOutSetupBuilder{builder: New()}
}

func (b *PunchOutSetupBuilder) PayloadID(id string) *PunchOutSetupBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *PunchOutSetupBuilder) Timestamp(ts string) *PunchOutSetupBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *PunchOutSetupBuilder) Version(version string) *PunchOutSetupBuilder {
	b.builder.Version(version)
	return b
}

func (b *PunchOutSetupBuilder) From(party *model.Party) *PunchOutSetupBuilder {
	b.builder.From(party)
	return b
}

func (b *PunchOutSetupBuilder) To(party *model.Party) *PunchOutSetupBuilder {
	b.builder.To(party)
	return b
}

func (b *PunchOutSetupBuilder) Sender(sender *model.Sender) *PunchOutSetupBuilder {
	b.builder.Sender(sender)
	return b
}

func (b *PunchOutSetupBuilder) Request(req *model.PunchOutSetupRequest) *PunchOutSetupBuilder {
	b.builder.Request(&model.Request{PunchOutSetupRequest: req})
	return b
}

func (b *PunchOutSetupBuilder) Build() *model.CXML {
	return b.builder.Build()
}
