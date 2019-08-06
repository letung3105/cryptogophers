package set01

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/pkg/errors"
)

// HexToB64 decodes a hex encoded buffer and encodes it back into base64
func HexToB64(src []byte) []byte {
	tmp := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(tmp, src)
	if err != nil {
		panic(errors.Wrap(err, "Could not decode hex string"))
	}

	encoding := base64.StdEncoding
	dst := make([]byte, encoding.EncodedLen(n))
	encoding.Encode(dst, tmp[:n])
	return dst
}
