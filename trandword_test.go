package trandword

import (
	"bytes"
	"testing"
)

func TestNewTrandword(t *testing.T) {
	reader := bytes.NewBufferString("foo\n")
	if tw := NewTrandword(reader); tw == nil {
		t.Fatalf("NewTrandword() : nil expected not nil")
	}

}
