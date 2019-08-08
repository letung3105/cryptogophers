package set01

import (
	"bufio"
	"encoding/hex"
	"math"
	"os"

	"github.com/pkg/errors"
)

// DetectSingleXOR find the single line the the given file that is single byte xor encrypted
func DetectSingleXOR(filepath string) ([]byte, byte, float64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, 0x00, 0, errors.Wrapf(err, "could not open: %s", filepath)
	}

	var dst []byte
	var key byte
	minScore := math.MaxFloat64
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		srcHex := scanner.Bytes()
		src := make([]byte, hex.DecodedLen(len(srcHex)))
		n, err := hex.Decode(src, srcHex)
		if err != nil {
			return nil, 0x00, 0, errors.Wrapf(err, "could not decode: %s", srcHex)
		}

		plain, pk, score := SingleXORDecrypt(src[:n])
		if score < minScore {
			minScore = score
			dst = plain
			key = pk
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, 0x00, 0, errors.Wrapf(err, "could not scan file: %s", filepath)
	}
	return dst, key, minScore, nil
}
