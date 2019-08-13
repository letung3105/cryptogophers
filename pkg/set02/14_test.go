package set02

import (
	"bytes"
	"testing"
)

func TestComplexECBOracle(t *testing.T) {
	t.Parallel()
	test := struct {
		sampleSize int
	}{
		512,
	}

	for i := 0; i <= test.sampleSize; i++ {
		_, err := ComplexECBOracle(bytes.Repeat([]byte("A"), i))
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
	}
}

func TestBreakComplexECBOracle(t *testing.T) {
	t.Parallel()
	out, err := BreakComplexECBOracle()
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	t.Logf("found:\n%s", out)
}
