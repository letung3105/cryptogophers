package set01

import (
	"fmt"
	"testing"
)

func TestDetectECB(t *testing.T) {
	t.Parallel()
	test := struct {
		filepath  string
		blocksize int
	}{
		"./testdata/08.txt",
		16,
	}

	ciphers, err := DetectECB(test.filepath, test.blocksize)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	out := ""
	for _, c := range ciphers {
		out += fmt.Sprintf("+ %s\n", c)
	}
	t.Logf("found:\n%s", out)
}
