package lineman

import "bytes"

// Интерфейс, реализующий ряд функциональных возможностей для разных типов парсеров
type Liner interface {
	IncPos()                   // Вперёд на один символ
	IsLetter() int             // Проверяем, является ли текущий символ буквенным и возвращаем его длину (нулевая длина - символ не является буквенным)
	IsSpace() bool             // Проверяем, является ли символ пробельным
	IsEndLine() bool           // Проверяем, является ли текущий символ окончанием строки
	CheckEndLine(ch byte) bool // Проверя, является ли переданный символ окончанием строки
}

// Инициализирует овый обходчик массива байтов
func NewByteLine(src []byte) *ByteLine {
	res := &ByteLine{src: src}
	res.liner = res
	return res
}

// Структура обходчика массива байтов
type ByteLine struct {
	src   []byte
	pos   int
	liner Liner
}

func (s *ByteLine) SetLiner(liner Liner) {
	s.liner = liner
}

// Сдвигает позицию вперёд на 1 символ
func (s *ByteLine) IncPos() { s.pos++ }

// Возвращает символ на текущей позиции
func (s *ByteLine) Char() byte {
	if s.pos < len(s.src) {
		return s.src[s.pos]
	} else {
		return 0
	}
}

func (s *ByteLine) NextChar() byte {
	if s.pos < len(s.src)-1 {
		return s.src[s.pos+1]
	} else {
		return 0
	}
}

func (s *ByteLine) PrevChar() byte {
	if s.pos > 0 && s.pos < len(s.src) {
		return s.src[s.pos-1]
	} else {
		return 0
	}
}

// Проверка соответствия текущей позиции концу документа
func (s *ByteLine) IsEndDocument() bool { return len(s.src) <= s.pos }

// Проверка соответствия текущего символа концу документа или заверщению оператора
func (s *ByteLine) IsEndLine() bool {
	if !s.IsEndDocument() {
		return CheckEndLine(s.src[s.pos])
	}
	return true
}

func (s *ByteLine) CheckEndLine(ch byte) bool {
	return CheckEndLine(ch)
}

// Сдвигает позицию вперёд до первого непробельного символа
func (s *ByteLine) PassSpaces() {
	for !s.IsEndDocument() && s.liner.IsSpace() {
		s.liner.IncPos()
	}
}

// Пропускает все пробелы и символы конца строки
func (s *ByteLine) PassEndLines() {
	for !s.IsEndDocument() && (s.liner.IsSpace() || s.liner.IsEndLine()) {
		s.liner.IncPos()
	}
}

// Проверяем, является ли текущий символ буквенным и возвращает его длину (в юникоде). Нулевая длина - символ не является буквенным
func (s *ByteLine) IsLetter() int {
	return CheckUnicodeLetter(s.src[s.pos:])
}

// Проверяем, является ли текущий символ пробельным
func (s *ByteLine) IsSpace() bool {
	return s.src[s.pos] == ' '
}

// Читает слово (в юикоде)
func (s *ByteLine) ReadWord() (res []byte, check bool) {
	pos := s.pos
	for !s.IsEndDocument() {
		if size := CheckUnicodeLetter(s.src[s.pos:]); size > 0 {
			for i := 0; i < size; i++ {
				s.liner.IncPos()
			}
		} else {
			break
		}
	}
	res = s.src[pos:s.pos]
	check = len(res) > 0
	return
}

// Двигаемся вперёд по пробелам и читаем слово
func (s *ByteLine) ReadWordSpaces() ([]byte, bool) {
	s.PassSpaces()
	return s.ReadWord()
}

// Двигаемся вперёд на offset позиций
func (s *ByteLine) ForwardPos(offset int) {
	for i := 0; i < offset; i++ {
		s.liner.IncPos()
	}
}

// Взвращает срез от текущего элемента до конца строки (без сдвига позиции)
func (s *ByteLine) EndLineContent() []byte {
	pos := s.pos
	for pos < len(s.src) && !s.liner.CheckEndLine(s.src[pos]) {
		pos++
	}
	return s.src[s.pos:pos]
}

// Двигаемся вперёд ндо тех пор, пока не встретим ch. Если ch не встретили, вернётся false
func (s *ByteLine) ToChar(ch byte) bool {
	for !s.IsEndDocument() {
		if s.src[s.pos] == ch {
			return true
		}
		s.liner.IncPos()
	}
	return false
}

// Возвращает срез до текущей позиции
func (s *ByteLine) Left() (res []byte) {
	if s.pos <= len(s.src) {
		res = s.src[:s.pos]
	} else {
		return s.src
	}
	return
}

// Возаращает срез с текущей позиции до конца
func (s *ByteLine) Right() (res []byte) {
	if s.pos < len(s.src) {
		res = s.src[s.pos:]
	}
	return
}

// Возвращает текущую позицию
func (s *ByteLine) Pos() int {
	return s.pos
}

// Возвращает длину среза контента
func (s *ByteLine) Len() int {
	return len(s.src)
}

// Возвращает длину среза с текущей позиции до конца
func (s *ByteLine) LeftLen() int {
	return len(s.src[s.pos:])
}

// Двигается врерёд до совпадения. Если совпадений до конца среза не найдено, вернёт false
func (s *ByteLine) ToSlice(match []byte) bool {
	if len(match) == 0 {
		return false
	}
	for !s.IsEndDocument() {
		if s.src[s.pos] == match[0] && bytes.Index(s.src[s.pos:], match) == 0 {
			return true
		} else {
			s.IncPos()
		}
	}
	return false
}

// Двигаемся вперед до совпадения, затем ещё вперёд на длину среза совпадения
func (s *ByteLine) ToSlicePost(match []byte) (res bool) {
	if res = s.ToSlice(match); res {
		s.ForwardPos(len(match))
	}
	return
}

func (s *ByteLine) PosMatchSlice(match []byte) bool {
	if len(match) == 0 {
		return false
	}
	return s.src[s.pos] == match[0] && bytes.Index(s.src[s.pos:], match) == 0
}

func (s *ByteLine) MatchSliceIndexPos(match []byte) (pos int) {
	return bytes.Index(s.src[s.pos:], match)
}

func (s *ByteLine) MatchSliceIndex(match []byte) (pos int) {
	if pos = s.MatchSliceIndexPos(match); pos >= 0 {
		pos += s.pos
	}
	return
}
