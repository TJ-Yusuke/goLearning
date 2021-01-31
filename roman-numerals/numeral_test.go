package roman_numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to Ⅰ", 1, "Ⅰ"},
		{"2 gets converted to ⅠⅠ", 2, "ⅠⅠ"},
		{"3 gets converted to ⅠⅠⅠ", 3, "ⅠⅠⅠ"},
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
