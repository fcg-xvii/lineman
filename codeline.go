package lineman

func NewCodeLine(src []byte) *CodeLine {
	res := &CodeLine{NewDocLine(src)}
	res.liner = res
	return res
}

type CodeLine struct {
	*DocLine
}

func checkFirsNameChar(src []byte) (size int) {
	if len(src) > 0 {
		if src[0] == '_' {
			size = 1
		} else {
			size = checkUnicodeLetter(src)
		}
	}
	return
}

func checkBodyNameChar(src []byte) (size int) {
	if size = checkFirsNameChar(src); size == 0 && len(src) > 0 && checkNumber(src[0]) {
		size = 1
	}
	return
}

// В этой реализации пробельным символом является так же табуляция
func (s *CodeLine) IsSpace() bool {
	return s.src[s.pos] == '\t' || s.DocLine.IsSpace()
}

func (s *CodeLine) IsEndLine() bool {
	return s.src[s.pos] == ';' || s.DocLine.IsEndLine()
}

func (s *CodeLine) CheckEndLine(ch byte) bool {
	return s.src[s.pos] == ';' || s.DocLine.CheckEndLine(ch)
}

func (s *CodeLine) ReadName() (res []byte, check bool) {
	if !s.IsEndDocument() {
		s.SetupMark()
		var size int
		if size = checkFirsNameChar(s.src[s.pos:]); size > 0 {
			s.ForwardPos(size)
		}
		for {
			if size = checkBodyNameChar(s.src[s.pos:]); size > 0 {
				s.ForwardPos(size)
			} else {
				break
			}
		}
		res = s.MarkVal(0)
		check = len(res) > 0
	}
	return
}

func (s *CodeLine) ReadNameSpaces() ([]byte, bool) {
	s.PassSpaces()
	return s.ReadName()
}
