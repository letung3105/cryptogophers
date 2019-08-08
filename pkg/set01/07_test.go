package set01

import "testing"

func TestECBDecryptB64AES(t *testing.T) {
	t.Parallel()
	test := struct {
		filepath string
		key      []byte
	}{
		"../../data/07.txt",
		[]byte("YELLOW SUBMARINE"),
	}

	out, err := ECBDecryptB64AES(test.filepath, test.key)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	t.Logf("output:\n%s", out)
}
