package reloaded

import (
	// "fmt"
	"strconv"
	"strings"
)

func EditText(text string) string {
	// the name of the variable should have a meaning

	words := strings.Split(fixTheVowel(text), " ")

	for idx, word := range words {

		if word == "(up)" {
			words[idx-1] = strings.ToUpper(words[idx-1])
			words[idx] = ""

		} else if word == "(low)" {

			words[idx-1] = strings.ToLower(words[idx-1])
			words[idx] = ""
		} else if word == "(bin)" {
			decimal, _ := strconv.ParseInt(words[idx-1], 2, 64)

			words[idx-1] = strconv.FormatInt(decimal, 10)
			words[idx] = ""
		} else if word == "(hex)" {
			decimal, _ := strconv.ParseInt(words[idx-1], 16, 64)
			words[idx-1] = strconv.FormatInt(decimal, 10)
			words[idx] = ""
		} else if word == "(cap)" {
			if idx == 0 {
				return ""
			}
			words[idx-1] = toCapitalize(words[idx-1])
			words[idx] = ""
		} else if word == "(low," {

			nbrStr := (strings.TrimRight(words[idx+1], "),"))
			nbr, _ := strconv.Atoi(nbrStr)
			for l := 0; l < nbr; l++ {
				words[idx-nbr+l] = strings.ToLower(words[idx-nbr+l])
				words[idx] = ""
				words[idx+1] = ""
				// fmt.Println(words[idx-nbr+l])
			}
		} else if word == "(cap," {
			if idx+1 == len(words) {
				return ""
			}
			nbrStr := (strings.TrimRight(words[idx+1], ")!:;,"))
			nbr, _ := strconv.Atoi(nbrStr)
			for l := 0; l < nbr; l++ {
				words[idx-nbr+l] = toCapitalize(words[idx-nbr+l])
				words[idx] = ""
				words[idx+1] = ""
				// fmt.Println(words[idx-nbr+l])
			}
		} else if word == "(up," {
			nbrStr := (strings.TrimRight(words[idx+1], ")!:;,"))
			nbr, _ := strconv.Atoi(nbrStr)
			for l := 0; l < nbr; l++ {
				words[idx-nbr+l] = strings.ToUpper(words[idx-nbr+l])
				words[idx] = ""
				words[idx+1] = ""
				// fmt.Println(words[idx-nbr+l])
			}
		}
		// else if word == "(hex," {
		// 	decimal, _ := strconv.ParseInt(words[idx-1], 16, 64)
		// 	words[idx-1] = strconv.FormatInt(decimal, 10)
		// 	words[idx] = ""
		// 	words[idx+1] = ""
		// }

	}
	return strings.Join(words, " ")
}
