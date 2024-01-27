package fastcdc

import (
	"bufio"
	"io"
)

// Chunker splits bytes into content defined chunks.
type Chunker struct {
	reader  *bufio.Reader
	options *Options
}

// NewChunker returns a new chunker that returns content defined chunks from the given reader.
func NewChunker(r io.Reader, options *Options) (*Chunker, error) {
	reader := bufio.NewReaderSize(r, options.maxSize)

	// make sure the buffer is filled
	_, err := reader.Peek(options.maxSize)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return &Chunker{
		reader:  reader,
		options: options,
	}, nil
}

// HasNext returns true if there are bytes remaining in the reader.
func (c *Chunker) HasNext() bool {
	return c.reader.Buffered() != 0
}

// Next returns the next chunk of bytes from the reader.
func (c *Chunker) Next() ([]byte, error) {
	src, err := c.reader.Peek(c.options.maxSize)
	if err != nil && err != io.EOF {
		return nil, err
	}
	chunk := make([]byte, Boundary(src, c.options))
	if _, err := io.ReadFull(c.reader, chunk); err != nil {
		return nil, err
	}
	return chunk, nil
}
