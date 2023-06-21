package main

import "fmt"

const polindrom string = "шалаш"

func main() {
	var wordRunes = []rune(polindrom)
	var j = len(wordRunes)
	for i := 0; i <= len(wordRunes)-1; i++ {
		j--
		if wordRunes[i] == wordRunes[j] {
			continue
		}
		fmt.Printf("Word not polindrom!")
		break
	}
	if j == 0 {
		fmt.Printf("Word polindrom!")
	}
}