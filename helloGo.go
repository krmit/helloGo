// +build linux,386
package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v1"
	"github.com/krmit/helloGo/html"
)

func main() {

	//Parse commadline flags.
	repeat := kingpin.Flag("r", "Number of time to writte message.").Default("1").Int()
	text := kingpin.Flag("t", "Text that the message should containe.").Default("Hello Go!").String()
	headline := kingpin.Flag("hl", "If it should be a headline").Default("false").Bool()
	
	kingpin.Parse()
	
	for i := 0; i < *repeat; i++ {
		if *headline {
			fmt.Printf(html.Headline(*text))
		} else {
			fmt.Printf(*text)
		}
	}
}
