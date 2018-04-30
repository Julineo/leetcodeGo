package main

import (
	"fmt"
)

func main () {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) int {
	morseDic := map[string]string{"a": ".-",
				"b": "-...",
				"c": "-.-.",
				"d": "-..",
				"e": ".",
				"f": "..-.",
				"g": "--.",
				"h": "....",
				"i": "..",
				"j": ".---",
				"k": "-.-",
				"l": ".-..",
				"m": "--",
				"n": "-.",
				"o": "---",
				"p": ".--.",
				"q": "--.-",
				"r": ".-.",
				"s": "...",
				"t": "-",
				"u": "..-",
				"v": "...-",
				"w": ".--",
				"x": "-..-",
				"y": "-.--",
				"z": "--.."}
	morseCode := make(map[string] int)
	for _, word := range(words) {
		newWord := ""
		for _, r := range(word) {
			newWord = newWord + morseDic[string(r)]
		}
		morseCode[newWord] = morseCode[newWord] + 1
	}
	return len(morseCode)
}
