package roman_numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "Ⅰ"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
