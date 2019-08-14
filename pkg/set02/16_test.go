package set02

import (
	"bytes"
	"crypto/aes"
	"math/rand"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

func TestCBCOracle(t *testing.T) {
	t.Parallel()
	keyLen := 16
	prefixBytes := []byte("comment1=cooking%20MCs;userdata=")
	suffixBytes := []byte(";comment2=%20like%20a%20pound%20of%20bacon")
	tt := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{
			"InputNormal",
			[]byte("This is a test input"),
			bytes.Join([][]byte{
				prefixBytes,
				[]byte("This is a test input"),
				suffixBytes,
			}, nil),
		},
		{
			"InputWithInvalidChars",
			[]byte(";This=is a test input;"),
			bytes.Join([][]byte{
				prefixBytes,
				[]byte("Thisis a test input"),
				suffixBytes,
			}, nil),
		},
		{
			"InputEmpty",
			[]byte{},
			bytes.Join([][]byte{
				prefixBytes,
				suffixBytes,
			}, nil),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			key := make([]byte, keyLen)
			if _, err := rand.Read(key); err != nil {
				t.Fatal(errors.Wrapf(err, "could not create random key of size: %d", keyLen))
			}

			c, err := aes.NewCipher(key)
			if err != nil {
				t.Fatal(errors.Wrapf(err, "could not create cipher from key: %x", key))
			}

			iv := make([]byte, c.BlockSize())
			if _, err := rand.Read(iv); err != nil {
				t.Fatal(errors.Wrapf(err, "could not create random IV of size: %d", c.BlockSize()))
			}

			out := CBCOracle(tc.in, c, iv)
			decrypter := crypts.NewCBCDecrypter(c, iv)
			decrypter.CryptBlocks(out, out)

			out = utils.PKCS7Unpad(out)
			if !bytes.Equal(out, tc.out) {
				t.Errorf("unexpected output:\nhave %s\nwant %s", out, tc.out)
			}
		})
	}
}

func TestCBCCheckAdmin(t *testing.T) {
	t.Parallel()
	keyLen := 16
	tt := []struct {
		name string
		in   []byte
		out  bool
	}{
		{
			"IsAdmin",
			[]byte("comment1=cooking%20MCs;userdata=;admin=true;comment2=%20like%20a%20pound%20of%20bacon"),
			true,
		},
		{
			"IsNotAdmin",
			[]byte("comment1=cooking%20MCs;userdata=;comment2=%20like%20a%20pound%20of%20bacon"),
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			key := make([]byte, keyLen)
			if _, err := rand.Read(key); err != nil {
				t.Fatal(errors.Wrapf(err, "could not create random key of size: %d", keyLen))
			}

			c, err := aes.NewCipher(key)
			if err != nil {
				t.Fatal(errors.Wrapf(err, "could not create cipher from key: %x", key))
			}

			iv := make([]byte, c.BlockSize())
			if _, err := rand.Read(iv); err != nil {
				t.Fatal(errors.Wrapf(err, "could not create random IV of size: %d", c.BlockSize()))
			}

			data := make([]byte, len(tc.in))
			copy(data, tc.in)
			data = utils.PKCS7Pad(data, len(data)+c.BlockSize()-(len(data)%c.BlockSize()))

			encrypter := crypts.NewCBCEncrypter(c, iv)
			encrypter.CryptBlocks(data, data)

			out := CBCCheckAdmin(data, c, iv)
			if out != tc.out {
				t.Errorf("unexpected output:\nhave %t\nwant %t", out, tc.out)
			}
		})
	}
}

func TestCBCBitFlip(t *testing.T) {
	keyLen := 16

	key := make([]byte, keyLen)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(errors.Wrapf(err, "could not create random key of size: %d", keyLen))
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "could not create cipher from key: %x", key))
	}

	iv := make([]byte, c.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		t.Fatal(errors.Wrapf(err, "could not create random IV of size: %d", c.BlockSize()))
	}

	data := CBCBitFlipping(c, iv)
	isAdmin := CBCCheckAdmin(data, c, iv)
	if !isAdmin {
		t.Errorf("unexpected output:\nhave %t\nwant %t", isAdmin, true)
	}
}
