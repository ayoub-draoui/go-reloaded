package reloaded

import (
	// "fmt"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EditText(text string) string {
	// the name of the variable should have a meaning

	words := strings.Fields(text)

	for idx, word := range words {
		if word == "(up)" {
			if idx == 0 {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			words[idx-1] = strings.ToUpper(words[idx-1])
			words[idx] = ""

		} else if word == "(low)" {
			if idx == 0 {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			words[idx-1] = strings.ToLower(words[idx-1])
			words[idx] = ""
		} else if word == "(bin)" {
			if idx == 0  {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			decimal, _ := strconv.ParseInt(words[idx-1], 2, 64)

			words[idx-1] = strconv.FormatInt(decimal, 10)
			words[idx] = ""
		} else if word == "(hex)" {
			if idx == 0 {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			decimal, _ := strconv.ParseInt(words[idx-1], 16, 64)
			words[idx-1] = strconv.FormatInt(decimal, 10)
			words[idx] = ""
		} else if word == "(cap)" {
			if idx == 0  {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			words[idx-1] = toCapitalize(words[idx-1])
			words[idx] = ""
		} else if word == "(low," {
			if idx == 0 || idx < len(words) {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			nbrStr := (strings.TrimRight(words[idx+1], ")"))
			nbr, _ := strconv.Atoi(nbrStr)
			for l := 0; l < nbr; l++ {
				if nbr > idx {
					fmt.Println("Err: the number next the function should at least equal the range of str ")
					os.Exit(1)
				}
				words[idx-nbr+l] = strings.ToLower(words[idx-nbr+l])
				words[idx] = ""
				words[idx+1] = ""
				// fmt.Println(words[idx-nbr+l])
			}
		} else if word == "(cap," {
			if idx == 0 || idx < len(words) {

				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			nbrStr := (strings.TrimRight(words[idx+1], ")!:;,"))
			nbr, _ := strconv.Atoi(nbrStr)
			for l := 0; l < nbr; l++ {
				if nbr > idx {
					fmt.Println("Err: bla bla bla...")
					os.Exit(1)
				}
				words[idx-nbr+l] = toCapitalize(words[idx-nbr+l])
				words[idx] = ""
				words[idx+1] = ""
				// fmt.Println(words[idx-nbr+l])
			}
		} else if word == "(up," {
			if idx == 0 || idx < len(words) {

				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			nbrStr := (strings.TrimRight(words[idx+1], ")!:;,"))
			nbr, _ := strconv.Atoi(nbrStr)
			for l := 0; l < nbr; l++ {
				if nbr > idx {
					fmt.Println("Err: bla bla bla...")
					os.Exit(1)
				}
				words[idx-nbr+l] = strings.ToUpper(words[idx-nbr+l])
				words[idx] = ""
				words[idx+1] = ""
				// fmt.Println(words[idx-nbr+l])
			}

		}
	}
	return fixTheVowel(strings.Join(words, " "))
}
