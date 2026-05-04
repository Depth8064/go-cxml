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

	// Optional DOCTYPE maybe derived from c.Version or DTD not attached yet.
	if strings.TrimSpace(c.Version) != "" {
		buf.WriteString("<!DOCTYPE cXML SYSTEM \"http://xml.cxml.org/schemas/cXML/1.2.014/cXML.dtd\">\n")
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

func (s *Serializer) Deserialize(data []byte) (*model.CXML, error) {
	input := bytes.TrimSpace(data)
	if len(input) == 0 {
		return nil, errors.New("cxml: empty input")
	}

	// Strip DOCTYPE if present (basic approach)
	str := string(input)
	if idx := strings.Index(str, "<!DOCTYPE"); idx >= 0 {
		start := idx
		end := strings.Index(str[idx:], ">")
		if end > 0 {
			str = str[:start] + str[idx+end+1:]
		}
	}

	var doc model.CXML
	if err := xml.Unmarshal([]byte(str), &doc); err != nil {
		return nil, err
	}

	return &doc, nil
}
