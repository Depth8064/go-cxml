package processor

import (
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
