package rememberizer

import "strings"

// ToWords convert string to easy to remember sentences using dict
// Default dict is EnglishDict
func ToWords(s string, dict *Dict) []string {
	if dict == nil {
		dict = EnglishDict
	}
	sentence := strings.Split(s, "")
	for idx, char := range sentence {
		sentence[idx] = dict.Get(char)
	}
	return sentence
}

func ToSentence(s string, dict *Dict) string {
	return strings.Join(ToWords(s, dict), " ")
}
