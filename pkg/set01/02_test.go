package set01

import (
	"reflect"
	"testing"
)

func TestFixedXORHex(t *testing.T) {
	input1 := []byte("1c0111001f010100061a024b53535009181c")
	input2 := []byte("686974207468652062756c6c277320657965")
	expected := []byte("746865206b696420646f6e277420706c6179")

	output := FixedXORHex(input1, input2)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Unexpected output: got %s, expected %s", output, expected)
	}
}
