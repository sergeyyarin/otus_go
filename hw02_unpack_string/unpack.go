package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func RepeatRune(r rune, n int) (repeated []rune) {
	if n <= 0 {
		return []rune{}
	}
	for i := 0; i < n; i++ {
		repeated = append(repeated, r)
	}
	return repeated
}

func TrimRuneSlice(r []rune, n int) []rune {
	if n <= 0 {
		return []rune{}
	}
	l := len(r)
	if l <= n {
		return []rune{}
	}
	return r[:l-n]
}

func runeToInt(r rune) int {
	val, _ := strconv.Atoi(string(r))
	return val
}

func Unpack(s string) (result string, err error) {
	if len(s) == 0 {
		return "", nil
	}
	packed := []rune(s)
	if unicode.IsDigit(packed[0]) { // valid string doesn't start with a digit
		return "", ErrInvalidString
	}
	var last rune
	var unpacked []rune
	for _, v := range packed {
		if unicode.IsDigit(v) {
			if unicode.IsDigit(last) { // no more than one digit allowed
				return "", ErrInvalidString
			}
			unpacked = TrimRuneSlice(unpacked, 1)
			times := runeToInt(v)
			unpacked = append(unpacked, RepeatRune(last, times)...)
		} else {
			unpacked = append(unpacked, v)
		}
		last = v
	}
	result = string(unpacked)
	return result, nil
}
