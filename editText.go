package reloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EditText(text string) string {
	words := strings.Fields(text)
	if words[len(words)-1] == "(low," || words[len(words)-1] == "(up," || words[len(words)-1] == "(cap," {
		fmt.Println("Err: the number next to the function should at least equal the range of str ")
		os.Exit(1)
	}

	findekey := false
	for !findekey {
		for idx, word := range words {
			if (word == "(up)" || word == "(cap)" || word == "(hex)" || word == "(low)" || word == "(bin)") && idx == 0 {
				fmt.Println("err: please enter a text first")
				os.Exit(1)
			} else if word == "(up," || word == "(cap," || word == "(low," {
				if idx < len(words)-1 && words[idx+1][len(words[idx+1])-1] == ')' {
					if word == "(up)" {
						words[idx-1] = strings.ToUpper(words[idx-1])
						words[idx] = ""
						break
					} else if word == "(low)" {
						words[idx-1] = strings.ToLower(words[idx-1])
						words[idx] = ""
						break
					} else if word == "(bin)" {
						decimal, err := strconv.ParseInt(words[idx-1], 2, 64)
						if err != nil {
							fmt.Println("invalid bin input")
							os.Exit(1)
						}
						words[idx-1] = strconv.FormatInt(decimal, 10)
						words[idx] = ""
						break
					} else if word == "(hex)" {
						decimal, err := strconv.ParseInt(words[idx-1], 16, 64)
						if err != nil {
							fmt.Println("invalid hex input")
							os.Exit(1)
						}
						words[idx-1] = strconv.FormatInt(decimal, 10)
						words[idx] = ""
						break
					} else if word == "(cap)" {
						words[idx-1] = toCapitalize(words[idx-1])
						words[idx] = ""
						break
					} else if word == "(low," {
						nbrStr := strings.TrimRight(words[idx+1], ")")
						nbr, _ := strconv.Atoi(nbrStr)
						for l := 0; l < nbr; l++ {
							if nbr > idx {
								fmt.Println("Err: the number next to the function should at least equal to the range of str ")
								return text
							}
							words[idx-nbr+l] = strings.ToLower(words[idx-nbr+l])
						}
						words[idx] = ""
						words[idx+1] = ""
						break
					} else if word == "(cap," {
						nbrStr := strings.TrimRight(words[idx+1], ")")
						nbr, _ := strconv.Atoi(nbrStr)
						for l := 0; l < nbr; l++ {
							if nbr > idx {
								fmt.Println("Err: the number next the function should at least equal to the range of str")
								return text
							}
							words[idx-nbr+l] = toCapitalize(words[idx-nbr+l])
						}
						words[idx] = ""
						words[idx+1] = ""
						break
					} else if word == "(up," {
						nbrStr := strings.TrimRight(words[idx+1], ")")
						nbr, _ := strconv.Atoi(nbrStr)
						for l := 0; l < nbr; l++ {
							if nbr > idx {
								fmt.Println("Err: the number next the function should at least equal to the range of str")
								return text
							}
							words[idx-nbr+l] = strings.ToUpper(words[idx-nbr+l])
						}
						words[idx] = ""
						words[idx+1] = ""
						break
					} else {
						fmt.Println("Err: please enter a correct format")
						return text
					}
				} else {
					fmt.Println("Err: please enter a correct format")
					return text
				}
			}
			if idx == len(words)-1 {
				findekey = true
			}
		}
		words = strings.Fields(strings.Join(words, " "))
	}
	return FixSingleQuotes(Punctuation(fixTheVowel(strings.Join(words, " "))))
}
