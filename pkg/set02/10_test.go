package set02

import "testing"

func TestDecryptCBC(t *testing.T) {
	t.Parallel()
	tc := struct {
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

	out, err := DecryptCBC(tc.filepath, tc.key, tc.iv)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	// TODO: add result file to tc against output
	t.Logf("output:\n%s", out)
}
