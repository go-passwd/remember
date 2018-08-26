package rememberizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeDict(t *testing.T) {
	dict := MakeDict("go,password,toolkit", ",")
	e := Dict{
		"g": "go",
		"p": "password",
		"t": "toolkit",
	}
	assert.EqualValues(t, e, *dict)
}

func TestDict_Get(t *testing.T) {
	assert.Equal(t, "skype", EnglishDict.Get("s"))
	assert.Equal(t, "SKYPE", EnglishDict.Get("S"))
	assert.Equal(t, "4", EnglishDict.Get("4"))
}
