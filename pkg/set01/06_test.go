package set01

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/pkg/errors"
)

func TestRepeatingXORDecrypt(t *testing.T) {
	t.Parallel()
	tc := struct {
		filepath      string
		keysizeMax    int
		keysizeTrials int
	}{
		"./testdata/06.txt", 40, 1,
	}

	out, key, _, err := RepeatingXORDecrypt(tc.filepath, tc.keysizeMax, tc.keysizeTrials)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	inB64, err := ioutil.ReadFile(tc.filepath)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "could not read: %s", tc.filepath))
	}

	b64 := base64.StdEncoding
	in := make([]byte, b64.EncodedLen(len(inB64)))
	n, err := b64.Decode(in, inB64)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "could not decode: %s", inB64))
	}
	in = in[:n]

	constructed := crypts.RepeatingXOR(out, key)
	if !bytes.Equal(constructed, in) {
		t.Errorf("incorrect reconstructed input\nhave %x\nwant %x", constructed, in)
	}
}
