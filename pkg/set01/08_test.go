package set01

import "testing"

func TestDetectECB(t *testing.T) {
	t.Parallel()
	filepath := "../../data/09.txt"
	blocksize := 16

	ciphers, err := DetectECB(filepath, blocksize)
	if err != nil {
		t.Fatalf("DetectECB(%q) = %+v", filepath, err)
	}
	for _, cipher := range ciphers {
		t.Logf("blocksize: %d\n%x", blocksize, cipher)
	}
}
