# go-cxml

Pure Go implementation of the cXML (Commerce XML) standard. Provides parsing, building, serialization, and processing of cXML documents with zero external runtime dependencies.

## Build & Test

```powershell
go test ./cxml/...
go vet ./cxml/...
```

## Architecture

```
Input XML → Serializer (deserialize) → Authenticator → Processor (route to Handler) → Serializer (serialize) → Output XML
```

| Package | Purpose |
|---------|---------|
| `cxml/endpoint/` | Main entry point — orchestrates validate → auth → process → serialize |
| `cxml/model/` | cXML data structures (CXML, Header, Request, Response, OrderRequest, etc.) |
| `cxml/builder/` | Fluent builders for cXML documents (OrderRequest, OrderChange) |
| `cxml/serializer/` | XML marshal/unmarshal using `encoding/xml` |
| `cxml/processor/` | Routes documents to registered handlers by payload type |
| `cxml/handler/` | `Handler` interface + `Registry` for pluggable business logic |
| `cxml/auth/` | `Authenticator` interface, SharedSecret implementation |
| `cxml/credential/` | `CredentialRepository` interface for credential storage/validation |
| `cxml/validation/` | DTD validation |
| `cxml/document/` | `DocumentRegistry` for tracking processed documents |

## Conventions

- **Interface-driven**: All major components (`Handler`, `Authenticator`, `CredentialRepository`, `DocumentRegistry`) are interfaces
- **No global state**: Dependencies are passed explicitly
- **Nil-safe**: Check `nil` before use throughout
- **Error handling**: Return `error`, never panic in core logic
- **Testing**: Use `testify` for assertions; stub handlers implement `Handler` interface
- **XML tags**: Attributes use `xml:"name,attr"`, elements use `xml:"Name"`

## Consumer

This library is imported by the [Ariba-to-JobBOSS Connector](../Aribra%20to%20JobBOSS%20Connector/) as `github.com/mstrhakr/go-cxml`. The connector registers custom `Handler` implementations (e.g., `orderRequestHandler`) via the `handler.Registry`.
