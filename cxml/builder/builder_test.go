package builder

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Basic(t *testing.T) {
	doc := New().
		PayloadID("p1").
		Timestamp("2026-03-24T00:00:00").
		Version("1.2.014").
		From(&model.Party{Identity: "From"}).
		To(&model.Party{Identity: "To"}).
		Sender(&model.Sender{UserAgent: "go-cxml"}).
		Build()

	assert.NotNil(t, doc)
	assert.Equal(t, "p1", doc.PayloadID)
	assert.Equal(t, "From", doc.From.Identity)
	assert.Equal(t, "To", doc.To.Identity)
}

func TestBuilder_BuildError(t *testing.T) {
	doc := New().BuildError("500", "Server Error")
	assert.NotNil(t, doc.Response)
	assert.Equal(t, "500", doc.Response.Status.Code)
}

func TestOrderRequestBuilder(t *testing.T) {
	order := &model.OrderRequest{
		OrderRequestHeader: &model.OrderRequestHeader{OrderID: "order-1", OrderDate: "2026-03-24"},
	}
	doc := NewOrderRequestBuilder().
		PayloadID("p2").
		Request(order).
		Build()

	assert.NotNil(t, doc)
	assert.NotNil(t, doc.Request)
	assert.Equal(t, "order-1", doc.Request.OrderRequest.OrderRequestHeader.OrderID)
}

func TestOrderChangeBuilder(t *testing.T) {
	change := &model.OrderChangeRequest{
		OrderRequestReference: &model.OrderRequestHeader{OrderID: "order-1", OrderDate: "2026-03-24"},
	}
	doc := NewOrderChangeBuilder().
		PayloadID("p3").
		Request(change).
		Build()

	assert.NotNil(t, doc)
	assert.NotNil(t, doc.Request)
	assert.Equal(t, "order-1", doc.Request.OrderChangeRequest.OrderRequestReference.OrderID)
}
