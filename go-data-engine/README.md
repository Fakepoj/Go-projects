# Go Data Processing Engine

A tested Go toolkit built from core slice and string algorithms developed during Go programming training and extended into reusable data-processing components.

## Features

- Chunk integer data into batches
- Concatenate datasets without mutating inputs
- Interleave ordered datasets
- Normalize word endings
- CLI interface and unit tests

## Run

```bash
go test ./...
go run ./cmd/data-engine chunk 3 0 1 2 3 4 5 6 7
go run ./cmd/data-engine concat 1 2 3 -- 4 5 6
go run ./cmd/data-engine alternate 1 2 3 -- 4 5 6 7
go run ./cmd/data-engine normalize "First SMALL TesT"
```

## Engineering value

Demonstrates slice allocation, data transformation, input validation, package organization, CLI design, unit testing, and edge-case handling.

## Security / AI relevance

Chunking and transformation are common building blocks in log processing, batch workloads, feature preparation, and security-data pipelines. This project can later grow into a streaming security-log preprocessing component for AI security workflows.
