package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(arg string) (string, error) {
	out := strings.Builder{}
	var prev rune
	for _, r := range arg {
		m, err := strconv.Atoi(string(r))
		switch {
		case err == nil && prev != 0: // цифра после буквы(не цифры)
			out.WriteString(strings.Repeat(string(prev), m))
			prev = 0
		case err != nil && prev == 0: // первая буква или после цифры
			prev = r
		case err != nil && prev != 0: // буква после буквы
			out.WriteRune(prev)
			prev = r
		case err == nil && prev == 0: // цифра первая в строке или после цифры
			return "", ErrInvalidString
		}
	}
	if prev != 0 {
		out.WriteRune(prev)
	}
	return out.String(), nil
}
