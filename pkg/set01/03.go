package set01

import (
	"encoding/hex"

	"github.com/pkg/errors"
	"github.com/letung3105/cryptogophers/pkg/decrypts"
)

// SingleByteXORHexDecrypt takes in hex string that was sngle byte xor enscrypted
// and find the plain text
func SingleByteXORHexDecrypt(cipherHex []byte) ([]byte, error) {
	cipher := make([]byte, hex.DecodedLen(len(cipherHex)))
	if _, err := hex.Decode(cipher, cipherHex); err != nil {
		return nil, errors.Wrap(err, "Could not decode hex string")
	}
	return decrypts.SingleByteXOR(cipher), nil
}
