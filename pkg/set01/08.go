package set01

import (
	"bufio"
	"encoding/hex"
	"os"

	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// DetectECB finds the line in the given file that was encrypted with AES in ECB mode
func DetectECB(filepath string, blocksize int) ([][]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not open: %s", filepath)
	}

	var found [][]byte
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		srcHex := scanner.Bytes()
		src := make([]byte, hex.DecodedLen(len(srcHex)))
		n, err := hex.Decode(src, srcHex)
		if err != nil {
			return nil, errors.Wrapf(err, "could not decode: %s", srcHex)
		}
		src = src[:n]

		if utils.HasNonOverlapDup(src, blocksize) {
			found = append(found, srcHex)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.Wrapf(err, "could not scan file: %s", filepath)
	}

	return found, nil
}
