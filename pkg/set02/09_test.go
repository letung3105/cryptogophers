package set02

import (
	"bytes"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/utils"
)

func TestPKCS7Pad(t *testing.T) {
	t.Parallel()
	test := struct {
		paddedLen int
		in        []byte
		out       []byte
	}{
		20,
		[]byte("YELLOW SUBMARINE"),
		[]byte("YELLOW SUBMARINE\x04\x04\x04\x04"),
	}

	out := utils.PKCS7Pad(test.in, test.paddedLen)
	if !bytes.Equal(out, test.out) {
		t.Errorf("unexpected output:\nhave %s\nwant %s", out, test.out)
	}
}
