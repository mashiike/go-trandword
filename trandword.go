package trandword

type Trandword struct {
	Vocabs []Vocab
}

func NewTrandword() *Trandword {
	return &Trandword{
		Vocabs: make([]Vocab, 0),
	}
}
