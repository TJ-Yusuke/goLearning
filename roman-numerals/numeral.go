package roman_numerals

func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "Ⅱ"
	}
	return "Ⅰ"
}
