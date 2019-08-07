package decrypts

import (
	"reflect"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
)

func TestSingleByteXOR(t *testing.T) {
	t.Parallel()
	plain := []byte("This is a testing plain text for single byte xor encryption")
	key := byte('K')

	cipher := crypts.SingleByteXOR(plain, key)
	output := SingleByteXOR(cipher)
	if !reflect.DeepEqual(output, plain) {
		t.Errorf("Incorrect plain text: got %s, expected %s", output, plain)
	}
}
