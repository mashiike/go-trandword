package trandword

import (
	"io"
	"sort"

	"github.com/ikawaha/kagome/tokenizer"
)

type Trandword struct {
	Vocabs  []Vocab
	Decoder SentenceDecoder
}

func NewTrandword(reader io.Reader) *Trandword {
	return &Trandword{
		Vocabs:  make([]Vocab, 0),
		Decoder: NewSentenceDecorder(reader),
	}
}

func (tw *Trandword) Analyze() {
	t := tokenizer.New()
	for {
		s, err := tw.Decoder.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		tokens := t.Tokenize(s.Message)
		for _, token := range tokens {
			if token.Class == tokenizer.DUMMY {
				continue
			}
			features := token.Features()
			if features[0] == "名詞" ||
				features[0] == "形容詞" ||
				features[0] == "動詞" {
				flag := false
				word := features[6]
				for i, _ := range tw.Vocabs {
					if tw.Vocabs[i].Word == word {
						tw.Vocabs[i].Freq++
						flag = true
						break
					}
				}
				if !flag {
					tw.Vocabs = append(tw.Vocabs, Vocab{
						Word: word,
						Freq: 1,
					})
				}
			}
		}

	}
	sort.Slice(tw.Vocabs, func(i, j int) bool { return tw.Vocabs[i].Freq > tw.Vocabs[j].Freq })
}
