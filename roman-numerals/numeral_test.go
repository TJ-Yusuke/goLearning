package roman_numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to Ⅰ", 1, "I"},
		{"2 gets converted to ⅠⅠ", 2, "II"},
		{"3 gets converted to ⅠⅠⅠ", 3, "III"},
		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to Ⅹ", 10, "Ⅹ"},
		{"14 gets converted to ⅩIV", 14, "ⅩIV"},
		{"18 gets converted to ⅩVIII", 18, "ⅩVIII"},
		{"20 gets converted to ⅩⅩ", 20, "ⅩⅩ"},
		{"39 gets converted to ⅩⅩⅩIX", 39, "ⅩⅩⅩIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
