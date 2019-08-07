package ciphers

import (
	"fmt"

	"github.com/pkg/errors"
)

// FixedXOR produces fixed xor of each byte in two equal length buffers
func FixedXOR(src, target []byte) ([]byte, error) {
	if len(src) != len(target) {
		return nil,
			errors.New(fmt.Sprintf(
				"Length mismatch: got %d and %d",
				len(src), len(target),
			))
	}

	dst := make([]byte, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ target[i]
	}
	return dst, nil
}

// SingleByteXOR produces the xor combination of each byte in the buffer against a single byte
func SingleByteXOR(src []byte, target byte) []byte {
	dst := make([]byte, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ target
	}
	return dst
}
