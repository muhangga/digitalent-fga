package main

import (
	"fmt"
	"strings"
)

func main() {
	assignment()
}

func assignment() {
	words := "selamat malam"

	count := make(map[string]int)
	splitWord := strings.Split(words, "")

	for _, word := range splitWord {
		for _, char := range word {
			string := string(char)
			count[string]++
			fmt.Printf("%c\n", char)
		}
	}
	fmt.Printf("%v\n", count)
}
