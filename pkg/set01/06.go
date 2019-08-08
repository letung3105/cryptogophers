package set01

import (
	"encoding/base64"
	"io/ioutil"
	"math"
	"sort"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// RepeatingXORDecrypt decrypts the base64 encoded cipher text stored in the given file
func RepeatingXORDecrypt(filepath string, keysizeMax, keysizeTrials int) ([]byte, []byte, float64, error) {
	srcB64, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, nil, 0, errors.Wrapf(err, "could not read: %s", filepath)
	}

	b64Encoding := base64.StdEncoding
	src := make([]byte, b64Encoding.DecodedLen(len(srcB64)))
	n, err := b64Encoding.Decode(src, srcB64)
	if err != nil {
		return nil, nil, 0, errors.Wrapf(err, "could not decode: %s", srcB64)
	}
	src = src[:n]

	keyGuesses, err := guessKeysize(src, keysizeMax, keysizeTrials)
	if err != nil {
		return nil, nil, 0, errors.Wrap(err, "could not get possible keys size")
	}

	var key []byte
	var dst []byte
	minScore := math.MaxFloat64
	for _, guess := range keyGuesses {
		pk := guessKey(src, guess.size)
		plain := crypts.RepeatingXOR(src, pk)
		score := utils.ScoreTxtEn(plain)
		if score < minScore {
			minScore = score
			key = pk
			dst = plain
		}
	}

	return dst, key, minScore, nil
}

func guessKey(src []byte, keysize int) []byte {
	blocks := utils.BytesBlocksTranspose(utils.BytesBlockMake(src, keysize))
	dst := make([]byte, keysize)
	for i, block := range blocks {
		_, k, _ := SingleXORDecrypt(block)
		dst[i] = k
	}
	return dst
}

func guessKeysize(src []byte, keysizeMax, keysizeTrials int) (sortKeyGuess, error) {
	var keyGuesses sortKeyGuess
	for keysize := 1; keysize <= keysizeMax; keysize++ {
		d, err := getNormHamming(src, keysize)
		if err != nil {
			return nil, errors.Wrapf(err, "could not get normalized average hamming distance")
		}

		if d > 0 {
			keyGuesses = append(keyGuesses, &keyGuess{
				size:         keysize,
				normDistance: d,
			})
		}
	}

	sort.Sort(keyGuesses)
	if len(keyGuesses) > keysizeTrials {
		return keyGuesses[:keysizeTrials], nil
	}
	return keyGuesses, nil
}

func getNormHamming(src []byte, blocksize int) (float64, error) {
	var normDistance float64
	var blocks int
	x2Blocksize := 2 * blocksize

	for len(src) > x2Blocksize {
		distance, err := utils.HammingDistance(src[:blocksize], src[blocksize:x2Blocksize])
		if err != nil {
			return -1, errors.Wrap(err, "could not get hamming distance")
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

type keyGuess struct {
	size         int
	normDistance float64
}

type sortKeyGuess []*keyGuess

func (s sortKeyGuess) Len() int           { return len(s) }
func (s sortKeyGuess) Less(i, j int) bool { return s[i].normDistance < s[j].normDistance }
func (s sortKeyGuess) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
