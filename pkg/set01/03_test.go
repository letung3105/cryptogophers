package set01

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
)

func TestHexSingleXORDecrypt(t *testing.T) {
	t.Parallel()
	test := struct {
		in []byte
	}{
		[]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"),
	}

	out, key, _, err := HexSingleXORDecrypt(test.in)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	constructed := crypts.SingleXOR(out, key)
	constructedHex := make([]byte, hex.EncodedLen(len(constructed)))
	n := hex.Encode(constructedHex, constructed)
	constructedHex = constructedHex[:n]

	if !bytes.Equal(constructedHex, test.in) {
		t.Errorf("incorrect reconstructed input:\nhave %s\nwant %s", constructedHex, test.in)
	}
}
