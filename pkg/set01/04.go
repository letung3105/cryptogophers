package set01

import (
	"bufio"
	"encoding/hex"
	"math"
	"os"

	"github.com/pkg/errors"
)

// DetectSingleXOR find the single line the the given file that is single byte xor encrypted
func DetectSingleXOR(filepath string) ([]byte, []byte, byte, float64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, nil, 0x00, 0, errors.Wrapf(err, "could not open: %s", filepath)
	}

	var dst []byte
	var key byte
	var src []byte
	minScore := math.MaxFloat64
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		cipherHex := scanner.Bytes()
		cipher := make([]byte, hex.DecodedLen(len(cipherHex)))
		n, err := hex.Decode(cipher, cipherHex)
		if err != nil {
			return nil, nil, 0x00, 0, errors.Wrapf(err, "could not decode: %s", cipherHex)
		}
		cipher = cipher[:n]

		plain, pk, score := SingleXORDecrypt(cipher)
		if score < minScore {
			minScore = score
			src = cipher
			dst = plain
			key = pk
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, 0x00, 0, errors.Wrapf(err, "could not scan file: %s", filepath)
	}
	return dst, src, key, minScore, nil
}
