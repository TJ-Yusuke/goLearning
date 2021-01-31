package roman_numerals

import "strings"

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic > 4:
			result.WriteString("Ⅴ")
			arabic -= 5
		case arabic > 3:
			result.WriteString("ⅠⅤ")
			arabic -= 4
		default:
			result.WriteString("Ⅰ")
			arabic--
		}
	}

	return result.String()
}
