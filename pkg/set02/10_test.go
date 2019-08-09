package set02

import "testing"

func TestDecryptCBC(t *testing.T) {
	t.Parallel()
	test := struct {
		filepath string
		key      []byte
		iv       []byte
	}{
		"./testdata/10.txt",
		[]byte("YELLOW SUBMARINE"),
		[]byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}

	out, err := DecryptCBC(test.filepath, test.key, test.iv)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	t.Logf("output:\n%s", out)
}
