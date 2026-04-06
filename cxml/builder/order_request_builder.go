package builder

import "github.com/Depth8064/go-cxml/cxml/model"

type OrderRequestBuilder struct {
	builder *Builder
}

func NewOrderRequestBuilder() *OrderRequestBuilder {
	return &OrderRequestBuilder{builder: New()}
}

func (b *OrderRequestBuilder) PayloadID(id string) *OrderRequestBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *OrderRequestBuilder) Timestamp(ts string) *OrderRequestBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *OrderRequestBuilder) Version(version string) *OrderRequestBuilder {
	b.builder.Version(version)
	return b
}

func (b *OrderRequestBuilder) From(party *model.Party) *OrderRequestBuilder {
	b.builder.From(party)
	return b
}

func (b *OrderRequestBuilder) To(party *model.Party) *OrderRequestBuilder {
	b.builder.To(party)
	return b
}

func (b *OrderRequestBuilder) Sender(sender *model.Sender) *OrderRequestBuilder {
	b.builder.Sender(sender)
	return b
}

func (b *OrderRequestBuilder) Request(orderRequest *model.OrderRequest) *OrderRequestBuilder {
	b.builder.Request(&model.Request{OrderRequest: orderRequest})
	return b
}

func (b *OrderRequestBuilder) Build() *model.CXML {
	return b.builder.Build()
}
