# Security Toolkit

A lightweight Go command-line toolkit built from algorithmic exercises in my Go training and extended into reusable text and number-analysis utilities.

## Why this project?

The original exercises focused on algorithms, strings, slices, validation, and integer factorization. This project turns those foundations into a small CLI that demonstrates how individual algorithms can be organized into reusable functions and tested as software components.

## Features

- **Prime factorization** — decomposes integers into ascending prime factors.
- **Hidden sequence detection** — checks whether one string appears in another while preserving character order.
- **Character intersection** — finds unique characters shared by two strings.
- **Character union** — combines unique characters while preserving first appearance.
- **Alternating text selection** — saves and skips fixed-size chunks.

## Quick start

```bash
git clone <your-repository-url>
cd security-toolkit
go test ./...
go run . factor 225225
go run . hidden faya fgvvfdxcacpolhyghbreda
go run . inter padinton paqefwtdjetyiytjneytjoeyj
go run . union zpadinton paqefwtdjetyiytjneytjoeyj
go run . mask 123456789 3
```

## Example

```text
$ go run . factor 42
2*3*7

$ go run . hidden abc 2altrb53c.sse
true

$ go run . inter padinton paqefwtdjetyiytjneytjoeyj
padinto

$ go run . union zpadinton paqefwtdjetyiytjneytjoeyj
zpadintoqefwjy

$ go run . mask 123456789 3
123789
```

## Engineering practices demonstrated

- Go functions and modular logic
- Command-line argument handling
- Input validation
- Slice and string processing
- Algorithmic problem solving
- Unit testing with Go's standard testing package
- Clean separation between CLI behavior and reusable functions

## Security relevance

The toolkit is intentionally lightweight; it is **not a cryptographic library or security scanner**. Its security relevance comes from the underlying concepts: factorization, pattern/sequence detection, input validation, character analysis, and controlled data masking.

These components can serve as building blocks for later security-focused tooling such as log analysis, sensitive-data detection, and input-validation systems.

## Roadmap

- Add benchmarks for the algorithms.
- Add UTF-8/rune-aware text handling.
- Add structured error handling.
- Add fuzz tests for string-processing functions.
- Add CI with `go test ./...`.
- Add security-oriented log-analysis features.

## License

MIT
