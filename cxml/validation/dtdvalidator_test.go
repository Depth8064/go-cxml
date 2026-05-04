package validation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDTDValidator_Valid(t *testing.T) {
	xml := []byte(`<?xml version="1.0"?>
<!DOCTYPE cXML SYSTEM "http://xml.cxml.org/schemas/cXML/1.2.014/cXML.dtd">
<cXML payloadID="abc"></cXML>`)
	v := NewDTDValidator()
	err := v.Validate(xml)
	assert.NoError(t, err)
}

func TestDTDValidator_MissingDoctype(t *testing.T) {
	xml := []byte(`<?xml version="1.0"?><cXML payloadID="abc"></cXML>`)
	v := NewDTDValidator()
	err := v.Validate(xml)
	assert.Error(t, err)
}

func TestDTDValidator_EmptyDocument(t *testing.T) {
	v := NewDTDValidator()
	err := v.Validate(nil)
	assert.EqualError(t, err, "validation: empty document")
}

func TestDTDValidator_MissingCXMLRoot(t *testing.T) {
	xml := []byte(`<?xml version="1.0"?><root></root>`)
	v := NewDTDValidator()
	err := v.Validate(xml)
	assert.EqualError(t, err, "validation: document does not contain cXML root")
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
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation: local DTD file not found")
}
