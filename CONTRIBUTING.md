# Contributing

Thanks for helping improve go-cxml.

## Ground Rules

1. Keep changes focused and minimal.
2. Avoid breaking public APIs unless discussed first.
3. Add or update tests with every behavior change.
4. Preserve zero external runtime dependency goals in core packages.
5. Prefer interface-driven, nil-safe, explicit error handling.

## Development Setup

```bash
go test ./cxml/...
go vet ./cxml/...
```

## Coding Rules

1. Return errors; do not panic in library code.
2. Keep XML tags consistent with cXML and existing model style.
3. Do not introduce global mutable state.
4. Keep package boundaries clear (model, serializer, processor, endpoint, etc.).

## Pull Request Rules

1. Explain what changed and why.
2. Link related issue(s).
3. Include tests for new or changed behavior.
4. Keep PRs reviewable; split large unrelated work.

## Commit Message Guidance

Use short, descriptive commit subjects in imperative mood.

Examples:

- `Add tests for endpoint auth failures`
- `Refactor serializer encoder injection`
