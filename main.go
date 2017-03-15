package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {

	t := tokenizer.New()
	scanner := bufio.NewScanner(os.Stdin)

	vcab := make(map[string]int)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		tokens := t.Tokenize(scanner.Text())
		for _, token := range tokens {
			if token.Class == tokenizer.DUMMY {
				continue
			}
			features := token.Features()
			if features[0] == "名詞" {
				vcab[token.Surface]++
			}
		}
	}

	fmt.Println("word,frequency")
	for k, v := range vcab {
		fmt.Printf("\"%s\",\"%d\"\n", k, v)
	}

}
