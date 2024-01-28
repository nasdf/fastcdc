package fastcdc

//go:generate go run ./cmd/gear

import (
	_ "embed"
	"encoding/binary"
)

var (
	//go:embed gear
	gear []byte
	// gearN is the normal gear table
	gearN [256]uint64
	// gearL is the left shifted gear table
	gearL [256]uint64
)

func init() {
	// pre-compute gear tables
	for i := 0; i < 256; i++ {
		gearN[i] = binary.LittleEndian.Uint64(gear[i*8:])
		gearL[i] = gearN[i] << 1
	}
}

// Boundary returns the next chunk boundary for the given bytes.
func Boundary(src []byte, options *Options) int {
	size := min(len(src), options.maxSize)
	if size <= options.minSize {
		return size
	}

	hash := uint64(0)
	norm := min(options.avgSize, size)

	for i := options.minSize; i < (norm / 2); i++ {
		hash = (hash << 2) + gearL[src[i*2]]
		if (hash & (options.maskS << 1)) == 0 {
			return i * 2
		}
		hash = hash + gearN[src[i*2+1]]
		if (hash & options.maskS) == 0 {
			return (i * 2) + 1
		}
	}

	for i := (norm / 2); i < (size / 2); i++ {
		hash = (hash << 2) + gearL[src[i*2]]
		if (hash & (options.maskL << 1)) == 0 {
			return i * 2
		}
		hash = hash + gearN[src[i*2+1]]
		if (hash & options.maskL) == 0 {
			return (i * 2) + 1
		}
	}

	return size
}
