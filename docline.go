package lineman

import "fmt"

func NewDocLine(src []byte) *DocLine {
	res := &DocLine{
		NewByteLine(src),
		1, 1,
		&mark{0, 1, 1},
	}
	res.liner = res
	return res
}

type mark struct {
	pos, line, linePos int
}

// Структура обходчика многострочного документа
type DocLine struct {
	*ByteLine
	line, linePos int
	_mark         *mark
}

// Сдвигаем указатель вперёд на одну позицию, устанавливаем значения строки и позиции в строке
func (s *DocLine) IncPos() {
	s.ByteLine.IncPos()
	if !s.IsEndDocument() && s.src[s.pos] == '\n' {
		s.line++
		s.linePos = 1
	} else {
		s.linePos++
	}
}

// Установка метки на текущей позиции
func (s *DocLine) SetupMark() {
	s._mark.pos, s._mark.line, s._mark.linePos = s.pos, s.line, s.linePos
}

// Откат к позиции согласно установленной метке (с возможным вперёд)
func (s *DocLine) RollbackMark(forward int) {
	s.pos, s.line, s.linePos = s._mark.pos, s._mark.line, s._mark.linePos
	for i := 0; i < forward; i++ {
		s.liner.IncPos()
	}
}

// Получение среза от установленной метки до текущей позиции со смещением
func (s *DocLine) MarkVal(rOffset int) (res []byte) {
	if s._mark != nil {
		if s._mark.pos < s.pos-rOffset {
			res = s.src[s._mark.pos : s.pos-rOffset]
		}
	}
	return
}

// Получение строки от установленной метки со смещением
func (s *DocLine) MarkValString(rOffset int) string {
	return string(s.MarkVal(rOffset))
}

// Позиция указателя (если не установлен, будет возвращён 0)
func (s *DocLine) MarkPos() (pos int) {
	if s._mark != nil {
		pos = s._mark.pos
	}
	return
}

func (s *DocLine) MarkLine() (line int) {
	if s._mark != nil {
		line = s._mark.line
	}
	return
}

func (s *DocLine) MarkLinePos() (linePos int) {
	if s._mark != nil {
		linePos = s._mark.linePos
	}
	return
}

func (s *DocLine) Line() int {
	return s.line
}

func (s *DocLine) LinePos() int {
	return s.linePos
}

// Установка ошибки (к тексту добавляется информация о номере текущей строки и позиции)
func (s *DocLine) InitError(text string) (err error) {
	err = fmt.Errorf("%s [ line: %d, position: %d ]", text, s.line, s.linePos)
	return
}
