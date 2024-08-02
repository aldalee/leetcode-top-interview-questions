package main

import (
	"fmt"
	"testing"
)

func getInfo(s []int) {
	fmt.Printf("s=%v\tlen(s)=%d\tcap(s)=%d\n", s, len(s), cap(s))
}

func TestSlice(t *testing.T) {
	s := make([]int, 5)
	getInfo(s)
	s = append(s, 1)
	s = append(s, 2)
	getInfo(s)
}
