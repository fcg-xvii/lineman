package lineman

import (
	_ "log"
	"testing"
)

func TestPassSpaces(t *testing.T) {
	src := []byte("          ")
	p := NewByteLine(src)
	p.PassSpaces()
}

func TestDocLine(t *testing.T) {
	src := []byte("      CoolAway")
	p := NewDocLine(src)
	w, _ := p.ReadWordSpaces()
	t.Log(string(w))
}

func TestCodeLine(t *testing.T) {
	src := []byte("x = 10;   __CoolAway18  vdf")
	p := NewCodeLine(src)
	//t.Log("END_LINE_CONTENT", string(p.EndLineContent()))
	w, _ := p.ReadNameSpaces()
	t.Log(string(w))
}
