package model

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestMoney_UnmarshalXML(t *testing.T) {
	t.Run("parses and trims value", func(t *testing.T) {
		var m Money
		err := xml.Unmarshal([]byte(`<Money currency="USD"> 123.45 </Money>`), &m)
		if err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		if got, want := m.Currency, "USD"; got != want {
			t.Fatalf("unexpected currency: got %q want %q", got, want)
		}
		if got, want := m.Value, "123.45"; got != want {
			t.Fatalf("unexpected value: got %q want %q", got, want)
		}
		if got, want := m.Amount, 123.45; got != want {
			t.Fatalf("unexpected amount: got %v want %v", got, want)
		}
	})

	t.Run("keeps zero amount on non-numeric value", func(t *testing.T) {
		var m Money
		err := xml.Unmarshal([]byte(`<Money currency="USD">abc</Money>`), &m)
		if err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		if got, want := m.Value, "abc"; got != want {
			t.Fatalf("unexpected value: got %q want %q", got, want)
		}
		if got, want := m.Amount, 0.0; got != want {
			t.Fatalf("unexpected amount: got %v want %v", got, want)
		}
	})

	t.Run("returns decode error on malformed XML", func(t *testing.T) {
		decoder := xml.NewDecoder(strings.NewReader(`<Money currency="USD">123`))
		tok, err := decoder.Token()
		if err != nil {
			t.Fatalf("unexpected token error: %v", err)
		}

		start, ok := tok.(xml.StartElement)
		if !ok {
			t.Fatal("expected start element token")
		}

		var m Money
		err = m.UnmarshalXML(decoder, start)
		if err == nil {
			t.Fatal("expected decode error")
		}
	})
}

func TestMoney_MarshalXML(t *testing.T) {
	t.Run("uses explicit value", func(t *testing.T) {
		m := Money{Currency: "USD", Value: "99.99", Amount: 100}
		out, err := xml.Marshal(m)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		if !strings.Contains(string(out), `currency="USD"`) {
			t.Fatal("expected USD currency attribute")
		}
		if !strings.Contains(string(out), `>99.99<`) {
			t.Fatal("expected explicit money value")
		}
	})

	t.Run("falls back to amount", func(t *testing.T) {
		m := Money{Currency: "EUR", Amount: 42.5}
		out, err := xml.Marshal(m)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		if !strings.Contains(string(out), `currency="EUR"`) {
			t.Fatal("expected EUR currency attribute")
		}
		if !strings.Contains(string(out), `>42.5<`) {
			t.Fatal("expected fallback money amount")
		}
	})
}

func TestCXML_GetPayloadTypeAndFlags(t *testing.T) {
	if got, want := (&CXML{}).GetPayloadType(), ""; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}

	reqDoc := &CXML{Request: &Request{OrderRequest: &OrderRequest{}}}
	if got, want := reqDoc.GetPayloadType(), "OrderRequest"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
	if !reqDoc.IsRequest() || reqDoc.IsResponse() || reqDoc.IsMessage() {
		t.Fatal("unexpected request/response/message flags for request doc")
	}

	respDoc := &CXML{Response: &Response{Status: &Status{Code: "200"}}}
	if got, want := respDoc.GetPayloadType(), "Status"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
	if respDoc.IsRequest() || !respDoc.IsResponse() || respDoc.IsMessage() {
		t.Fatal("unexpected request/response/message flags for response doc")
	}

	msgDoc := &CXML{Message: &Message{Subject: "hello"}}
	if got, want := msgDoc.GetPayloadType(), "Message"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
	if msgDoc.IsRequest() || msgDoc.IsResponse() || !msgDoc.IsMessage() {
		t.Fatal("unexpected request/response/message flags for message doc")
	}
}

func TestPrimaryCredential(t *testing.T) {
	cred := &Credential{Domain: "NetworkID", Identity: "buyer", SharedSecret: "secret"}

	if (*From)(nil).PrimaryCredential() != nil || (*To)(nil).PrimaryCredential() != nil || (*Sender)(nil).PrimaryCredential() != nil {
		t.Fatal("expected nil credentials for nil receivers")
	}

	if (&From{Credential: cred}).PrimaryCredential() != cred {
		t.Fatal("expected from primary credential")
	}
	if (&To{Credential: cred}).PrimaryCredential() != cred {
		t.Fatal("expected to primary credential")
	}
	if (&Sender{Credential: cred}).PrimaryCredential() != cred {
		t.Fatal("expected sender primary credential")
	}
}

func TestMessage_PayloadType(t *testing.T) {
	if got, want := (*Message)(nil).PayloadType(), ""; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
	if got, want := (&Message{Subject: "x"}).PayloadType(), "Message"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
	if got, want := (&Message{Payload: &PayloadWrapper{Content: "<x/>"}}).PayloadType(), "Payload"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
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
			if got := tt.r.PayloadType(); got != tt.want {
				t.Fatalf("unexpected payload type: got %q want %q", got, tt.want)
			}
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
			if got := tt.r.PayloadType(); got != tt.want {
				t.Fatalf("unexpected payload type: got %q want %q", got, tt.want)
			}
		})
	}
}

func TestRequestPayloadName_Implementations(t *testing.T) {
	if got, want := (&PunchOutOrderMessage{}).RequestPayloadName(), "PunchOutOrderMessage"; got != want {
		t.Fatalf("unexpected payload name: got %q want %q", got, want)
	}
	if got, want := (&OrderRequest{}).RequestPayloadName(), "OrderRequest"; got != want {
		t.Fatalf("unexpected payload name: got %q want %q", got, want)
	}
}

func TestMessagePayloadWrapper_ContentRoundTrip(t *testing.T) {
	m := &Message{Payload: &PayloadWrapper{Content: `<Foo attr="1"/>`}}
	out, err := xml.Marshal(m)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if !(strings.Contains(string(out), `<Payload><Foo attr="1"></Foo></Payload>`) || strings.Contains(string(out), `<Payload><Foo attr="1"/></Payload>`)) {
		t.Fatalf("unexpected payload wrapper output: %s", string(out))
	}
}
