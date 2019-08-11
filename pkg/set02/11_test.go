package set02

import (
	"bytes"
	"testing"
)

func TestEncryptOracle(t *testing.T) {
	t.Parallel()
	test := struct {
		sampleSize int
		in         []byte
	}{
		512,
		[]byte("This is a test input.\nThis is added for nothing!\n"),
	}

	count := map[string]int{
		"ECB": 0,
		"CBC": 0,
	}

	for i := 0; i < test.sampleSize; i++ {
		_, isECB, err := EncryptOracle(test.in, 16)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		if isECB {
			count["ECB"]++
		} else {
			count["CBC"]++
		}
	}

	if count["ECB"] == 0 || count["CBC"] == 0 {
		t.Errorf(
			"encryption mode was not randomized: got %d ECB cases and %d CBC cases",
			count["ECB"], count["CBC"],
		)
	}
}

func TestDetectOracle(t *testing.T) {
	t.Parallel()
	test := struct {
		sampleSize int
		in         []byte
	}{
		512,
		bytes.Repeat([]byte("B"), 128),
	}

	for i := 0; i < test.sampleSize; i++ {
		cipher, isECB, err := EncryptOracle(test.in, 16)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		guessIsECB := DetectOracle(cipher)
		if isECB != guessIsECB {
			t.Errorf("incorrect guess:\nhave %t\nwant %t", isECB, guessIsECB)
		}
	}
}
