package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(src string) (string, error) {

	var res string

	runes := []rune(src)

	var tmp string
	var count int
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			count, _ = strconv.Atoi(string(runes[i]))
			res += strings.Repeat(tmp, count)
		} else {
			tmp = string(runes[i])
		}
	}

	return res, nil
}
