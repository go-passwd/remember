package rememberizer

import "strings"

type Dict map[string]string

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

// English dict
var English = Dict{
	"a": "apple",
	"b": "bestbuy",
	"c": "coffee",
	"d": "drip",
	"e": "egg",
	"f": "fruit",
	"g": "golf",
	"h": "hulu",
	"i": "iphone",
	"j": "jack",
	"k": "korean",
	"l": "laptop",
	"m": "music",
	"n": "nut",
	"o": "omlet",
	"p": "park",
	"q": "queen",
	"r": "rope",
	"s": "skype",
	"t": "tokyo",
	"u": "usa",
	"v": "visa",
	"w": "walmart",
	"x": "xbox",
	"y": "yelp",
	"z": "zip",
}
