package set01

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/pkg/errors"
)

// HexToB64 decodes a hex encoded buffer and encodes it back into base64
func HexToB64(src []byte) ([]byte, error) {
	tmp := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(tmp, src)
	if err != nil {
		return nil, errors.Wrapf(err, "could not decode: %s", src)
	}
	tmp = tmp[:n]

	encoding := base64.StdEncoding
	dst := make([]byte, encoding.EncodedLen(n))
	encoding.Encode(dst, tmp)
	return dst, nil
}
