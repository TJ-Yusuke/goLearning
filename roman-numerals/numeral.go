package roman_numerals

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("Ⅰ")
	}

	return result.String()
}