package reloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EditText(text string) string {
	words := strings.Fields(text)
	for idx, word := range words {
		if words[len(words)-1] == "(low," || words[len(words)-1] == "(up," || words[len(words)-1] == "(cap," {
			fmt.Println("Err: the number next to the function should at least equal the range of str ")
			os.Exit(1)
		}
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
			if idx == 0 {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			decimal, err := strconv.ParseInt(words[idx-1], 2, 64)
			if err != nil {
				fmt.Println("invalid bin input")
				os.Exit(1)
			}
			words[idx-1] = strconv.FormatInt(decimal, 10)
			words[idx] = ""
		} else if word == "(hex)" {
			if idx == 0 {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}

			decimal, err := strconv.ParseInt(words[idx-1], 16, 64)
			if err != nil {
				fmt.Println("invalid hex input")
				os.Exit(1)
			}
			words[idx-1] = strconv.FormatInt(decimal, 10)
			words[idx] = ""
		
		} else if word == "(cap)" {
			if idx == 0 {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			words[idx-1] = toCapitalize(words[idx-1])
			words[idx] = ""
		} else if word == "(low," {
			if idx < len(words)-1 {
				if words[idx+1][len(words[idx+1])-1] != ')' {
					return text
				}
				nbrStr := (strings.TrimRight(words[idx+1], ")"))
				nbr, _ := strconv.Atoi(nbrStr)
				for l := 0; l < nbr; l++ {
					if nbr > idx {
						fmt.Println("Err: the number next to the function should at least equal to the range of str ")
						return (text)
					}
					words[idx-nbr+l] = strings.ToLower(words[idx-nbr+l])
					words[idx] = ""
					words[idx+1] = ""
					// fmt.Println(words[idx-nbr+l])
				}
			} else {
				fmt.Println("Err: please enter a correct format")
				return (text)
			}
		} else if word == "(cap," {
			if idx < len(words)-1 {
				if words[idx+1][len(words[idx+1])-1] != ')' {
					return text
				}
				nbrStr := (strings.TrimRight(words[idx+1], ")"))
				nbr, _ := strconv.Atoi(nbrStr)
				for l := 0; l < nbr; l++ {
					if nbr > idx {
						fmt.Println("Err: the number next the function should at least equal to the range of str")
						return (text)
					}
					words[idx-nbr+l] = toCapitalize(words[idx-nbr+l])
					words[idx] = ""
					words[idx+1] = ""
					// fmt.Println(words[idx-nbr+l])
				}
			} else {
				fmt.Println("Err: please enter a correct format")
				return text
			}
		} else if word == "(up," {
			if idx < len(words)-1 {
				if words[idx+1][len(words[idx+1])-1] != ')' {
					return text
				}
				nbrStr := (strings.TrimRight(words[idx+1], ")!:;,"))
				nbr, _ := strconv.Atoi(nbrStr)
				for l := 0; l < nbr; l++ {
					if nbr > idx {
						fmt.Println("Err: the number next the function should at least equal to the range of str")
						return (text)
					}
					words[idx-nbr+l] = strings.ToUpper(words[idx-nbr+l])
					words[idx] = ""
					words[idx+1] = ""
					// fmt.Println(words[idx-nbr+l])
				}
			} else {
				fmt.Println("Err: please enter a correct format")
				return (text)
			}
		}
	}
	return FixSingleQuotes(Punctuation(fixTheVowel(strings.Join(words, " "))))
}
