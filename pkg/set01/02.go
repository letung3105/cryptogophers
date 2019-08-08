package set01

import (
	"encoding/hex"
	"fmt"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/pkg/errors"
)

// FixedXORCipher computes fixed xor combination of two hex encodes buffers
// and returns to hex encoded result
func FixedXORCipher(srcHex, keyHex []byte) ([]byte, error) {
	src := make([]byte, hex.DecodedLen(len(srcHex)))
	if _, err := hex.Decode(src, srcHex); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"could not decode %x", srcHex,
		))
	}

	key := make([]byte, hex.DecodedLen(len(keyHex)))
	if _, err := hex.Decode(key, keyHex); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"could not decode %x", keyHex,
		))
	}

	dst, err := crypts.FixedXOR(src, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"could not evaluate %x and %x", src, key,
		))
	}

	dstHex := make([]byte, hex.EncodedLen(len(dst)))
	n := hex.Encode(dstHex, dst)
	return dstHex[:n], nil
}
