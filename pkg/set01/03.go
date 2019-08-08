package set01

import (
	"encoding/hex"
	"math"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// HexSingleXORDecrypt takes in hex string that was single byte xor encrypted
// and find the plain text
func HexSingleXORDecrypt(srcHex []byte) ([]byte, byte, float64, error) {
	src := make([]byte, hex.DecodedLen(len(srcHex)))
	if _, err := hex.Decode(src, srcHex); err != nil {
		return nil, 0x00, 0, errors.Wrapf(err, "could not decode: %s", srcHex)
	}
	plain, key, score := SingleXORDecrypt(src)
	return plain, key, score, nil
}

// SingleXORDecrypt guesses plain text from cipher text by brute-forcing the key
func SingleXORDecrypt(src []byte) ([]byte, byte, float64) {
	var dst []byte
	var key byte
	minScore := math.MaxFloat64
	for k := 0x00; k <= 0xff; k++ {
		plain := crypts.SingleXOR(src, byte(k))
		score := utils.ScoreTxtEn(plain)
		if score < minScore {
			minScore = score
			dst = plain
			key = byte(k)
		}
	}

	return dst, key, minScore
}
