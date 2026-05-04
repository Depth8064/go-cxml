package processor

import (
	"errors"
	"testing"

	"github.com/Depth8064/go-cxml/cxml/handler"
	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
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

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "x1", resp.PayloadID)
	assert.Equal(t, "200", resp.Response.Status.Code)
}

func TestProcessor_Process_PunchOutOrderMessage(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&stubPunchOutHandler{})

	p := NewProcessor(reg)

	req := &model.CXML{PayloadID: "x2", Request: &model.Request{PunchOutOrderMessage: &model.PunchOutOrderMessage{PunchOutOrderMessageHeader: &model.PunchOutOrderMessageHeader{Operation: "create"}}}}
	resp, err := p.Process(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "x2", resp.PayloadID)
	assert.Equal(t, "201", resp.Response.Status.Code)
}

func TestProcessor_Process_OrderChangeRequest(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&stubOrderChangeHandler{})

	p := NewProcessor(reg)

	req := &model.CXML{PayloadID: "x3", Request: &model.Request{OrderChangeRequest: &model.OrderChangeRequest{OrderRequestReference: &model.OrderRequestHeader{OrderID: "PO123"}}}}
	resp, err := p.Process(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "x3", resp.PayloadID)
	assert.Equal(t, "202", resp.Response.Status.Code)
}

func TestProcessor_Process_NoHandler(t *testing.T) {
	p := NewProcessor(handler.NewRegistry())
	_, err := p.Process(&model.CXML{Request: &model.Request{OrderRequest: &model.OrderRequest{}}})
	assert.Error(t, err)
}

func TestProcessor_NewProcessor_NilRegistry(t *testing.T) {
	p := NewProcessor(nil)
	assert.NotNil(t, p)
}

func TestProcessor_Process_NilDocument(t *testing.T) {
	p := NewProcessor(nil)
	resp, err := p.Process(nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestProcessor_Process_UnsupportedPayload(t *testing.T) {
	p := NewProcessor(handler.NewRegistry())
	resp, err := p.Process(&model.CXML{})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestProcessor_Process_ResponsePayload(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&namedHandler{name: "Status"})
	p := NewProcessor(reg)

	resp, err := p.Process(&model.CXML{PayloadID: "resp1", Response: &model.Response{Status: &model.Status{Code: "200"}}})
	assert.NoError(t, err)
	assert.Equal(t, "299", resp.Response.Status.Code)
}

func TestProcessor_Process_MessagePayload(t *testing.T) {
	reg := handler.NewRegistry()
	reg.Register(&namedHandler{name: "Message"})
	p := NewProcessor(reg)

	resp, err := p.Process(&model.CXML{PayloadID: "msg1", Message: &model.Message{Subject: "hello"}})
	assert.NoError(t, err)
	assert.Equal(t, "299", resp.Response.Status.Code)
}
