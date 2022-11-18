package main

import (
	"log"
	"testing"
)

func TestHelloworld(t *testing.T) {
	got := "hi"
	want := Helloworld("hello")
	if got != want {
		log.Println("error")
	}
}
