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

func TestBuilder_ResponseMessageStatus(t *testing.T) {
	doc := New().
		Request(&model.Request{OrderRequest: &model.OrderRequest{}}).
		Response(&model.Response{Status: &model.Status{Code: "200"}}).
		Message(&model.Message{Subject: "note"}).
		Status(&model.Status{Code: "299"}).
		Build()

	assert.NotNil(t, doc.Message)
	assert.Nil(t, doc.Request)
	assert.Nil(t, doc.Response)
	assert.NotNil(t, doc.Status)
	assert.Equal(t, "299", doc.Status.Code)
}

func TestOrderRequestBuilder(t *testing.T) {
	order := &model.OrderRequest{
		OrderRequestHeader: &model.OrderRequestHeader{OrderID: "order-1", OrderDate: "2026-03-24"},
	}
	doc := NewOrderRequestBuilder().
		PayloadID("p2").
		Timestamp("2026-03-24T00:00:00").
		Version("1.2.014").
		From(&model.Party{Identity: "Buyer"}).
		To(&model.Party{Identity: "Supplier"}).
		Sender(&model.Sender{UserAgent: "go-cxml-test"}).
		Request(order).
		Build()

	assert.NotNil(t, doc)
	assert.NotNil(t, doc.Request)
	assert.Equal(t, "order-1", doc.Request.OrderRequest.OrderRequestHeader.OrderID)
	assert.Equal(t, "2026-03-24T00:00:00", doc.Timestamp)
	assert.Equal(t, "1.2.014", doc.Version)
}

func TestOrderChangeBuilder(t *testing.T) {
	change := &model.OrderChangeRequest{
		OrderRequestReference: &model.OrderRequestHeader{OrderID: "order-1", OrderDate: "2026-03-24"},
	}
	doc := NewOrderChangeBuilder().
		PayloadID("p3").
		Timestamp("2026-03-24T00:00:00").
		Version("1.2.014").
		Request(change).
		Build()

	assert.NotNil(t, doc)
	assert.NotNil(t, doc.Request)
	assert.Equal(t, "order-1", doc.Request.OrderChangeRequest.OrderRequestReference.OrderID)
	assert.Equal(t, "2026-03-24T00:00:00", doc.Timestamp)
	assert.Equal(t, "1.2.014", doc.Version)
}

func TestShipNoticeBuilder(t *testing.T) {
	sn := &model.ShipNoticeRequest{
		ShipNoticeHeader: &model.ShipNoticeHeader{
			ShipmentID: "SN-001",
			NoticeDate: "2026-04-01T00:00:00",
			Operation:  "new",
		},
	}
	doc := NewShipNoticeBuilder().
		PayloadID("p4").
		Timestamp("2026-04-01T00:00:00").
		Version("1.2.069").
		From(&model.Party{Identity: "From"}).
		To(&model.Party{Identity: "To"}).
		Sender(&model.Sender{UserAgent: "go-cxml-test"}).
		Request(sn).
		Build()

	assert.NotNil(t, doc)
	assert.NotNil(t, doc.Request)
	assert.NotNil(t, doc.Request.ShipNoticeRequest)
	assert.Equal(t, "SN-001", doc.Request.ShipNoticeRequest.ShipNoticeHeader.ShipmentID)
	assert.Equal(t, "ShipNoticeRequest", doc.Request.PayloadType())
	assert.Equal(t, "1.2.069", doc.Version)
}

func TestInvoiceDetailBuilder(t *testing.T) {
	inv := &model.InvoiceDetailRequest{
		InvoiceDetailRequestHeader: &model.InvoiceDetailRequestHeader{
			InvoiceID:   "INV-001",
			InvoiceDate: "2026-04-01T00:00:00",
			Operation:   "new",
		},
	}
	doc := NewInvoiceDetailBuilder().
		PayloadID("p5").
		Timestamp("2026-04-01T00:00:00").
		Version("1.2.069").
		From(&model.Party{Identity: "From"}).
		To(&model.Party{Identity: "To"}).
		Sender(&model.Sender{UserAgent: "go-cxml-test"}).
		Request(inv).
		Build()

	assert.NotNil(t, doc)
	assert.NotNil(t, doc.Request)
	assert.NotNil(t, doc.Request.InvoiceDetailRequest)
	assert.Equal(t, "INV-001", doc.Request.InvoiceDetailRequest.InvoiceDetailRequestHeader.InvoiceID)
	assert.Equal(t, "InvoiceDetailRequest", doc.Request.PayloadType())
	assert.Equal(t, "1.2.069", doc.Version)
}
