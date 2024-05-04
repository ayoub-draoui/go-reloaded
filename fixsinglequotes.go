package reloaded

import "strings"

func FixSingleQuotes(text string) string {
	result := ""
	first := true
	rm_next := false
	if len(text) <= 3 {
		return text
	}
	for idx, char := range text {
		if char == '\'' {
			// if single cote
			if idx == 0 {
				// if is first char
				result += string(char)
				first = false
				if text[idx+1] == ' ' {
					rm_next = true
				}
			} else if idx == len(text)-1 {
				// if is last char
				if text[idx-1] == ' ' {
					result = result[:len(result)-1]
				}
				result += "'"
			} else {
				// your logical code
				if text[idx-1] != ' ' && text[idx+1] != ' ' {
					result += "'"
					continue
				}
				if first {
					// first cote
					if result[len(result)-1] != ' ' {
						result += " "
					}
					result += "'"
					if text[idx+1] == ' ' {
						rm_next = true
					}
					first = false
				} else {
					// last cote
					first = true
					if result[len(result)-1] == ' ' {
						result = result[:len(result)-1]
					}
					is_punc := strings.Contains(",;:.!? ", string(text[idx+1]))
					result += "'"
					if !is_punc {
						result += " "
					}
				}
			}
		} else {
			// if not single cote
			if rm_next {
				rm_next = false
				continue
			}
			result += string(char)
		}
	}
	return result
}
