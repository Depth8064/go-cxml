package validation

import (
	"os"
	"testing"
)

func TestDTDValidator_LocalDTDSucceeds(t *testing.T) {
	// Local-only DTD check, enabled by env var
	os.Setenv("CXML_DTD_DIR", "testdata")
	defer os.Unsetenv("CXML_DTD_DIR")

	xml := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE cXML SYSTEM "http://xml.cxml.org/schemas/cXML/1.2.069/cXML.dtd">
<cXML payloadID="abc" timestamp="2026-03-24T12:34:56" version="1.2.069">
  <Header>
    <From>
      <Credential domain="DUNS">
        <Identity>Buyer</Identity>
      </Credential>
      <Identity>FromCompany</Identity>
    </From>
    <To>
      <Credential domain="DUNS">
        <Identity>Supplier</Identity>
      </Credential>
      <Identity>ToCompany</Identity>
    </To>
    <Sender>
      <Credential domain="DUNS">
        <Identity>Sender</Identity>
        <SharedSecret>secret</SharedSecret>
      </Credential>
      <UserAgent>go-cxml-test</UserAgent>
    </Sender>
  </Header>
  <Request>
    <OrderRequest>
      <OrderRequestHeader orderID="PO-123" orderDate="2026-03-24">
        <Total>
          <Money currency="USD">100.00</Money>
        </Total>
      </OrderRequestHeader>
    </OrderRequest>
  </Request>
</cXML>`)

	validator := NewDTDValidator()
	err := validator.Validate(xml)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
