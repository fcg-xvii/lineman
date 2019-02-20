package lineman

import (
	"unicode"
	"unicode/utf8"
)

// Проверка соответствия символу завершения оператора (ему соответствует символ конца строки или ";"
func CheckEndLine(ch byte) bool { return ch == '\n' }

// Проверка соответствия байта буквенному значению (соответствует регулярке [0-9A-Za-z])
func CheckLetter(ch byte) bool { return (ch >= 97 && ch <= 122) || (ch >= 65 && ch <= 90) }

// Проверка, является ли байт числовым значением (соответствует регулярке [0-9])
func CheckNumber(ch byte) bool { return ch >= 48 && ch <= 57 }

func CheckUnicodeLetter(src []byte) (size int) {
	var r rune
	r, size = utf8.DecodeRune(src)
	if !unicode.IsLetter(r) {
		size = 0
	}
	return
}
