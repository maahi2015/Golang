package main

import (
	"log"
	"testing"
)

func TestHelloworld(t *testing.T) {
	got := "hi"
	want := Helloworld("hi")
	if got != want {
		log.Println("error")
	}
}
