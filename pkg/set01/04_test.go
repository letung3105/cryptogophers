package set01

import (
	"bytes"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
)

func TestDetectSingleXOR(t *testing.T) {
	t.Parallel()
	test := struct {
		filepath string
	}{
		"./testdata/04.txt",
	}

	out, in, key, _, err := DetectSingleXOR(test.filepath)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	constructed := crypts.SingleXOR(out, key)
	if !bytes.Equal(constructed, in) {
		t.Errorf("incorrect reconstructed input:\nhave %x\nwant %x", constructed, in)
	}
}
