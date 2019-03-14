package lineman

import (
	"log"
	"testing"
)

func TestPassSpaces(t *testing.T) {
	src := []byte("vdbvdhvdfvh{{1234567890")
	p := NewByteLine(src)
	p.ForwardPos(2)
	log.Println(p.ToChar('t'))
	log.Println(p.MatchSliceIndex([]byte("{{")))
	log.Println(p.MatchSliceIndexPos([]byte("{{")))
	log.Println(string(p.Right()))
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
