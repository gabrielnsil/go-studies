package iteration

import "strings"

func Repeat(character string, repeatCount int) string {

	var repeated strings.Builder
	
	if repeatCount < 1 {
		return character
	}

	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}

