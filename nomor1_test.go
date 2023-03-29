package jastis

import (
	"fmt"
	"testing"
)

func change(A int, B int) (int, int) {
	a, b := B, A

	return a, b
}

func TestChange(t *testing.T) {
	A := 100
	B := 6
	fmt.Println(change(A, B))
}
