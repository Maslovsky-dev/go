package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builder strings.Builder
	var isShielding = false
	runes := []rune(s) //Приводим строку к массиву рун, чтобы обрабатывать все символы Unicode
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\\' {
			if i == len(runes)-1 {
				return "", ErrInvalidString
			}
			isShielding = true
			builder.WriteRune(runes[i+1]) //Неэкранированный символ
			i++
			continue
		}
		if unicode.IsDigit(runes[i]) { //Проверка на число
			if i == 0 {
				return "", ErrInvalidString
			}
			if unicode.IsDigit(runes[i-1]) && !isShielding { //Два числа подряд, кроме случая с экранированием
				return "", ErrInvalidString
			}
			num, _ := strconv.Atoi(string(runes[i]))
			if num == 0 {
				str := builder.String()
				builder.Reset()
				builder.WriteString(str[:len(str)-1])
			}
			if num > 1 {
				builder.WriteString(strings.Repeat(string(runes[i-1]), num-1))
			}
			isShielding = false
			continue
		}
		builder.WriteRune(runes[i])
	}
	return builder.String(), nil
}
