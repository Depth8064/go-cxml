package builder

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
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

	if doc == nil {
		t.Fatal("expected built doc")
	}
	if got, want := doc.PayloadID, "p1"; got != want {
		t.Fatalf("unexpected payloadID: got %q want %q", got, want)
	}
	if got, want := doc.From.Identity, "From"; got != want {
		t.Fatalf("unexpected from identity: got %q want %q", got, want)
	}
	if got, want := doc.To.Identity, "To"; got != want {
		t.Fatalf("unexpected to identity: got %q want %q", got, want)
	}
}

func TestBuilder_BuildError(t *testing.T) {
	doc := New().BuildError("500", "Server Error")
	if doc.Response == nil {
		t.Fatal("expected response")
	}
	if got, want := doc.Response.Status.Code, "500"; got != want {
		t.Fatalf("unexpected status code: got %q want %q", got, want)
	}
}

func TestBuilder_ResponseMessageStatus(t *testing.T) {
	doc := New().
		Request(&model.Request{OrderRequest: &model.OrderRequest{}}).
		Response(&model.Response{Status: &model.Status{Code: "200"}}).
		Message(&model.Message{Subject: "note"}).
		Status(&model.Status{Code: "299"}).
		Build()

	if doc.Message == nil {
		t.Fatal("expected message")
	}
	if doc.Request != nil {
		t.Fatal("expected request to be cleared")
	}
	if doc.Response != nil {
		t.Fatal("expected response to be cleared")
	}
	if doc.Status == nil {
		t.Fatal("expected status")
	}
	if got, want := doc.Status.Code, "299"; got != want {
		t.Fatalf("unexpected status code: got %q want %q", got, want)
	}
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

	if doc == nil || doc.Request == nil {
		t.Fatal("expected built request doc")
	}
	if got, want := doc.Request.OrderRequest.OrderRequestHeader.OrderID, "order-1"; got != want {
		t.Fatalf("unexpected order id: got %q want %q", got, want)
	}
	if got, want := doc.Timestamp, "2026-03-24T00:00:00"; got != want {
		t.Fatalf("unexpected timestamp: got %q want %q", got, want)
	}
	if got, want := doc.Version, "1.2.014"; got != want {
		t.Fatalf("unexpected version: got %q want %q", got, want)
	}
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

	if doc == nil || doc.Request == nil {
		t.Fatal("expected built request doc")
	}
	if got, want := doc.Request.OrderChangeRequest.OrderRequestReference.OrderID, "order-1"; got != want {
		t.Fatalf("unexpected order id: got %q want %q", got, want)
	}
	if got, want := doc.Timestamp, "2026-03-24T00:00:00"; got != want {
		t.Fatalf("unexpected timestamp: got %q want %q", got, want)
	}
	if got, want := doc.Version, "1.2.014"; got != want {
		t.Fatalf("unexpected version: got %q want %q", got, want)
	}
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

	if doc == nil || doc.Request == nil || doc.Request.ShipNoticeRequest == nil {
		t.Fatal("expected built ship notice request")
	}
	if got, want := doc.Request.ShipNoticeRequest.ShipNoticeHeader.ShipmentID, "SN-001"; got != want {
		t.Fatalf("unexpected shipment id: got %q want %q", got, want)
	}
	if got, want := doc.Request.PayloadType(), "ShipNoticeRequest"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
	if got, want := doc.Version, "1.2.069"; got != want {
		t.Fatalf("unexpected version: got %q want %q", got, want)
	}
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

	if doc == nil || doc.Request == nil || doc.Request.InvoiceDetailRequest == nil {
		t.Fatal("expected built invoice detail request")
	}
	if got, want := doc.Request.InvoiceDetailRequest.InvoiceDetailRequestHeader.InvoiceID, "INV-001"; got != want {
		t.Fatalf("unexpected invoice id: got %q want %q", got, want)
	}
	if got, want := doc.Request.PayloadType(), "InvoiceDetailRequest"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
	if got, want := doc.Version, "1.2.069"; got != want {
		t.Fatalf("unexpected version: got %q want %q", got, want)
	}
}
