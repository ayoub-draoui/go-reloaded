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
		lastItem := words[len(words)-1]
		if words[len(words)-1] == "(low," || words[len(words)-1] == "(up," || words[len(words)-1] == "(cap," {
			fmt.Println("Err: the number next the function should at least equal the range of str ")
			os.Exit(1)
		}
		// if lastItem[len(lastItem)-1] == ')' && idx < len(words)-1 {
		// 	// fmt.Println("!! last " + string(lastItem[len(lastItem)-1]))
		// 	fmt.Println("Err: not correct !!!!!! ")
		// 	continue
		// 	// os.Exit(1)
		// }
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
			if idx == 0 {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			}
			words[idx-1] = toCapitalize(words[idx-1])
			words[idx] = ""
		} else if word == "(low," {
			if lastItem[len(lastItem)-1] == ')' && idx < len(words)-1 {
				nbrStr := (strings.TrimRight(words[idx+1], ")"))
				nbr, _ := strconv.Atoi(nbrStr)
				for l := 0; l < nbr; l++ {
					if nbr > idx {
						fmt.Println("Err: the number next the function should at least equal the range of str ")
						return (text)
					}
					words[idx-nbr+l] = strings.ToLower(words[idx-nbr+l])
					words[idx] = ""
					words[idx+1] = ""
					// fmt.Println(words[idx-nbr+l])
				}

			}else{
				fmt.Println("Err: not correct !!!!!! ")
				return (text)
			}
		} else if word == "(cap," {
			if lastItem[len(lastItem)-1] == ')' && idx < len(words)-1 {
				nbrStr := (strings.TrimRight(words[idx+1], ")"))
				nbr, _ := strconv.Atoi(nbrStr)
				for l := 0; l < nbr; l++ {
					if nbr > idx {
						fmt.Println("Err: the number next the function should at least equal the range of str")
						return (text)
					}
					words[idx-nbr+l] = toCapitalize(words[idx-nbr+l])
					words[idx] = ""
					words[idx+1] = ""
					// fmt.Println(words[idx-nbr+l])
				}

			} else {
				fmt.Println("Err: not correct !!!!!! ")
				os.Exit(1)
			}
		} else if word == "(up," {
			if lastItem[len(lastItem)-1] == ')' && idx < len(words)-1 {
				nbrStr := (strings.TrimRight(words[idx+1], ")!:;,"))
				nbr, _ := strconv.Atoi(nbrStr)
				for l := 0; l < nbr; l++ {
					if nbr > idx {
						fmt.Println("Err: bla bla bla...")
						return (text)
					}
					words[idx-nbr+l] = strings.ToUpper(words[idx-nbr+l])
					words[idx] = ""
					words[idx+1] = ""
					// fmt.Println(words[idx-nbr+l])
				}
			} else {
				fmt.Println("Err: not correct !!!!!! ")
				return (text)
			}

			//}
		}
	}
	// for _, word := range words {
	// 	fmt.Println(word)
	// }
	// fmt.Println(words)
	// return ""

	return Punctuation(fixTheVowel(strings.Join(words, " ")))
}
