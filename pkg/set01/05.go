package set01

import (
	"encoding/hex"

	"github.com/letung3105/cryptogophers/pkg/crypts"
)

// RepeatingXORHex computes repeating xor combination from src and key
// then hex encode the result
func RepeatingXORHex(src, key []byte) []byte {
	dst := crypts.RepeatingXOR(src, key)
	dstHex := make([]byte, hex.EncodedLen(len(dst)))
	n := hex.Encode(dstHex, dst)
	return dstHex[:n]
}
