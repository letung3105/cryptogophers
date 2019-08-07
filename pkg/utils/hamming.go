package utils

import (
	"fmt"
	"math/bits"

	"github.com/pkg/errors"
)

// HammingDistance computes the hamming distance between two equal length buffers
func HammingDistance(src, target []byte) (int, error) {
	if len(src) != len(target) {
		return -1, errors.New(fmt.Sprintf(
			"Mismatch length: got %d and %d", len(src), len(target),
		))
	}
	distance := 0
	for i := 0; i < len(src); i++ {
		distance += bits.OnesCount(uint(src[i] ^ target[i]))
	}
	return distance, nil
}
