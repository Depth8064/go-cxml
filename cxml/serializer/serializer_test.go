package serializer

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
)

type encoderStub struct {
	encodeErr error
	flushErr  error
}

func (e *encoderStub) Encode(v any) error {
	if e.encodeErr != nil {
		return e.encodeErr
	}
	return nil
}

func (e *encoderStub) Flush() error {
	return e.flushErr
}

func (e *encoderStub) Indent(prefix, indent string) {}

func TestSerializeAndDeserialize(t *testing.T) {
	doc := &model.CXML{
		PayloadID: "12345",
		Timestamp: "2026-03-24T12:34:56",
		Version:   "1.2.014",
		From:      &model.Party{Identity: "FromCompany"},
		To:        &model.Party{Identity: "ToCompany"},
		Sender:    &model.Sender{UserAgent: "go-cxml"},
		Request: &model.Request{OrderRequest: &model.OrderRequest{
			OrderRequestHeader: &model.OrderRequestHeader{OrderID: "PO-1001", OrderDate: "2026-03-24"},
		}},
	}

	s := NewSerializer()
	encoded, err := s.Serialize(doc)
	assert.NoError(t, err)
	assert.Contains(t, string(encoded), "<?xml")
	assert.Contains(t, string(encoded), "OrderRequest")

	decoded, err := s.Deserialize(encoded)
	assert.NoError(t, err)
	assert.Equal(t, "12345", decoded.PayloadID)
	if assert.NotNil(t, decoded.Request) {
		assert.NotNil(t, decoded.Request.OrderRequest)
		assert.Equal(t, "PO-1001", decoded.Request.OrderRequest.OrderRequestHeader.OrderID)
	}
}

