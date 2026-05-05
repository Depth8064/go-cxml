# go-cxml

[![CI](https://github.com/Depth8064/go-cxml/actions/workflows/ci.yml/badge.svg)](https://github.com/Depth8064/go-cxml/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Depth8064/go-cxml)](go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/Depth8064/go-cxml)](https://goreportcard.com/report/github.com/Depth8064/go-cxml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/Depth8064/go-cxml.svg)](https://pkg.go.dev/github.com/Depth8064/go-cxml)
[![License](https://img.shields.io/github/license/Depth8064/go-cxml)](LICENSE)

Pure Go implementation of cXML (Commerce XML) with zero external runtime dependencies.

## Highlights

- Parse and serialize cXML using the standard library `encoding/xml`
- Build outbound cXML documents with fluent builders
- Route inbound payloads to pluggable handlers by payload type
- Shared-secret authentication support
- Optional DTD validation and document registry hooks

## Module And Import Paths

This module is:

- `github.com/Depth8064/go-cxml`

For the simple serialize/deserialize entry point, prefer importing the module root:

- `github.com/Depth8064/go-cxml`

Granular packages remain available under:

- `github.com/Depth8064/go-cxml/cxml/...`

Example:

```go
import (
	"github.com/Depth8064/go-cxml"
	"github.com/Depth8064/go-cxml/cxml/model"
)
```

## Installation

```bash
go get github.com/Depth8064/go-cxml
```

## Quick Start

### Serialize / Deserialize

```go
package main

import (
	"fmt"

	gcxml "github.com/Depth8064/go-cxml"
	"github.com/Depth8064/go-cxml/cxml/model"
)

func main() {
	ep := gcxml.NewEndpoint()

	doc := &model.CXML{
		PayloadID: "example-1",
		Version:   "1.2.014",
		Request:   &model.Request{OrderRequest: &model.OrderRequest{}},
	}

	out, err := ep.Serialize(doc)
	if err != nil {
		panic(err)
	}

	parsed, err := ep.Deserialize(out)
	if err != nil {
		panic(err)
	}

	fmt.Println(parsed.PayloadID)
}
```

### Processing Pipeline Endpoint

The full-pipeline endpoint in `cxml/endpoint` runs:

1. Validate input (DTD)
2. Deserialize XML into model
3. Authenticate sender credentials
4. Route payload to registered handler
5. Serialize response cXML

Use `cxml/endpoint` when receiving inbound cXML. Use the module root package when you only need simple serialization.

This is a breaking cleanup: the old entry-point import path `github.com/Depth8064/go-cxml/cxml` has been removed. Import `github.com/Depth8064/go-cxml` instead.

```go
package main

import (
	"fmt"

	"github.com/Depth8064/go-cxml/cxml/builder"
	"github.com/Depth8064/go-cxml/cxml/credential"
	"github.com/Depth8064/go-cxml/cxml/endpoint"
	"github.com/Depth8064/go-cxml/cxml/handler"
	"github.com/Depth8064/go-cxml/cxml/model"
	"github.com/Depth8064/go-cxml/cxml/processor"
)

type orderHandler struct{}

func (h *orderHandler) Handle(doc *model.CXML) (*model.CXML, error) {
	// Business logic here.
	return builder.New().
		PayloadID(doc.PayloadID + "-resp").
		Version("1.2.069").
		Response(&model.Response{Status: &model.Status{Code: "200", Text: "OK"}}).
		Build(), nil
}

func main() {
	reg := handler.NewRegistry()
	reg.Register("OrderRequest", &orderHandler{})

	proc := processor.NewProcessor(reg)

	creds := credential.NewRegistry([]credential.Entry{
		{Domain: "NetworkID", Identity: "buyer", SharedSecret: "secret"},
	})

	ep := endpoint.NewEndpoint(proc, nil, creds)

	// Receive inbound cXML bytes (from HTTP POST body, etc.)
	var inboundXML []byte // populate from request

	outXML, err := ep.Process(inboundXML)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %s\n", outXML)
}
```

### Building Outbound cXML

Use the typed builders in `cxml/builder` to construct outbound documents:

```go
doc := builder.NewConfirmationRequestBuilder().
	PayloadID("conf-001").
	Version("1.2.069").
	From(&model.Party{Identity: "supplier"}).
	To(&model.Party{Identity: "buyer"}).
	Sender(&model.Sender{UserAgent: "my-system"}).
	Request(&model.ConfirmationRequest{
		ConfirmationHeader: &model.ConfirmationHeader{
			ConfirmID: "CONF-001",
			Operation: "accept",
		},
		OrderReference: &model.OrderReference{OrderID: "PO-100"},
	}).
	Build()
```

Available builders:

| Builder | Payload type |
|---|---|
| `NewOrderRequestBuilder` | `OrderRequest` |
| `NewOrderChangeBuilder` | `OrderChangeRequest` |
| `NewConfirmationRequestBuilder` | `ConfirmationRequest` |
| `NewPunchOutSetupBuilder` | `PunchOutSetupRequest` |
| `NewStatusUpdateBuilder` | `StatusUpdateRequest` |
| `NewProfileRequestBuilder` | `ProfileRequest` |
| `NewReceivingAdviceBuilder` | `ReceivingAdviceRequest` |
| `NewShipNoticeBuilder` | `ShipNoticeRequest` |
| `NewInvoiceDetailBuilder` | `InvoiceDetailRequest` |

## Build And Test

```bash
go test ./cxml/...
go test ./cxml/... -coverprofile=coverage
go tool cover -func=coverage
go vet ./cxml/...
```

## Project Structure

- `cxml/endpoint`: orchestration entry point for processing
- `cxml/model`: cXML data model types
- `cxml/builder`: fluent builders for outbound cXML
- `cxml/serializer`: XML marshal/unmarshal helpers
- `cxml/processor`: payload-to-handler routing
- `cxml/handler`: handler interfaces and registry
- `cxml/auth`: authentication interfaces and implementation
- `cxml/credential`: credential repository interfaces/registry
- `cxml/validation`: DTD validation
- `cxml/document`: document registry interfaces/implementation

## Compatibility

- Language: Go 1.25+
- Runtime dependencies: none (standard library only)

## Security

For vulnerability reporting instructions, see [SECURITY.md](SECURITY.md).

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE).
