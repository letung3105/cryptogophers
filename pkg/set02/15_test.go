package set02

import (
	"bytes"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/utils"
)

func TestPKCS7Unpad(t *testing.T) {
	t.Parallel()
	tt := []struct {
		in  []byte
		out []byte
	}{
		{
			[]byte("ICE ICE BABY\x04\x04\x04\x04"),
			[]byte("ICE ICE BABY"),
		},
		{
			[]byte("ICE ICE BABY\x05\x05\x05\x05"),
			[]byte("ICE ICE BABY\x05\x05\x05\x05"),
		},
		{
			[]byte("ICE ICE BABY\x01\x02\x03\x04"),
			[]byte("ICE ICE BABY\x01\x02\x03\x04"),
		},
	}

	for _, tc := range tt {
		out := utils.PKCS7Unpad(tc.in)
		if !bytes.Equal(out, tc.out) {
			t.Errorf("unexpected output:\nhave %x\nwant %x", out, tc.out)
		}
	}
}
