package main

import (
	"fmt"
	"os"

	"github.com/mashiike/go-trandword"
)

func main() {

	tw := trandword.NewTrandword(os.Stdin)
	tw.Analyze()

	fmt.Println("word\tfrequency")
	for _, v := range tw.Vocabs {
		fmt.Printf("%s\t%d\n", v.Word, v.Freq)
	}

}
