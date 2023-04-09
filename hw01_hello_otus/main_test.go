package main

import "testing"

func TestReverse(t *testing.T) {
	str := "Hello, OTUS!"
	if reverse(str) != "!SUTO ,olleH" {
		t.Error(`reverse("Hello, OTUS!") != "!SUTO ,olleH"`)
	}
}
