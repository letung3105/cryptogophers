package set01

import (
	"encoding/hex"

	"github.com/pkg/errors"
	"github.com/letung3105/cryptogophers/pkg/decrypts"
)

// SingleByteXORDecrypt takes in hex string that was sngle byte xor enscrypted
// and find the plain text
func SingleByteXORDecrypt(cipherHex []byte) ([]byte, byte, error) {
	var key byte
	cipher := make([]byte, hex.DecodedLen(len(cipherHex)))
	if _, err := hex.Decode(cipher, cipherHex); err != nil {
		return nil, key, errors.Wrap(err, "Could not decode hex string")
	}
	plain, key, _ := decrypts.SingleByteXOR(cipher)
	return plain, key, nil
}
