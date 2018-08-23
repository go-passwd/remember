package main

import (
	"fmt"
	"os"

	"github.com/go-passwd/rememberizer"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg, ": ", rememberizer.ToSentences(arg, nil))
	}
}
