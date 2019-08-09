package set01

import (
	"bytes"
	"testing"
)

func TestHexToB64(t *testing.T) {
	t.Parallel()
	tc := struct {
		in  []byte
		out []byte
	}{
		in:  []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"),
		out: []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"),
	}

	out, err := HexToB64(tc.in)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	if !bytes.Equal(out, tc.out) {
		t.Errorf("unexpected output:\nhave %s\nwant %s", out, tc.out)
	}
}
