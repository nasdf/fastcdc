package fastcdc

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChunkerNext(t *testing.T) {
	var input bytes.Buffer
	var reader io.Reader

	reader = io.LimitReader(rand.Reader, 4096)
	reader = io.TeeReader(reader, &input)

	options, err := WithAverageSize(256, 2)
	require.NoError(t, err)

	chunker, err := NewChunker(reader, options)
	require.NoError(t, err)

	var output []byte
	for {
		chunk, err := chunker.Next()
		require.NoError(t, err)
		output = append(output, chunk...)

		if !chunker.HasNext() {
			break
		}

		assert.True(t, len(chunk) <= options.maxSize, "maxSize=%d len=%d", options.maxSize, len(chunk))
		assert.True(t, len(chunk) >= options.minSize, "minSize=%d len=%d", options.minSize, len(chunk))
	}

	expect, err := io.ReadAll(&input)
	require.NoError(t, err)
	assert.True(t, bytes.Equal(expect, output))
}
