package reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func EditText(text string) string {
	// the name of the variable should have a meaning
	words := strings.Split(text, " ")

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
			words[idx-1] = strings.Title(words[idx-1])
			words[idx] = ""
		} else if word == "(low," {
			nbrStr := (strings.TrimRight(words[idx+1], "),"))
			nbr, _ := strconv.Atoi(nbrStr)
			for l := 0; l < nbr; l++ {
				words[idx-nbr+l] = strings.ToLower(words[idx-nbr+l])
				fmt.Println(words[idx-nbr+l])
			}
		} else if word == "(cap," {
			nbrStr := (strings.TrimRight(words[idx+1], "),"))
			nbr, _ := strconv.Atoi(nbrStr)
			for l := 0; l < nbr; l++ {
				words[idx-nbr+l] = strings.Title(words[idx-nbr+l])
				fmt.Println(words[idx-nbr+l])
			}
		}

	}
	return strings.Join(words, " ")
}
