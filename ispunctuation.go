package reloaded

import (
	 
	"strings"
)


func Punctuation(text string) string {
	result := ""
	if len(text) < 2 {
		return text
	}
	 for idx, char:= range text {
		is_punc := strings.Contains(".,;:!?", string(char))
		if is_punc {
			if idx == 0 {
				result += string(char)
				if text[idx+1] != ' ' {
					result += " "
				}
				// if the special character is located at the last of str
			}else if idx == len(text)-1{
				if result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result += string(char)
			}else{
				if result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result += string(char)
				if text[idx+1] != ' ' {
					result += " "
				}
			}
			} else {
			result += string(char)

		}
	 }

return result
}
