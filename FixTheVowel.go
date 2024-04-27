package reloaded

import (
	"strings"
)

func fixTheVowel(text string) string {
	words := strings.Split(text, " ")
	data := words
	// data := text
	res := ""
	for idx, word := range data {
		if idx < len(data)-1 {
			item := data[idx+1]

			if word == "a" {
				if item[0] == 'e' || item[0] == 'o' || item[0] == 'u' || item[0] == 'h' || item[0] == 'a' || item[0] == 'i' {
					res += "an "
				} else if item[0] == 'E' || item[0] == 'O' || item[0] == 'U' || item[0] == 'H' || item[0] == 'A' || item[0] == 'I' {
					res += "an "
				}
			} else if word == "A" {
				if item[0] == 'e' || item[0] == 'o' || item[0] == 'u' || item[0] == 'h' || item[0] == 'a' || item[0] == 'i' {
					res += "An "
				} else if item[0] == 'E' || item[0] == 'O' || item[0] == 'U' || item[0] == 'H' || item[0] == 'A' || item[0] == 'I' {
					res += "An "
				}
			} else {
				res += word + " "
			}
		}
		if idx == len(data)-1 {
			res += word
		}
	}
	return res
}
