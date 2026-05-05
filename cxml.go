package cxml

import internal "github.com/Depth8064/go-cxml/internal/cxml"

// Endpoint is the primary entry point for simple cXML serialization.
type Endpoint = internal.Endpoint

func NewEndpoint() *Endpoint { return internal.NewEndpoint() }
