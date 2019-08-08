package set01

import (
	"encoding/base64"
	"io/ioutil"
	"math"
	"sort"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/decrypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// RepeatingXORDecrypt decrypts the base64 encoded cipher text stored in the given file
func RepeatingXORDecrypt(filepath string, keysizeMax, keysizeTrials int) ([]byte, []byte, error) {
	cipherB64, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Could not open file")
	}

	b64Encoding := base64.StdEncoding
	cipher := make([]byte, b64Encoding.DecodedLen(len(cipherB64)))
	n, err := b64Encoding.Decode(cipher, cipherB64)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Could not decode base64")
	}
	cipher = cipher[:n]

	possibles, err := getPossibleKeysizes(cipher, keysizeMax, keysizeTrials)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Could not get possible keys information")
	}

	var possibleKey []byte
	var possiblePlain []byte
	minScore := math.MaxFloat64
	for _, p := range possibles {
		key := getPossibleKey(cipher, p.size)
		plain := crypts.RepeatingXOR(cipher, key)
		score := utils.ScoreTxtEn(plain)
		if score < minScore {
			minScore = score
			possibleKey = key
			possiblePlain = plain
		}
	}

	return possiblePlain, possibleKey, nil
}

func getPossibleKey(cipher []byte, keysize int) []byte {
	blocks := utils.BytesBlocksTranspose(utils.BytesBlockMake(cipher, uint(keysize)))
	key := make([]byte, keysize)
	for i, block := range blocks {
		_, k, _ := decrypts.SingleByteXOR(block)
		key[i] = k
	}
	return key
}

func getPossibleKeysizes(cipher []byte, keysizeMax, keysizeTrials int) (sortPossibleKey, error) {
	var keysInfo sortPossibleKey
	for keysize := 1; keysize <= keysizeMax; keysize++ {
		normDistance, err := normAvgDistance(cipher, keysize)
		if err != nil {
			return nil, errors.Wrap(err, "Could not get normalized average hamming distance")
		}

		if normDistance > 0 {
			keysInfo = append(keysInfo, &possibleKey{
				size:         keysize,
				normDistance: normDistance,
			})
		}
	}

	sort.Sort(keysInfo)
	if len(keysInfo) > keysizeTrials {
		return keysInfo[:keysizeTrials], nil
	}
	return keysInfo, nil
}

func normAvgDistance(src []byte, blocksize int) (float64, error) {
	var normDistance float64
	var blocks int
	x2Blocksize := 2 * blocksize

	for len(src) > x2Blocksize {
		distance, err := utils.HammingDistance(src[:blocksize], src[blocksize:x2Blocksize])
		if err != nil {
			return -1, errors.Wrap(err, "Could not get hamming distance")
		}

		normDistance += float64(distance) / float64(blocksize)
		src = src[x2Blocksize:]
		blocks++
	}

	if blocks > 0 {
		normDistance /= float64(blocks)
	}
	return normDistance, nil
}

type possibleKey struct {
	size         int
	normDistance float64
}

type sortPossibleKey []*possibleKey

func (s sortPossibleKey) Len() int           { return len(s) }
func (s sortPossibleKey) Less(i, j int) bool { return s[i].normDistance < s[j].normDistance }
func (s sortPossibleKey) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
