package set01

import (
	"reflect"
	"testing"
)

func TestRepeatingXORHex(t *testing.T) {
	t.Parallel()
	plain := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	expected := []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")

	output := RepeatingXORHex(plain, key)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Unexpected output: got %s, expected %s", output, expected)
	}
}
