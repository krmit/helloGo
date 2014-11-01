// +build linux,386
package main

import (
	"flag"
	"fmt"
	"github.com/krmit/helloGo/html"
)

func main() {

	//Parse commadline flags.
	repeat := 0
	text := ""
	headline := true
	flag.IntVar(&repeat, "r", 1, "Number of time to writte message.")
	flag.StringVar(&text, "t", "Hello Go!", "Test that message should containe.")
	flag.BoolVar(&headline, "hl", true, "If it should be a headline")
	flag.Parse()

	for i := 0; i < repeat; i++ {
		if headline {
			fmt.Printf(html.Headline(text))
		} else {
			fmt.Printf(text)
		}
	}
}
