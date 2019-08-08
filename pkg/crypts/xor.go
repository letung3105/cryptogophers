package crypts

import (
	"github.com/pkg/errors"
)

// FixedXOR produces fixed xor of each byte in two equal length buffers
func FixedXOR(src, key []byte) ([]byte, error) {
	if len(src) != len(key) {
		return nil, errors.Errorf("length mismatch: got %d and %d", len(src), len(key))
	}

	dst := make([]byte, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ key[i]
	}
	return dst, nil
}

// SingleXOR produces the xor combination of each byte in the buffer against a single byte
func SingleXOR(src []byte, key byte) []byte {
	dst := make([]byte, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ key
	}
	return dst
}

// RepeatingXOR implements repeating key xor
func RepeatingXOR(src, key []byte) []byte {
	if len(key) == 0 {
		return src
	}
	dst := make([]byte, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ key[i%len(key)]
	}
	return dst
}
