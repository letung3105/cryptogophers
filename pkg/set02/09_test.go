package set02

import (
	"bytes"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/utils"
)

func TestPKCS7Pad(t *testing.T) {
	t.Parallel()
	tc := struct {
		paddedLen int
		in        []byte
		out       []byte
	}{
		20,
		[]byte("YELLOW SUBMARINE"),
		[]byte("YELLOW SUBMARINE\x04\x04\x04\x04"),
	}

	out := utils.PKCS7Pad(tc.in, tc.paddedLen)
	if !bytes.Equal(out, tc.out) {
		t.Errorf("unexpected output:\nhave %s\nwant %s", out, tc.out)
	}
}
