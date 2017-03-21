package trandword

import (
	"bufio"
	"io"
	"sort"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

type Trandword struct {
	Vocabs  []Vocab
	scanner *bufio.Scanner
}

func NewTrandword(reader io.Reader) *Trandword {
	return &Trandword{
		Vocabs:  make([]Vocab, 0),
		scanner: bufio.NewScanner(reader),
	}
}

func (tw *Trandword) Analyze() {
	t := tokenizer.New()
	nounCount := float32(0)
	for tw.scanner.Scan() {
		if err := tw.scanner.Err(); err != nil {
			panic(err)
		}
		s := tw.scanner.Text()
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
				for i, _ := range tw.Vocabs {
					if tw.Vocabs[i].Word == token.Surface {
						tw.Vocabs[i].Freq++
						flag = true
						break
					}
				}
				if !flag {
					tw.Vocabs = append(tw.Vocabs, Vocab{
						Word: token.Surface,
						Freq: 1,
					})
				}
				nounCount++
			}
		}
	}
	sort.Slice(tw.Vocabs, func(i, j int) bool { return tw.Vocabs[i].Freq > tw.Vocabs[j].Freq })
}