func TestDeserializeWithDoctype(t *testing.T) {
	xmlStr := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE cXML SYSTEM "http://xml.cxml.org/schemas/cXML/1.2.014/cXML.dtd">
<cXML payloadID="abc" timestamp="2026-03-24T12:34:56" version="1.2.014">
  <Header>
    <From>
      <Identity>FromA</Identity>
    </From>
  </Header>
  <Request>
    <OrderRequest>
      <OrderRequestHeader orderID="PO-99" orderDate="2026-03-24">
        <Total>
          <Money currency="USD">100.00</Money>
        </Total>
      </OrderRequestHeader>
    </OrderRequest>
  </Request>
</cXML>`

	s := NewSerializer()
	decoded, err := s.Deserialize([]byte(xmlStr))
	assert.NoError(t, err)
	assert.Equal(t, "abc", decoded.PayloadID)
	assert.NotNil(t, decoded.Request)
	assert.Equal(t, "PO-99", decoded.Request.OrderRequest.OrderRequestHeader.OrderID)
}

func TestSerializeAndDeserialize_PunchOutOrderMessage(t *testing.T) {
	doc := &model.CXML{
		PayloadID: "punch1",
		Request:   &model.Request{PunchOutOrderMessage: &model.PunchOutOrderMessage{PunchOutOrderMessageHeader: &model.PunchOutOrderMessageHeader{Operation: "create"}}},
	}

	s := NewSerializer()
	encoded, err := s.Serialize(doc)
	assert.NoError(t, err)
	assert.Contains(t, string(encoded), "PunchOutOrderMessage")

	decoded, err := s.Deserialize(encoded)
	assert.NoError(t, err)
	assert.NotNil(t, decoded.Request)
	assert.NotNil(t, decoded.Request.PunchOutOrderMessage)
	assert.Equal(t, "create", decoded.Request.PunchOutOrderMessage.PunchOutOrderMessageHeader.Operation)
}

func TestSerializeAndDeserialize_ShipNoticeRequest(t *testing.T) {
	doc := &model.CXML{
		PayloadID: "sn-round-trip",
		Request: &model.Request{
			ShipNoticeRequest: &model.ShipNoticeRequest{
				ShipNoticeHeader: &model.ShipNoticeHeader{
					ShipmentID: "SHIP-42",
					NoticeDate: "2026-04-01T00:00:00",
					Operation:  "new",
				},
			},
		},
	}

	s := NewSerializer()
	encoded, err := s.Serialize(doc)
	assert.NoError(t, err)
	assert.Contains(t, string(encoded), "ShipNoticeRequest")
	assert.Contains(t, string(encoded), "SHIP-42")

	decoded, err := s.Deserialize(encoded)
	assert.NoError(t, err)
	assert.NotNil(t, decoded.Request)
	assert.NotNil(t, decoded.Request.ShipNoticeRequest)
	assert.Equal(t, "SHIP-42", decoded.Request.ShipNoticeRequest.ShipNoticeHeader.ShipmentID)
	assert.Equal(t, "ShipNoticeRequest", decoded.Request.PayloadType())
}

func TestSerializeAndDeserialize_InvoiceDetailRequest(t *testing.T) {
	doc := &model.CXML{
		PayloadID: "inv-round-trip",
		Request: &model.Request{
			InvoiceDetailRequest: &model.InvoiceDetailRequest{
				InvoiceDetailRequestHeader: &model.InvoiceDetailRequestHeader{
					InvoiceID:   "INV-99",
					InvoiceDate: "2026-04-01T00:00:00",
				},
				InvoiceDetailSummary: &model.InvoiceDetailSummary{
					SubtotalAmount: &model.SubtotalAmount{Money: &model.Money{Currency: "USD", Value: "500.00"}},
					Tax:            &model.Tax{Money: &model.Money{Currency: "USD", Value: "0.00"}},
					NetAmount:      &model.NetAmount{Money: &model.Money{Currency: "USD", Value: "500.00"}},
				},
			},
		},
	}

	s := NewSerializer()
	encoded, err := s.Serialize(doc)
	assert.NoError(t, err)
	assert.Contains(t, string(encoded), "InvoiceDetailRequest")
	assert.Contains(t, string(encoded), "INV-99")

	decoded, err := s.Deserialize(encoded)
	assert.NoError(t, err)
	assert.NotNil(t, decoded.Request)
	assert.NotNil(t, decoded.Request.InvoiceDetailRequest)
	assert.Equal(t, "INV-99", decoded.Request.InvoiceDetailRequest.InvoiceDetailRequestHeader.InvoiceID)
	assert.Equal(t, "InvoiceDetailRequest", decoded.Request.PayloadType())
}

func TestSerializeAndDeserialize_ConfirmationRequest(t *testing.T) {
	doc := &model.CXML{
		PayloadID: "conf-round-trip",
		Request: &model.Request{
			ConfirmationRequest: &model.ConfirmationRequest{
				ConfirmationHeader: &model.ConfirmationHeader{
					ConfirmID:  "CONF-1",
					Operation:  "accept",
					NoticeDate: "2026-04-01T00:00:00",
				},
				OrderReference: &model.OrderReference{OrderID: "PO-123"},
			},
		},
	}

	s := NewSerializer()
	encoded, err := s.Serialize(doc)
	assert.NoError(t, err)
	assert.Contains(t, string(encoded), "ConfirmationRequest")
	assert.Contains(t, string(encoded), "CONF-1")

	decoded, err := s.Deserialize(encoded)
	assert.NoError(t, err)
	assert.NotNil(t, decoded.Request)
	assert.NotNil(t, decoded.Request.ConfirmationRequest)
	assert.Equal(t, "CONF-1", decoded.Request.ConfirmationRequest.ConfirmationHeader.ConfirmID)
	assert.Equal(t, "PO-123", decoded.Request.ConfirmationRequest.OrderReference.OrderID)
	assert.Equal(t, "ConfirmationRequest", decoded.Request.PayloadType())
}

func TestSerializeNilDocument(t *testing.T) {
	s := NewSerializer()
	out, err := s.Serialize(nil)
	assert.Error(t, err)
	assert.Nil(t, out)
}

func TestDeserializeEmptyInput(t *testing.T) {
	s := NewSerializer()
	doc, err := s.Deserialize(nil)
	assert.Error(t, err)
	assert.Nil(t, doc)
}

func TestDeserializeInvalidXML(t *testing.T) {
	s := NewSerializer()
	doc, err := s.Deserialize([]byte(`<cXML><Request></cXML>`))
	assert.Error(t, err)
	assert.Nil(t, doc)
}

func TestSerialize_EncodeError(t *testing.T) {
	s := NewSerializer()
	s.newBuffer = func() *bytes.Buffer { return &bytes.Buffer{} }
	s.newEncoder = func(w io.Writer) xmlEncoder {
		return &encoderStub{encodeErr: errors.New("encode failed")}
	}

	out, err := s.Serialize(&model.CXML{PayloadID: "x"})
	assert.EqualError(t, err, "encode failed")
	assert.Nil(t, out)
}

func TestSerialize_FlushError(t *testing.T) {
	s := NewSerializer()
	s.newBuffer = func() *bytes.Buffer { return &bytes.Buffer{} }
	s.newEncoder = func(w io.Writer) xmlEncoder {
		return &encoderStub{flushErr: errors.New("flush failed")}
	}

	out, err := s.Serialize(&model.CXML{PayloadID: "x"})
	assert.EqualError(t, err, "flush failed")
	assert.Nil(t, out)
}
