package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrimRune(t *testing.T) {
	tests := []struct {
		trim     []rune
		cut      int
		expected []rune
	}{
		{[]rune{'a', 'b', 'c'}, 2, []rune{'a'}},
		{[]rune{'a', 'b', 'c'}, 3, []rune{}},
		{[]rune{'a', 'b', 'c'}, 4, []rune{}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("TrimRune", func(t *testing.T) {
			result := TrimRuneSlice(tc.trim, tc.cut)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestRepeatRune(t *testing.T) {
	tests := []struct {
		repeat   rune
		times    int
		expected []rune
	}{
		{repeat: 'a', times: 3, expected: []rune{'a', 'a', 'a'}},
		{repeat: 'a', times: 0, expected: []rune{}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("RepeatRune", func(t *testing.T) {
			result := RepeatRune(tc.repeat, tc.times)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
