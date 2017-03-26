package trandword_test

import (
	"bytes"
	"testing"

	"github.com/mashiike/go-trandword"
)

func TestNewTrandword(t *testing.T) {
	reader := bytes.NewBufferString("foo\n")
	if tw := trandword.NewTrandword(reader); tw == nil {
		t.Fatalf("NewTrandword() : nil expected not nil")
	}
}
