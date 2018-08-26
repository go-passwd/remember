package rememberizer

import "fmt"

func ExampleToWords() {
	fmt.Println(ToWords("rT4s", nil))
	// Output:
	// [rope TOKYO 4 skype]
}

func ExampleToSentence() {
	fmt.Println(ToSentence("rT4s", nil))
	// Output:
	// rope TOKYO 4 skype
}

func ExampleToSentence_ownDict() {
	dict := MakeDict("go,password,toolkit", ",")
	fmt.Println(ToSentence("pt4g", dict))
	// Output:
	// password toolkit 4 go
}
