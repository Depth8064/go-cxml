package serializer

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/Depth8064/go-cxml/cxml/model"
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
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}
	if !strings.Contains(string(encoded), "<?xml") {
		t.Fatal("expected xml header")
	}
	if !strings.Contains(string(encoded), "OrderRequest") {
		t.Fatal("expected OrderRequest payload in output")
	}

	decoded, err := s.Deserialize(encoded)
	if err != nil {
		t.Fatalf("deserialize failed: %v", err)
	}
	if got, want := decoded.PayloadID, "12345"; got != want {
		t.Fatalf("unexpected payload id: got %q want %q", got, want)
	}
	if decoded.Request == nil {
		t.Fatal("expected request")
	}
	if decoded.Request.OrderRequest == nil {
		t.Fatal("expected order request")
	}
	if got, want := decoded.Request.OrderRequest.OrderRequestHeader.OrderID, "PO-1001"; got != want {
		t.Fatalf("unexpected order id: got %q want %q", got, want)
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
	if err != nil {
		t.Fatalf("deserialize failed: %v", err)
	}
	if got, want := decoded.PayloadID, "abc"; got != want {
		t.Fatalf("unexpected payload id: got %q want %q", got, want)
	}
	if decoded.Request == nil {
		t.Fatal("expected request")
	}
	if got, want := decoded.Request.OrderRequest.OrderRequestHeader.OrderID, "PO-99"; got != want {
		t.Fatalf("unexpected order id: got %q want %q", got, want)
	}
}

func TestSerializeAndDeserialize_PunchOutOrderMessage(t *testing.T) {
	doc := &model.CXML{
		PayloadID: "punch1",
		Request:   &model.Request{PunchOutOrderMessage: &model.PunchOutOrderMessage{PunchOutOrderMessageHeader: &model.PunchOutOrderMessageHeader{Operation: "create"}}},
	}

	s := NewSerializer()
	encoded, err := s.Serialize(doc)
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}
	if !strings.Contains(string(encoded), "PunchOutOrderMessage") {
		t.Fatal("expected punchout payload in output")
	}

	decoded, err := s.Deserialize(encoded)
	if err != nil {
		t.Fatalf("deserialize failed: %v", err)
	}
	if decoded.Request == nil || decoded.Request.PunchOutOrderMessage == nil {
		t.Fatal("expected punchout order message")
	}
	if got, want := decoded.Request.PunchOutOrderMessage.PunchOutOrderMessageHeader.Operation, "create"; got != want {
		t.Fatalf("unexpected operation: got %q want %q", got, want)
	}
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
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}
	if !strings.Contains(string(encoded), "ShipNoticeRequest") || !strings.Contains(string(encoded), "SHIP-42") {
		t.Fatal("expected ship notice content in output")
	}

	decoded, err := s.Deserialize(encoded)
	if err != nil {
		t.Fatalf("deserialize failed: %v", err)
	}
	if decoded.Request == nil || decoded.Request.ShipNoticeRequest == nil {
		t.Fatal("expected ship notice request")
	}
	if got, want := decoded.Request.ShipNoticeRequest.ShipNoticeHeader.ShipmentID, "SHIP-42"; got != want {
		t.Fatalf("unexpected shipment id: got %q want %q", got, want)
	}
	if got, want := decoded.Request.PayloadType(), "ShipNoticeRequest"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
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
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}
	if !strings.Contains(string(encoded), "InvoiceDetailRequest") || !strings.Contains(string(encoded), "INV-99") {
		t.Fatal("expected invoice detail content in output")
	}

	decoded, err := s.Deserialize(encoded)
	if err != nil {
		t.Fatalf("deserialize failed: %v", err)
	}
	if decoded.Request == nil || decoded.Request.InvoiceDetailRequest == nil {
		t.Fatal("expected invoice detail request")
	}
	if got, want := decoded.Request.InvoiceDetailRequest.InvoiceDetailRequestHeader.InvoiceID, "INV-99"; got != want {
		t.Fatalf("unexpected invoice id: got %q want %q", got, want)
	}
	if got, want := decoded.Request.PayloadType(), "InvoiceDetailRequest"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
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
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}
	if !strings.Contains(string(encoded), "ConfirmationRequest") || !strings.Contains(string(encoded), "CONF-1") {
		t.Fatal("expected confirmation content in output")
	}

	decoded, err := s.Deserialize(encoded)
	if err != nil {
		t.Fatalf("deserialize failed: %v", err)
	}
	if decoded.Request == nil || decoded.Request.ConfirmationRequest == nil {
		t.Fatal("expected confirmation request")
	}
	if got, want := decoded.Request.ConfirmationRequest.ConfirmationHeader.ConfirmID, "CONF-1"; got != want {
		t.Fatalf("unexpected confirm id: got %q want %q", got, want)
	}
	if got, want := decoded.Request.ConfirmationRequest.OrderReference.OrderID, "PO-123"; got != want {
		t.Fatalf("unexpected order reference: got %q want %q", got, want)
	}
	if got, want := decoded.Request.PayloadType(), "ConfirmationRequest"; got != want {
		t.Fatalf("unexpected payload type: got %q want %q", got, want)
	}
}

