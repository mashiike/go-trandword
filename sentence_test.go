package trandword_test

import (
	"bytes"
	"testing"

	"github.com/mashiike/go-trandword"
)

func TestNewSentenceDecorder(t *testing.T) {
	reader := bytes.NewBufferString(``)
	if sr := trandword.NewSentenceDecorder(reader); sr == nil {
		t.Fatalf("NewSentenceDecorder() : nil expected not nil")
	}
}

func TestJsonRead(t *testing.T) {

	reader := bytes.NewBufferString(`
		{"document_id":1 , "message" : "hoge hoge fuga"}
	`)
	sr := trandword.NewSentenceDecorder(reader)
	s, err := sr.Decode()
	if err != nil {
		t.Fatalf("TestJsonRead: %#v expected %#v", err, nil)
	}
	if s.Message != "hoge hoge fuga" {
		t.Fatalf("TestJsonRead: not match message %#v", s)
	}
	if s.DocumentId != 1 {
		t.Fatalf("TestJsonRead: not match DocumentID  %#v expected DocumentID 1", s)
	}
	if s.Timestamp.IsZero() {
		t.Fatalf("TestJsonRead: timestamp is zero  %#v", s)
	}

}
