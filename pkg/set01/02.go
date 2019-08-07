package set01

import (
	"encoding/hex"

	"github.com/letung3105/cryptopals/pkg/ciphers"
	"github.com/pkg/errors"
)

// FixedXORCipher computes fixed xor combination of two hex encodes buffers
// and returns to hex encoded result
func FixedXORCipher(srcHex, targetHex []byte) ([]byte, error) {
	src := make([]byte, hex.DecodedLen(len(srcHex)))
	if _, err := hex.Decode(src, srcHex); err != nil {
		return nil, errors.Wrap(err, "Could not decode hex string")
	}

	target := make([]byte, hex.DecodedLen(len(targetHex)))
	if _, err := hex.Decode(target, targetHex); err != nil {
		return nil, errors.Wrap(err, "Could not decode hex string")
	}

	res, err := ciphers.FixedXOR(src, target)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get fixed xor")
	}

	dst := make([]byte, hex.EncodedLen(len(res)))
	n := hex.Encode(dst, res)
	return dst[:n], nil
}
