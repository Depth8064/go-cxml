# go-cxml

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

Most consumers import packages under:

- `github.com/Depth8064/go-cxml/cxml/...`

Example:

```go
import (
	"github.com/Depth8064/go-cxml/cxml"
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

	gcxml "github.com/Depth8064/go-cxml/cxml"
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

The processing endpoint in `cxml/endpoint` runs this flow:

1. Validate input (DTD)
2. Deserialize XML into model
3. Authenticate sender
4. Route payload to registered handler
5. Serialize response cXML

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
