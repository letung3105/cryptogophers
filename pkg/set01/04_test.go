package set01

import "testing"

func TestDetectSingleXOR(t *testing.T) {
	t.Parallel()
	test := struct {
		filepath string
	}{
		"../../data/04.txt",
	}

	out, key, score, err := DetectSingleXOR(test.filepath)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	t.Logf("key: %c | score: %.4f\n%s", key, score, out)
}
