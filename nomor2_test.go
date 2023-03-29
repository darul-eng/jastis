package jastis

import (
	"fmt"
	"testing"
)

func reverse(s string) string {
	var String string
	for i := len(s) - 1; i >= 0; i-- {
		String += string(s[i])
	}

	return String
}

func TestReverse(t *testing.T) {
	String := "jatis"
	fmt.Println(reverse(String))
}
