package set02

import (
	"bytes"
	"testing"
)

func TestECBOracle(t *testing.T) {
	t.Parallel()
	test := struct {
		sampleSize int
	}{
		512,
	}

	for i := 0; i <= test.sampleSize; i++ {
		_, err := ECBOracle(bytes.Repeat([]byte("A"), i))
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
	}
}

func TestBreakECBOracle(t *testing.T) {
	t.Parallel()
	out, blocksize, err := BreakECBOracle()
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	// TODO: add result file to test against output
	t.Logf("block size %d\n%s", blocksize, out)
}
