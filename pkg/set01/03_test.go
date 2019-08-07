package set01

import (
	"testing"
)

func TestSingleByteXORhexDecrypt(t *testing.T) {
	t.Parallel()
	cipherHex := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	plain, err := SingleByteXORHexDecrypt(cipherHex)
	if err != nil {
		t.Fatalf("Could not find plain text: %v", err)
	}
	t.Logf("Single byte XOR plaintext: %s", plain)
}
