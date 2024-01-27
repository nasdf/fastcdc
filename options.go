package fastcdc

import (
	"errors"
	"math"
)

var (
	ErrInvalidSize = errors.New("average size must be between 256B and 4MB")
	ErrInvalidNorm = errors.New("normalization must be between 0 and 3")
)

// mask values are uniform distributions of 1 bits for varying sized chunks.
var mask = [26]uint64{
	0,                  // padding
	0,                  // padding
	0,                  // padding
	0,                  // padding
	0,                  // padding
	0x0000000001804110, // unused except for NC 3
	0x0000000001803110, // 64B
	0x0000000018035100, // 128B
	0x0000001800035300, // 256B
	0x0000019000353000, // 512B
	0x0000590003530000, // 1KB
	0x0000d90003530000, // 2KB
	0x0000d90103530000, // 4KB
	0x0000d90303530000, // 8KB
	0x0000d90313530000, // 16KB
	0x0000d90f03530000, // 32KB
	0x0000d90303537000, // 64KB
	0x0000d90703537000, // 128KB
	0x0000d90707537000, // 256KB
	0x0000d91707537000, // 512KB
	0x0000d91747537000, // 1MB
	0x0000d91767537000, // 2MB
	0x0000d93767537000, // 4MB
	0x0000d93777537000, // 8MB
	0x0000d93777577000, // 16MB
	0x0000db3777577000, // unused except for NC 3
}

// Options is used to control the behavior of the FastCDC algorithm.
type Options struct {
	minSize int
	maxSize int
	avgSize int
	maskS   uint64
	maskL   uint64
}

// WithAverageSize returns options that will attempt to create chunks
// of the given average size.
func WithAverageSize(avgSize, norm int) (*Options, error) {
	if avgSize < 256 || avgSize > (4<<20) {
		return nil, ErrInvalidSize
	}
	if norm < 0 || norm > 3 {
		return nil, ErrInvalidNorm
	}

	maxSize := avgSize * 4
	minSize := avgSize / 4

	bits := int(math.Log2(float64(avgSize)))
	maskS := mask[bits+norm]
	maskL := mask[bits-norm]

	return &Options{
		minSize: minSize,
		maxSize: maxSize,
		avgSize: avgSize,
		maskS:   maskS,
		maskL:   maskL,
	}, nil
}
