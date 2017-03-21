package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

type Vocab struct {
	Word string
	Freq int
}

func main() {

	t := tokenizer.New()
	scanner := bufio.NewScanner(os.Stdin)

	vocabs := make([]Vocab, 0)
	nounCount := float32(0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		s := scanner.Text()
		s = strings.TrimSuffix(s, "\"")
		s = strings.TrimPrefix(s, "\"")
		tokens := t.Tokenize(s)
		for _, token := range tokens {
			if token.Class == tokenizer.DUMMY {
				continue
			}
			features := token.Features()
			if features[0] == "名詞" || features[0] == "形容詞" {
				flag := false
				for i, _ := range vocabs {
					if vocabs[i].Word == token.Surface {
						vocabs[i].Freq++
						flag = true
						break
					}
				}
				if !flag {
					vocabs = append(vocabs, Vocab{
						Word: token.Surface,
						Freq: 1,
					})
				}
				nounCount++
			}
		}
	}

	sort.Slice(vocabs, func(i, j int) bool { return vocabs[i].Freq > vocabs[j].Freq })
	fmt.Println("word\tfrequency\tprob")
	for _, v := range vocabs {
		fmt.Printf("%s\t%d\t%0.3f\n", v.Word, v.Freq, ((float32)(v.Freq) / nounCount))
	}

}
