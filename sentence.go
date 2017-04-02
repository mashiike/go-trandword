package trandword

import (
	"encoding/json"
	"io"
	"time"
)

type Sentence struct {
	Timestamp  int64  `json:"timestamp"`
	DocumentId int    `json:"document_id"`
	Message    string `json:"message"`
}

type SentenceDecoder interface {
	Decode() (*Sentence, error)
}

type JsonSentenceDecoder struct {
	Decoder *json.Decoder
}

func NewSentenceDecorder(reader io.Reader) SentenceDecoder {
	return &JsonSentenceDecoder{
		Decoder: json.NewDecoder(reader),
	}
}

func (jsd *JsonSentenceDecoder) Decode() (*Sentence, error) {

	s := Sentence{
		Timestamp: time.Now().Unix(),
	}
	if err := jsd.Decoder.Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil

}
