package decrypts

import (
	"math"
	"reflect"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
)

func SingleByteXOR(src []byte) []byte {
	minScore := math.MaxFloat64
	var dst []byte
	for k := 0x00; k <= 0xff; k++ {
		plain := crypts.SingleByteXOR(src, byte(k))
		score := utils.ScoreTxtEn(plain)
		if score < minScore {
			minScore = score
			dst = plain
		}
	}
	return dst
}

func TestSingleByteXOR(t *testing.T) {
	t.Parallel()
	plain := []byte("This is a testing plain text for single byte xor encryption")
	key := byte('K')

	cipher := crypts.SingleByteXOR(plain, key)
	output := SingleByteXOR(cipher)
	if !reflect.DeepEqual(output, plain) {
		t.Errorf("Incorrect plain text: got %s, expected %s", output, plain)
	}
}
