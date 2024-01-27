# FastCDC

This package implements the Fast Content Defined Chunking algorithm along with some utilities to make it simple to use.

## Usage

```go
import (
    "github.com/nasdf/fastcdc"
)

options, _ := fastcdc.WithAverageSize(256, 2)
chunker, _ := fastcdc.NewChunker(reader, options)

for chunker.HasNext() {
    chunk, _ := chunker.Next()
}
```
