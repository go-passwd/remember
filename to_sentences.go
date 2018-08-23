package rememberizer

import "strings"

// ToSentences convert string to easy to remember sentences using sentences
// Default sentences is English
func ToSentences(s string, sentences *Sentences) []string {
	if sentences == nil {
		sentences = &English
	}
	out := strings.Split(s, "")
	for idx := range out {
		c := out[idx]
		lowerC := strings.ToLower(c)
		sentence, ok := (*sentences)[lowerC]
		if ok {
			if c == lowerC {
				out[idx] = sentence
			} else {
				out[idx] = strings.ToUpper(sentence)
			}
		}
	}
	return out
}
