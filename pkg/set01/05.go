package set01

import (
	"encoding/hex"

	"github.com/letung3105/cryptogophers/pkg/crypts"
)

// RepeatingXORHex computes repeating xor combination from src and key
// then hex encode the result
func RepeatingXORHex(src, key []byte) []byte {
	cipher := crypts.RepeatingXOR(src, key)
	dst := make([]byte, hex.EncodedLen(len(cipher)))
	n := hex.Encode(dst, cipher)
	return dst[:n]
}
