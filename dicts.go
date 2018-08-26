package rememberizer

import "strings"

// EnglishWords is a set of a words from…
const EnglishWords = "apple,bestbuy,coffee,drip,egg,fruit,golf,hulu,iphone,jack,korean,laptop,music,nut,omlet,park,queen,rope,skype,tokyo,usa,visa,walmart,xbox,yelp,zip"

// EnglishDict is a dict made from EnglishWords
var EnglishDict = MakeDict(EnglishWords, ",")

// MakeDict create dict from set of words separated by separator
func MakeDict(words string, separator string) *Dict {
	dict := Dict{}
	for _, word := range strings.Split(words, separator) {
		dict[strings.Split(word, "")[0]] = word
	}
	return &dict
}

type Dict map[string]string

// Get a word mapped to key.
// If key is uppercase, word is uppercase too.
// If key doesn't exisis, return key.
func (d *Dict) Get(key string) string {
	lowerKey := strings.ToLower(key)
	word, ok := (*d)[lowerKey]
	if ok {
		if key == lowerKey {
			return word
		}
		return strings.ToUpper(word)
	}
	return key
}
