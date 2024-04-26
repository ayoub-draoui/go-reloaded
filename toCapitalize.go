package reloaded



func toCapitalize(str string) string {
	runes := []rune(str)
	for i, v := range runes {
		if i == 0 && v >= 'a' && v <= 'z' {
			runes[i] = v - 32
		} else if i > 0 && v >= 'A' && v <= 'Z' {
			runes[i] = v + 32
		}
	}
	return string(runes)
}