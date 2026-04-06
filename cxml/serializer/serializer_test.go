package serializer

import (
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/stretchr/testify/assert"
)

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
