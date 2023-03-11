package main

import "fmt"

func main() {
	assignment()
}

func assignment() {

	var characters = []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}

	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			for index, v := range characters {
				fmt.Printf("character %U '%s' starts at byte position %d\n", v, string(v), index*2)
			}
		}
		fmt.Printf("Nilai j = %d\n", j)
	}
}