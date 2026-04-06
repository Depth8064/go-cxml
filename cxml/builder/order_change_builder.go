package builder

import "github.com/Depth8064/go-cxml/cxml/model"

type OrderChangeBuilder struct {
	builder *Builder
}

func NewOrderChangeBuilder() *OrderChangeBuilder {
	return &OrderChangeBuilder{builder: New()}
}

func (b *OrderChangeBuilder) PayloadID(id string) *OrderChangeBuilder {
	b.builder.PayloadID(id)
	return b
}

func (b *OrderChangeBuilder) Timestamp(ts string) *OrderChangeBuilder {
	b.builder.Timestamp(ts)
	return b
}

func (b *OrderChangeBuilder) Version(version string) *OrderChangeBuilder {
	b.builder.Version(version)
	return b
}

func (b *OrderChangeBuilder) Request(orderChange *model.OrderChangeRequest) *OrderChangeBuilder {
	b.builder.Request(&model.Request{OrderChangeRequest: orderChange})
	return b
}

func (b *OrderChangeBuilder) Build() *model.CXML {
	return b.builder.Build()
}
