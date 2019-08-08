package set01

import (
	"encoding/hex"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/pkg/errors"
)

// FixedXORCipher computes fixed xor combination of two hex encodes buffers
// and returns to hex encoded result
func FixedXORCipher(srcHex, keyHex []byte) ([]byte, error) {
	src := make([]byte, hex.DecodedLen(len(srcHex)))
	if _, err := hex.Decode(src, srcHex); err != nil {
		return nil, errors.Wrapf(err, "could not decode: %s", srcHex)
	}

	key := make([]byte, hex.DecodedLen(len(keyHex)))
	if _, err := hex.Decode(key, keyHex); err != nil {
		return nil, errors.Wrapf(err, "could not decode: %s", keyHex)
	}

	dst, err := crypts.FixedXOR(src, key)
	if err != nil {
		return nil, errors.Wrapf(err, "could not evaluate: %s and %s", src, key)
	}

	dstHex := make([]byte, hex.EncodedLen(len(dst)))
	n := hex.Encode(dstHex, dst)
	return dstHex[:n], nil
}
