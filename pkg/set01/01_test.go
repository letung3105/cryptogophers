package set01

import (
	"reflect"
	"testing"
)

func TestHexToB64(t *testing.T) {
	t.Parallel()
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expected := []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")

	output, err := HexToB64(input)
	if err != nil {
		t.Fatalf("Could not convert hex to base64: %v", err)
	}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Unexpected output: got %s, expected %s", output, expected)
	}
}
