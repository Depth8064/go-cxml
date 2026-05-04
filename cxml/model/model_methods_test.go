package model

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_UnmarshalXML(t *testing.T) {
	t.Run("parses and trims value", func(t *testing.T) {
		var m Money
		err := xml.Unmarshal([]byte(`<Money currency="USD"> 123.45 </Money>`), &m)
		assert.NoError(t, err)
		assert.Equal(t, "USD", m.Currency)
		assert.Equal(t, "123.45", m.Value)
		assert.Equal(t, 123.45, m.Amount)
	})

	t.Run("keeps zero amount on non-numeric value", func(t *testing.T) {
		var m Money
		err := xml.Unmarshal([]byte(`<Money currency="USD">abc</Money>`), &m)
		assert.NoError(t, err)
		assert.Equal(t, "abc", m.Value)
		assert.Equal(t, 0.0, m.Amount)
	})

	t.Run("returns decode error on malformed XML", func(t *testing.T) {
		decoder := xml.NewDecoder(strings.NewReader(`<Money currency="USD">123`))
		tok, err := decoder.Token()
		assert.NoError(t, err)

		start, ok := tok.(xml.StartElement)
		assert.True(t, ok)

		var m Money
		err = m.UnmarshalXML(decoder, start)
		assert.Error(t, err)
	})
}

func TestMoney_MarshalXML(t *testing.T) {
	t.Run("uses explicit value", func(t *testing.T) {
		m := Money{Currency: "USD", Value: "99.99", Amount: 100}
		out, err := xml.Marshal(m)
		assert.NoError(t, err)
		assert.Contains(t, string(out), `currency="USD"`)
		assert.Contains(t, string(out), `>99.99<`)
	})

	t.Run("falls back to amount", func(t *testing.T) {
		m := Money{Currency: "EUR", Amount: 42.5}
		out, err := xml.Marshal(m)
		assert.NoError(t, err)
		assert.Contains(t, string(out), `currency="EUR"`)
		assert.Contains(t, string(out), `>42.5<`)
	})
}

func TestCXML_GetPayloadTypeAndFlags(t *testing.T) {
	assert.Equal(t, "", (&CXML{}).GetPayloadType())

	reqDoc := &CXML{Request: &Request{OrderRequest: &OrderRequest{}}}
	assert.Equal(t, "OrderRequest", reqDoc.GetPayloadType())
	assert.True(t, reqDoc.IsRequest())
	assert.False(t, reqDoc.IsResponse())
	assert.False(t, reqDoc.IsMessage())

	respDoc := &CXML{Response: &Response{Status: &Status{Code: "200"}}}
	assert.Equal(t, "Status", respDoc.GetPayloadType())
	assert.False(t, respDoc.IsRequest())
	assert.True(t, respDoc.IsResponse())
	assert.False(t, respDoc.IsMessage())

	msgDoc := &CXML{Message: &Message{Subject: "hello"}}
	assert.Equal(t, "Message", msgDoc.GetPayloadType())
	assert.False(t, msgDoc.IsRequest())
	assert.False(t, msgDoc.IsResponse())
	assert.True(t, msgDoc.IsMessage())
}

func TestPrimaryCredential(t *testing.T) {
	cred := &Credential{Domain: "NetworkID", Identity: "buyer", SharedSecret: "secret"}

	assert.Nil(t, (*From)(nil).PrimaryCredential())
	assert.Nil(t, (*To)(nil).PrimaryCredential())
	assert.Nil(t, (*Sender)(nil).PrimaryCredential())

	assert.Equal(t, cred, (&From{Credential: cred}).PrimaryCredential())
	assert.Equal(t, cred, (&To{Credential: cred}).PrimaryCredential())
	assert.Equal(t, cred, (&Sender{Credential: cred}).PrimaryCredential())
}

func TestMessage_PayloadType(t *testing.T) {
	assert.Equal(t, "", (*Message)(nil).PayloadType())
	assert.Equal(t, "Message", (&Message{Subject: "x"}).PayloadType())
	assert.Equal(t, "Payload", (&Message{Payload: &PayloadWrapper{Content: "<x/>"}}).PayloadType())
}

func TestRequest_PayloadType_AllBranches(t *testing.T) {
	tests := []struct {
		name string
		r    *Request
		want string
	}{
		{name: "order", r: &Request{OrderRequest: &OrderRequest{}}, want: "OrderRequest"},
		{name: "order change", r: &Request{OrderChangeRequest: &OrderChangeRequest{}}, want: "OrderChangeRequest"},
		{name: "confirmation", r: &Request{ConfirmationRequest: &ConfirmationRequest{}}, want: "ConfirmationRequest"},
		{name: "profile", r: &Request{ProfileRequest: &ProfileRequest{}}, want: "ProfileRequest"},
		{name: "status update", r: &Request{StatusUpdateRequest: &StatusUpdateRequest{}}, want: "StatusUpdateRequest"},
		{name: "punchout setup", r: &Request{PunchOutSetupRequest: &PunchOutSetupRequest{}}, want: "PunchOutSetupRequest"},
		{name: "punchout order", r: &Request{PunchOutOrderMessage: &PunchOutOrderMessage{}}, want: "PunchOutOrderMessage"},
		{name: "receiving advice", r: &Request{ReceivingAdviceRequest: &ReceivingAdviceRequest{}}, want: "ReceivingAdviceRequest"},
		{name: "ship notice", r: &Request{ShipNoticeRequest: &ShipNoticeRequest{}}, want: "ShipNoticeRequest"},
		{name: "invoice detail", r: &Request{InvoiceDetailRequest: &InvoiceDetailRequest{}}, want: "InvoiceDetailRequest"},
		{name: "none", r: &Request{}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.r.PayloadType())
		})
	}
}

func TestResponse_PayloadType_AllBranches(t *testing.T) {
	tests := []struct {
		name string
		r    *Response
		want string
	}{
		{name: "order response", r: &Response{OrderResponse: &OrderResponse{}}, want: "OrderResponse"},
		{name: "profile response", r: &Response{ProfileResponse: &ProfileResponse{}}, want: "ProfileResponse"},
		{name: "punchout setup response", r: &Response{PunchOutSetupResponse: &PunchOutSetupResponse{}}, want: "PunchOutSetupResponse"},
		{name: "status", r: &Response{Status: &Status{Code: "200"}}, want: "Status"},
		{name: "none", r: &Response{}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.r.PayloadType())
		})
	}
}

func TestRequestPayloadName_Implementations(t *testing.T) {
	assert.Equal(t, "PunchOutOrderMessage", (&PunchOutOrderMessage{}).RequestPayloadName())
	assert.Equal(t, "OrderRequest", (&OrderRequest{}).RequestPayloadName())
}

func TestMessagePayloadWrapper_ContentRoundTrip(t *testing.T) {
	m := &Message{Payload: &PayloadWrapper{Content: `<Foo attr="1"/>`}}
	out, err := xml.Marshal(m)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(string(out), `<Payload><Foo attr="1"></Foo></Payload>`) || strings.Contains(string(out), `<Payload><Foo attr="1"/></Payload>`))
}
