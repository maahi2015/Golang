package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(7, 5)
	want := 12
	if want != got {
		println("Error")
	}
}
func TestSubtract(t *testing.T) {
	got := Subtract(2, 1)
	want := 1
	if want != got {
		println("Error")
	}
}
