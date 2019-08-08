package set01

import (
	"testing"
)

func TestHexSingleXORDecrypt(t *testing.T) {
	t.Parallel()
	test := struct {
		in []byte
	}{
		[]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"),
	}

	out, key, score, err := HexSingleXORDecrypt(test.in)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	t.Logf("key: %c | score: %.4f\n%s", key, score, out)
}
