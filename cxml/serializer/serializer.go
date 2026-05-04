package serializer

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"strings"

	"github.com/Depth8064/go-cxml/cxml/model"
)

type xmlEncoder interface {
	Encode(v any) error
	Flush() error
	Indent(prefix, indent string)
}

type Serializer struct {
	newBuffer  func() *bytes.Buffer
	newEncoder func(io.Writer) xmlEncoder
}

func NewSerializer() *Serializer {
	return &Serializer{
		newBuffer: func() *bytes.Buffer { return &bytes.Buffer{} },
		newEncoder: func(w io.Writer) xmlEncoder {
			return xml.NewEncoder(w)
		},
	}
}

func (s *Serializer) Serialize(c *model.CXML) ([]byte, error) {
	if c == nil {
		return nil, errors.New("cxml: nil document")
	}

	buf := s.newBuffer()
	buf.WriteString(xml.Header)

	if v := strings.TrimSpace(c.Version); v != "" {
		buf.WriteString("<!DOCTYPE cXML SYSTEM \"http://xml.cxml.org/schemas/cXML/" + v + "/cXML.dtd\">\n")
	}

	enc := s.newEncoder(buf)
	enc.Indent("", "  ")

	if err := enc.Encode(c); err != nil {
		return nil, err
	}
	if err := enc.Flush(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// stripDoctype removes a <!DOCTYPE ...> declaration from s, handling both
// the simple form (ends at '>') and the internal-subset form (ends at ']>').
func stripDoctype(s string) string {
	idx := strings.Index(s, "<!DOCTYPE")
	if idx < 0 {
		return s
	}
	rest := s[idx:]
	depth := 0
	for i := 0; i < len(rest); i++ {
		switch rest[i] {
		case '[':
			depth++
		case ']':
			if depth > 0 {
				depth--
			}
		case '>':
			if depth == 0 {
				return s[:idx] + rest[i+1:]
			}
		}
	}
	// Malformed DOCTYPE — return original unchanged.
	return s
}

func (s *Serializer) Deserialize(data []byte) (*model.CXML, error) {
	input := bytes.TrimSpace(data)
	if len(input) == 0 {
		return nil, errors.New("cxml: empty input")
	}

	str := stripDoctype(string(input))

	var doc model.CXML
	if err := xml.Unmarshal([]byte(str), &doc); err != nil {
		return nil, err
	}

	return &doc, nil
}
