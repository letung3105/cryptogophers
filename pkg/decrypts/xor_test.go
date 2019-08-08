package decrypts

import (
	"reflect"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
)

func TestSingleByteXOR(t *testing.T) {
	t.Parallel()
	tt := []struct {
		plain []byte
		key   byte
	}{
		{plain: []byte("This is a testing plain text for single byte xor encryption"), key: byte('K')},
		{plain: []byte("short sentence!"), key: byte('K')},
		{plain: []byte("This! Contain punctuations?!"), key: byte('K')},
		{plain: []byte("")},
	}

	for _, tc := range tt {
		cipher := crypts.SingleByteXOR(tc.plain, tc.key)
		output, key, _ := SingleByteXOR(cipher)
		if key != tc.key {
			t.Fatalf("Incorrect key: got %c, expected %c", key, tc.key)
		}
		if !reflect.DeepEqual(output, tc.plain) {
			t.Errorf("Incorrect plain text: got %s, expected %s", output, tc.plain)
		}
	}
}