func TestSerializeNilDocument(t *testing.T) {
	s := NewSerializer()
	out, err := s.Serialize(nil)
	if err == nil {
		t.Fatal("expected serialize error")
	}
	if out != nil {
		t.Fatal("expected nil output on error")
	}
}

func TestDeserializeEmptyInput(t *testing.T) {
	s := NewSerializer()
	doc, err := s.Deserialize(nil)
	if err == nil {
		t.Fatal("expected deserialize error")
	}
	if doc != nil {
		t.Fatal("expected nil doc on error")
	}
}

func TestDeserializeInvalidXML(t *testing.T) {
	s := NewSerializer()
	doc, err := s.Deserialize([]byte(`<cXML><Request></cXML>`))
	if err == nil {
		t.Fatal("expected deserialize error")
	}
	if doc != nil {
		t.Fatal("expected nil doc on error")
	}
}

func TestSerialize_DocTypeVersionDerived(t *testing.T) {
	s := NewSerializer()

	// Version set — DOCTYPE should contain that version.
	out, err := s.Serialize(&model.CXML{Version: "1.2.069"})
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}
	if !strings.Contains(string(out), "1.2.069/cXML.dtd") {
		t.Fatalf("expected version 1.2.069 in DOCTYPE, got:\n%s", string(out))
	}
	if strings.Contains(string(out), "1.2.014") {
		t.Fatal("must not hardcode 1.2.014 when version is 1.2.069")
	}

	// No version — DOCTYPE should be omitted.
	outNoVer, err := s.Serialize(&model.CXML{})
	if err != nil {
		t.Fatalf("serialize failed: %v", err)
	}
	if strings.Contains(string(outNoVer), "<!DOCTYPE") {
		t.Fatal("expected no DOCTYPE when version is empty")
	}
}

func TestStripDoctype(t *testing.T) {
	// No DOCTYPE → unchanged.
	plain := `<cXML payloadID="x"/>`
	if got := stripDoctype(plain); got != plain {
		t.Fatalf("expected no change without DOCTYPE, got %q", got)
	}

	// Simple form (no internal subset).
	simple := `<?xml version="1.0"?><!DOCTYPE cXML SYSTEM "x.dtd"><cXML/>`
	stripped := stripDoctype(simple)
	if strings.Contains(stripped, "DOCTYPE") {
		t.Fatalf("expected DOCTYPE stripped, got %q", stripped)
	}
	if !strings.Contains(stripped, "<cXML/>") {
		t.Fatalf("expected element preserved after stripping, got %q", stripped)
	}

	// Internal subset form — DOCTYPE with [...].
	subset := `<?xml version="1.0"?><!DOCTYPE cXML [<!ENTITY foo "bar">]><cXML/>`
	strippedSubset := stripDoctype(subset)
	if strings.Contains(strippedSubset, "DOCTYPE") {
		t.Fatalf("expected internal-subset DOCTYPE stripped, got %q", strippedSubset)
	}
	if !strings.Contains(strippedSubset, "<cXML/>") {
		t.Fatalf("expected element preserved after internal-subset stripping, got %q", strippedSubset)
	}

	// Malformed DOCTYPE (internal subset opened but never closed) → returned unchanged.
	malformed := `<!DOCTYPE cXML [<!ENTITY foo "bar"`
	if got := stripDoctype(malformed); got != malformed {
		t.Fatalf("expected malformed DOCTYPE to be returned unchanged, got %q", got)
	}
}

func TestSerialize_EncodeError(t *testing.T) {
	s := NewSerializer()
	s.newBuffer = func() *bytes.Buffer { return &bytes.Buffer{} }
	s.newEncoder = func(w io.Writer) xmlEncoder {
		return &encoderStub{encodeErr: errors.New("encode failed")}
	}

	out, err := s.Serialize(&model.CXML{PayloadID: "x"})
	if err == nil {
		t.Fatal("expected encode error")
	}
	if got, want := err.Error(), "encode failed"; got != want {
		t.Fatalf("unexpected error: got %q want %q", got, want)
	}
	if out != nil {
		t.Fatal("expected nil output on error")
	}
}

func TestSerialize_FlushError(t *testing.T) {
	s := NewSerializer()
	s.newBuffer = func() *bytes.Buffer { return &bytes.Buffer{} }
	s.newEncoder = func(w io.Writer) xmlEncoder {
		return &encoderStub{flushErr: errors.New("flush failed")}
	}

	out, err := s.Serialize(&model.CXML{PayloadID: "x"})
	if err == nil {
		t.Fatal("expected flush error")
	}
	if got, want := err.Error(), "flush failed"; got != want {
		t.Fatalf("unexpected error: got %q want %q", got, want)
	}
	if out != nil {
		t.Fatal("expected nil output on error")
	}
}
