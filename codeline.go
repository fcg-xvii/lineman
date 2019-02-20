package lineman

func NewCodeLine(src []byte) *CodeLine {
	res := &CodeLine{NewDocLine(src)}
	res.liner = res
	return res
}

type CodeLine struct {
	*DocLine
}

func CheckFirsNameChar(src []byte) (size int) {
	if len(src) > 0 {
		if src[0] == '_' {
			size = 1
		} else {
			size = CheckUnicodeLetter(src)
		}
	}
	return
}

func CheckBodyNameChar(src []byte) (size int) {
	if size = CheckFirsNameChar(src); size == 0 && len(src) > 0 && CheckNumber(src[0]) {
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
	return ch == ';' || s.DocLine.CheckEndLine(ch)
}

func (s *CodeLine) ReadName() (res []byte, check bool) {
	if !s.IsEndDocument() {
		s.SetupMark()
		var size int
		if size = CheckFirsNameChar(s.src[s.pos:]); size > 0 {
			s.ForwardPos(size)
		}
		for {
			if size = CheckBodyNameChar(s.src[s.pos:]); size > 0 {
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
