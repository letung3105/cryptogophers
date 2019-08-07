package set01

import (
	"bufio"
	"encoding/hex"
	"math"
	"os"

	"github.com/letung3105/cryptogophers/pkg/decrypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// DetectSingleByteXOR find the single line the the given file that is single byte xor encrypted
func DetectSingleByteXOR(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "Could not open file")
	}

	minScore := math.MaxFloat64
	var dst []byte
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		srcHex := scanner.Bytes()
		src := make([]byte, hex.DecodedLen(len(srcHex)))
		n, err := hex.Decode(src, srcHex)
		if err != nil {
			return nil, errors.Wrap(err, "Could not decode hex string")
		}

		plain := decrypts.SingleByteXOR(src[:n])
		score := utils.ScoreTxtEn(plain)
		if score < minScore {
			minScore = score
			dst = plain
		}
	}
	return dst, nil
}
