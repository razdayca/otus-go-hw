package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(unpackStr string) (string, error) {
	var final strings.Builder
	final.Grow(32)
	for i, symbol := range []rune(unpackStr) {
		if unicode.IsDigit(symbol) && !unicode.IsDigit([]rune(unpackStr)[i+1]) {
			count, _ := strconv.Atoi(string(symbol))
			if count > 0 && i > 0 {
				final.WriteString(strings.Repeat(string([]rune(unpackStr)[i-1]), count-1))
			} else if i > 0 {
				_, buf, _ := strings.Cut(final.String(), string([]rune(unpackStr)[i-1]))
				final.Reset()
				final.WriteString(buf)
			} else {
				return "", ErrInvalidString
			}

		} else if unicode.IsDigit(symbol) && unicode.IsDigit([]rune(unpackStr)[i+1]) {
			return "", ErrInvalidString
		} else {
			final.WriteString(string(symbol))
		}
	}
	return final.String(), nil
}
