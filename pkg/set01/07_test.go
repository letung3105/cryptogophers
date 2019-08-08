package set01

import "testing"

func TestChallenge07(t *testing.T) {
	t.Parallel()
	filepath := "../../data/07.txt"
	key := []byte("YELLOW SUBMARINE")

	data, err := Challenge07(filepath, key)
	if err != nil {
		t.Fatalf("Challenge07(%s, %s) = %s", filepath, key, err)
	}
	t.Logf("Challenge07\n%s", data)
}
