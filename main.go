package main

import (
	"flag"
	"fmt"
	"github.com/bananaumai/snake2camel/converter"
)

var lower *bool

func init() {
	lower = flag.Bool("lower", false, "convert to lowerCamelCase")
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		fmt.Println(converter.Convert(arg, *lower))
	}

}
