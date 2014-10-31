// +build linux,386
package main

import (
	"fmt"
	"github.com/krmit/helloGo/html"
	)

func main() {
	fmt.Printf(html.Headline("Hello GO!"))
}
