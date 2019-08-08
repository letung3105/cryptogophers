package crypts

import (
	"bytes"
	"crypto/aes"
	"testing"
)

var ecbAESTests = []struct {
	name string
	blocksize int
	key  []byte
	in   []byte
	out  []byte
}{
	{
		"ECB-AES128",
		16,
		commonKey128,
		commonInput,
		[]byte{
			0x3a, 0xd7, 0x7b, 0xb0, 0x47, 0xd3, 0xa6, 0x6a, 0x09, 0x8c, 0xef, 0xa2, 0x36, 0x4e, 0xf6, 0x97,
			0xf5, 0xd3, 0xd8, 0x50, 0x5b, 0x36, 0x99, 0x9e, 0xd8, 0x78, 0x55, 0x99, 0xaf, 0x6b, 0xad, 0xaf,
			0x43, 0xb1, 0xc7, 0xd5, 0xf8, 0x9c, 0xe2, 0xe8, 0x31, 0x80, 0xbe, 0x0e, 0x30, 0xd0, 0x63, 0x88,
			0x7b, 0x0c, 0x75, 0x82, 0xee, 0x7a, 0x83, 0xd8, 0xf2, 0x22, 0x37, 0x00, 0x17, 0x45, 0xd2, 0xd4,
		},
	},
	{
		"ECB-AES192",
		24,
		commonKey192,
		commonInput,
		[]byte{
			0xbd, 0x33, 0x4f, 0x1d, 0x6e, 0x45, 0xf2, 0x5f, 0xf7, 0x12, 0xa2, 0x14, 0x57, 0x1f, 0xa5, 0xcc,
			0x97, 0x41, 0x04, 0x84, 0x6d, 0x0a, 0xd3, 0xad, 0x77, 0x34, 0xec, 0xb3, 0xec, 0xee, 0x4e, 0xef,
			0xef, 0x7a, 0xfd, 0x22, 0x70, 0xe2, 0xe6, 0x0a, 0xdc, 0xe0, 0xba, 0x2f, 0xac, 0xe6, 0x44, 0x4e,
			0x9a, 0x4b, 0x41, 0xba, 0x73, 0x8d, 0x6c, 0x72, 0xfb, 0x16, 0x69, 0x16, 0x03, 0xc1, 0x8e, 0x0e,
		},
	},
	{
		"ECB-AES256",
		32,
		commonKey256,
		commonInput,
		[]byte{
			0xf3, 0xee, 0xd1, 0xbd, 0xb5, 0xd2, 0xa0, 0x3c, 0x06, 0x4b, 0x5a, 0x7e, 0x3d, 0xb1, 0x81, 0xf8,
			0x59, 0x1c, 0xcb, 0x10, 0xd4, 0x10, 0xed, 0x26, 0xdc, 0x5b, 0xa7, 0x4a, 0x31, 0x36, 0x28, 0x70,
			0xb6, 0xed, 0x21, 0xb9, 0x9c, 0xa6, 0xf4, 0xf9, 0xf1, 0x53, 0xe7, 0xb1, 0xbe, 0xaf, 0xed, 0x1d,
			0x23, 0x30, 0x4b, 0x7a, 0x39, 0xf9, 0xf3, 0xff, 0x06, 0x7d, 0x8d, 0x8f, 0x9e, 0x24, 0xec, 0xc7,
		},
	},
}

func TestECBEncrypterAES(t *testing.T) {
	t.Parallel()
	for _, test := range ecbAESTests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c, err := aes.NewCipher(test.key)
			if err != nil {
				t.Fatalf("NewCipher(%d bytes) = %s", len(test.key), err)
			}

			encrypter := NewECBEncrypter(c)
			if encrypter.BlockSize() != test.blocksize {
				t.Errorf("ECBEncrypter: incorrect block size\nhave %x\nwant %x", encrypter.BlockSize(), test.blocksize)
			}

			data := make([]byte, len(test.in))
			copy(data, test.in)
			encrypter.CryptBlocks(data, data)
			if !bytes.Equal(test.out, data) {
				t.Errorf("ECBEncrypter\nhave %x\nwant %x", data, test.out)
			}
		})
	}
}

func TestECBDecrypterAES(t *testing.T) {
	t.Parallel()
	for _, test := range ecbAESTests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c, err := aes.NewCipher(test.key)
			if err != nil {
				t.Fatalf("NewCipher(%d bytes) = %s", len(test.key), err)
			}

			decrypter := NewECBDecrypter(c)
			if decrypter.BlockSize() != test.blocksize {
				t.Errorf("ECBEncrypter: incorrect block size\nhave %x\nwant %x", decrypter.BlockSize(), test.blocksize)
			}

			data := make([]byte, len(test.out))
			copy(data, test.out)
			decrypter.CryptBlocks(data, data)
			if !bytes.Equal(test.in, data) {
				t.Errorf("ECBDecrypter\nhave %x\nwant %x", data, test.in)
			}
		})
	}
}
