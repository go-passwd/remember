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
		k, ok := (*sentences)[out[idx]]
		if ok {
			out[idx] = k
		}
	}
	return out
}
