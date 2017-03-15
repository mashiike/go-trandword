package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	var input = flag.String("input", "data/default.txt", "input file")
	var modef = flag.Bool("mode-file", false, "input mode file read flag")
	flag.Parse()

	var reader io.Reader

	if *modef {
		var err error
		reader, err = os.Open(*input)
		if err != nil {
			panic(err)
		}
	} else {
		reader = os.Stdin
	}

	t := tokenizer.New()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		tokens := t.Tokenize(scanner.Text())
		for _, token := range tokens {
			if token.Class == tokenizer.DUMMY {
				// BOS: Begin Of Sentence, EOS: End Of Sentence.
				fmt.Printf("%s\n", token.Surface)
				continue
			}
			features := strings.Join(token.Features(), ",")
			fmt.Printf("%s\t%v\n", token.Surface, features)
		}
	}

}
