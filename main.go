package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	var input = flag.String("input", "data/default.txt", "input file")
	flag.Parse()

	fp, err := os.Open(*input)
	if err != nil {
		panic(err)
	}

	t := tokenizer.New()
	scanner := bufio.NewScanner(fp)
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
