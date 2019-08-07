package utils

import "bytes"

var lettersFrqOrder = []byte(" etaoinsrhldcumfgpywENTb,.vk-\"_'x)(;0j1q=2:z/*!?$35>{}49[]867\\+|&<%@#^`~")

// ScoreTextEn produces simple normalized score for given text
// based on the letters frequency in the english language
func ScoreTxtEn(src []byte) float64 {
	if len(src) == 0 {
		return 0
	}

	var score float64
	for _, c := range src {
		chrIndex := bytes.IndexByte(lettersFrqOrder, c)
		if chrIndex != -1 {
			score += float64(chrIndex)
		} else {
			score += float64(len(lettersFrqOrder))
		}
	}
	score /= float64(len(src))
	return score
}
