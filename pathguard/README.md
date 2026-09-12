# PathGuard

A Go path-validation engine built from an exact-jump algorithm developed during Go programming training.

## Features
- Exact-jump reachability validation
- Path reconstruction
- Boundary checking
- Cycle detection
- Unit tests
- CLI interface

## Run
```bash
go test ./...
go run ./cmd/pathguard 2 3 1 1 4
go run ./cmd/pathguard 3 2 1 0 4
```

Example:
```text
VALID: 0 -> 2 -> 3 -> 4
```

## Security relevance
The reachability and state-validation concepts can be applied to security workflows such as validating permitted state transitions, execution paths, and policy rules. This is a learning project, not a production authorization engine.

## Engineering concepts
Go packages, slices, maps, input validation, state tracking, edge-case handling, unit testing, and CLI design.
