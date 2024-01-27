package fastcdc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionsSizeTooSmall(t *testing.T) {
	options, err := WithAverageSize(128, 0)
	assert.ErrorIs(t, err, ErrInvalidSize)
	assert.Nil(t, options)
}

func TestOptionsSizeTooLarge(t *testing.T) {
	options, err := WithAverageSize(5<<20, 0)
	assert.ErrorIs(t, err, ErrInvalidSize)
	assert.Nil(t, options)
}

func TestOptionsNormNegative(t *testing.T) {
	options, err := WithAverageSize(256, -1)
	assert.ErrorIs(t, err, ErrInvalidNorm)
	assert.Nil(t, options)
}

func TestOptionsNormTooLarge(t *testing.T) {
	options, err := WithAverageSize(256, 4)
	assert.ErrorIs(t, err, ErrInvalidNorm)
	assert.Nil(t, options)
}

func TestOptionsSizeConfig(t *testing.T) {
	options, err := WithAverageSize(256, 2)
	require.NoError(t, err)
	assert.Equal(t, options.minSize, 64)
	assert.Equal(t, options.maxSize, 1024)
	assert.Equal(t, options.avgSize, 256)
}
