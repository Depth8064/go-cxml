package validation

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDTDValidator_Valid(t *testing.T) {
	xml := []byte(`<?xml version="1.0"?>
<!DOCTYPE cXML SYSTEM "http://xml.cxml.org/schemas/cXML/1.2.014/cXML.dtd">
<cXML payloadID="abc"></cXML>`)
	v := NewDTDValidator()
	err := v.Validate(xml)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDTDValidator_MissingDoctype(t *testing.T) {
	xml := []byte(`<?xml version="1.0"?><cXML payloadID="abc"></cXML>`)
	v := NewDTDValidator()
	err := v.Validate(xml)
	if err == nil {
		t.Fatal("expected error for missing doctype")
	}
}

func TestDTDValidator_EmptyDocument(t *testing.T) {
	v := NewDTDValidator()
	err := v.Validate(nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if got, want := err.Error(), "validation: empty document"; got != want {
		t.Fatalf("unexpected error: got %q want %q", got, want)
	}
}

func TestDTDValidator_MissingCXMLRoot(t *testing.T) {
	xml := []byte(`<?xml version="1.0"?><root></root>`)
	v := NewDTDValidator()
	err := v.Validate(xml)
	if err == nil {
		t.Fatal("expected error")
	}
	if got, want := err.Error(), "validation: document does not contain cXML root"; got != want {
		t.Fatalf("unexpected error: got %q want %q", got, want)
	}
}

func TestDTDValidator_LocalDTDFileMissing(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-dtd")
	os.Setenv("CXML_DTD_DIR", missing)
	defer os.Unsetenv("CXML_DTD_DIR")

	xml := []byte(`<?xml version="1.0"?>
<!DOCTYPE cXML SYSTEM "http://xml.cxml.org/schemas/cXML/1.2.014/cXML.dtd">
<cXML payloadID="abc"></cXML>`)

	v := NewDTDValidator()
	err := v.Validate(xml)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "validation: local DTD file not found") {
		t.Fatalf("unexpected error message: %v", err)
	}
}
