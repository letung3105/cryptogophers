package set01

import (
	"bytes"
	"testing"
)

func TestHexToB64(t *testing.T) {
	t.Parallel()
	test := struct {
		in  []byte
		out []byte
	}{
		in:  []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"),
		out: []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"),
	}

	out, err := HexToB64(test.in)
	if err != nil {
		t.Fatalf("HexToB64(%x) = %+v", test.in, err)
	}
	if !bytes.Equal(out, test.out) {
		t.Errorf("unexpected output\nhave %x\nwant %x", out, test.out)
	}
}
