package set01

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
)

// HexToB64 decodes a hex encoded buffer and encodes it back into base64
func HexToB64(src []byte) ([]byte, error) {
	tmp := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(tmp, src)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"could not decode %x", src,
		))
	}
	tmp = tmp[:n]

	encoding := base64.StdEncoding
	dst := make([]byte, encoding.EncodedLen(n))
	encoding.Encode(dst, tmp)
	return dst, nil
}
