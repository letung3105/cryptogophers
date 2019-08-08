package utils

import "bytes"

// ScoreTxtEn produces simple normalized score for given text
// based on the letters frequency in the english language
func ScoreTxtEn(src []byte) float64 {
	frq := []byte(" etaoinsrhldcumfgpywENTb,.vk-\"_'x)(;0j1q=2:z/*!?$35>{}49[]867\\+|&<%@#^`~")
	if len(src) == 0 {
		return 0
	}

	var score float64
	for _, c := range src {
		ind := bytes.IndexByte(frq, c)
		if ind != -1 {
			score += float64(ind)
		} else {
			score += float64(len(frq))
		}
	}
	score /= float64(len(src))
	return score
}
