package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-passwd/rememberizer"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg, ":", strings.Join(rememberizer.ToSentences(arg, nil), " "))
	}
}
