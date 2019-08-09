package set01

import (
	"bytes"
	"testing"
)

func TestFixedXORCipher(t *testing.T) {
	t.Parallel()
	tc := struct {
		key []byte
		in  []byte
		out []byte
	}{
		[]byte("686974207468652062756c6c277320657965"),
		[]byte("1c0111001f010100061a024b53535009181c"),
		[]byte("746865206b696420646f6e277420706c6179"),
	}

	out, err := FixedXORCipher(tc.in, tc.key)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	if !bytes.Equal(out, tc.out) {
		t.Errorf("unexpected output:\nhave %s\nwant %s", out, tc.out)
	}
}
