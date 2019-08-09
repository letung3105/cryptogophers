package set01

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/pkg/errors"
)

func TestECBDecryptB64AES(t *testing.T) {
	t.Parallel()
	test := struct {
		filepath string
		key      []byte
	}{
		"./testdata/07.txt",
		[]byte("YELLOW SUBMARINE"),
	}

	out, err := ECBDecryptB64AES(test.filepath, test.key)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	inB64, err := ioutil.ReadFile(test.filepath)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "could not read: %s", test.filepath))
	}

	b64 := base64.StdEncoding
	in := make([]byte, b64.DecodedLen(len(inB64)))
	n, err := b64.Decode(in, inB64)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "could not decode: %s", inB64))
	}
	in = in[:n]

	c, err := aes.NewCipher(test.key)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "could not create cipher from key: %s", test.key))
	}
	encrypter := crypts.NewECBEncrypter(c)
	constructed := make([]byte, len(out))
	encrypter.CryptBlocks(constructed, out)

	if !bytes.Equal(constructed, in) {
		t.Errorf("incorrect reconstructed input:\nhave %x\nwant %x", constructed, in)
	}
}
