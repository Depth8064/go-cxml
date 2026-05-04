package processor

import (
	"errors"
	"testing"

	"github.com/Depth8064/go-cxml/cxml/handler"
	"github.com/Depth8064/go-cxml/cxml/model"
)

type stubOrderHandler struct{}

func (h *stubOrderHandler) Handle(req *model.CXML) (*model.CXML, error) {
	resp := &model.CXML{
		PayloadID: req.PayloadID,
		Response:  &model.Response{Status: &model.Status{Code: "200", Text: "OK"}},
	}
	return resp, nil
}

func (h *stubOrderHandler) Name() string { return "OrderRequest" }

type stubPunchOutHandler struct{}

func (h *stubPunchOutHandler) Handle(req *model.CXML) (*model.CXML, error) {
	return &model.CXML{PayloadID: req.PayloadID, Response: &model.Response{Status: &model.Status{Code: "201", Text: "PunchOut OK"}}}, nil
}

func (h *stubPunchOutHandler) Name() string { return "PunchOutOrderMessage" }

type stubOrderChangeHandler struct{}

func (h *stubOrderChangeHandler) Handle(req *model.CXML) (*model.CXML, error) {
	return &model.CXML{PayloadID: req.PayloadID, Response: &model.Response{Status: &model.Status{Code: "202", Text: "OrderChange Accepted"}}}, nil
}

func (h *stubOrderChangeHandler) Name() string { return "OrderChangeRequest" }

type namedHandler struct {
	name string
}

func (h *namedHandler) Handle(req *model.CXML) (*model.CXML, error) {
	if req == nil {
		return nil, errors.New("nil")
	}
	return &model.CXML{PayloadID: req.PayloadID, Response: &model.Response{Status: &model.Status{Code: "299", Text: h.name}}}, nil
}

func (h *namedHandler) Name() string { return h.name }

func TestProcessor_Process_OrderRequest(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&stubOrderHandler{})

	p := NewProcessor(reg)

	req := &model.CXML{PayloadID: "x1", Request: &model.Request{OrderRequest: &model.OrderRequest{OrderRequestHeader: &model.OrderRequestHeader{OrderID: "PO"}}}}
	resp, err := p.Process(req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("expected response")
	}
	if got, want := resp.PayloadID, "x1"; got != want {
		t.Fatalf("unexpected payload id: got %q want %q", got, want)
	}
	if got, want := resp.Response.Status.Code, "200"; got != want {
		t.Fatalf("unexpected status code: got %q want %q", got, want)
	}
}

func TestProcessor_Process_PunchOutOrderMessage(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&stubPunchOutHandler{})

	p := NewProcessor(reg)

	req := &model.CXML{PayloadID: "x2", Request: &model.Request{PunchOutOrderMessage: &model.PunchOutOrderMessage{PunchOutOrderMessageHeader: &model.PunchOutOrderMessageHeader{Operation: "create"}}}}
	resp, err := p.Process(req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("expected response")
	}
	if got, want := resp.PayloadID, "x2"; got != want {
		t.Fatalf("unexpected payload id: got %q want %q", got, want)
	}
	if got, want := resp.Response.Status.Code, "201"; got != want {
		t.Fatalf("unexpected status code: got %q want %q", got, want)
	}
}

func TestProcessor_Process_OrderChangeRequest(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&stubOrderChangeHandler{})

	p := NewProcessor(reg)

	req := &model.CXML{PayloadID: "x3", Request: &model.Request{OrderChangeRequest: &model.OrderChangeRequest{OrderRequestReference: &model.OrderRequestHeader{OrderID: "PO123"}}}}
	resp, err := p.Process(req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("expected response")
	}
	if got, want := resp.PayloadID, "x3"; got != want {
		t.Fatalf("unexpected payload id: got %q want %q", got, want)
	}
	if got, want := resp.Response.Status.Code, "202"; got != want {
		t.Fatalf("unexpected status code: got %q want %q", got, want)
	}
}

func TestProcessor_Process_NoHandler(t *testing.T) {
	p := NewProcessor(handler.NewRegistry())
	_, err := p.Process(&model.CXML{Request: &model.Request{OrderRequest: &model.OrderRequest{}}})
	if err == nil {
		t.Fatal("expected error for missing handler")
	}
}

func TestProcessor_NewProcessor_NilRegistry(t *testing.T) {
	p := NewProcessor(nil)
	if p == nil {
		t.Fatal("expected processor instance")
	}
}

func TestProcessor_Process_NilDocument(t *testing.T) {
	p := NewProcessor(nil)
	resp, err := p.Process(nil)
	if err == nil {
		t.Fatal("expected error for nil document")
	}
	if resp != nil {
		t.Fatal("expected nil response on error")
	}
}

func TestProcessor_Process_UnsupportedPayload(t *testing.T) {
	p := NewProcessor(handler.NewRegistry())
	resp, err := p.Process(&model.CXML{})
	if err == nil {
		t.Fatal("expected error for unsupported payload")
	}
	if resp != nil {
		t.Fatal("expected nil response on error")
	}
}

func TestProcessor_Process_ResponsePayload(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&namedHandler{name: "Status"})
	p := NewProcessor(reg)

	resp, err := p.Process(&model.CXML{PayloadID: "resp1", Response: &model.Response{Status: &model.Status{Code: "200"}}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got, want := resp.Response.Status.Code, "299"; got != want {
		t.Fatalf("unexpected status code: got %q want %q", got, want)
	}
}

func TestProcessor_Process_MessagePayload(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&namedHandler{name: "Message"})
	p := NewProcessor(reg)

	resp, err := p.Process(&model.CXML{PayloadID: "msg1", Message: &model.Message{Subject: "hello"}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got, want := resp.Response.Status.Code, "299"; got != want {
		t.Fatalf("unexpected status code: got %q want %q", got, want)
	}
}
