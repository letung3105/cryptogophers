package set01

import (
	"bufio"
	"encoding/hex"
	"math"
	"os"

	"github.com/letung3105/cryptogophers/pkg/decrypts"
	"github.com/pkg/errors"
)

// DetectSingleByteXOR find the single line the the given file that is single byte xor encrypted
func DetectSingleByteXOR(filepath string) ([]byte, byte, error) {
	var key byte
	f, err := os.Open(filepath)
	if err != nil {
		return nil, key, errors.Wrap(err, "Could not open file")
	}

	var dst []byte
	minScore := math.MaxFloat64
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		srcHex := scanner.Bytes()
		src := make([]byte, hex.DecodedLen(len(srcHex)))
		n, err := hex.Decode(src, srcHex)
		if err != nil {
			return nil, key, errors.Wrap(err, "Could not decode hex string")
		}

		plain, pk, score := decrypts.SingleByteXOR(src[:n])
		if score < minScore {
			minScore = score
			dst = plain
			key = pk
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, key, errors.Wrap(err, "Could not scan file")
	}
	return dst, key, nil
}
