package set01

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// DetectECB finds the line in the given file that was encrypted with AES in ECB mode
func DetectECB(filepath string, blocksize int) ([][]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"Could not open %q", filepath,
		))
	}

	var ciphers [][]byte
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		bHex := scanner.Bytes()
		b := make([]byte, hex.DecodedLen(len(bHex)))
		n, err := hex.Decode(b, bHex)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf(
				"could not decode: %x", bHex,
			))
		}
		b = b[:n]

		if utils.HasNonOverlapDup(b, blocksize) {
			ciphers = append(ciphers, b)
		}
	}
	return ciphers, nil
}
