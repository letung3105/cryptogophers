package set01

import (
	"reflect"
	"testing"
)

func TestFixedXORHex(t *testing.T) {
	t.Parallel()
	src := []byte("1c0111001f010100061a024b53535009181c")
	target := []byte("686974207468652062756c6c277320657965")
	expected := []byte("746865206b696420646f6e277420706c6179")

	output, err := FixedXORHex(src, target)
	if err != nil {
		t.Fatalf("Could not compute fixed xor of the given buffers: %v", err)
	}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Unexpected output: got %s, expected %s", output, expected)
	}
}
