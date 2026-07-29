package luhn

import "testing"

func TestIsValid(t *testing.T) {
	cases := map[string]bool{
		"79927398713":      true,  // classic Luhn example
		"79927398710":      false, // wrong check digit
		"4539 5787 6362 1486": true, // separators are ignored
		"":                 true,  // empty sums to zero
	}
	for in, want := range cases {
		if got := IsValid(in); got != want {
			t.Errorf("IsValid(%q) = %v, want %v", in, got, want)
		}
	}
}
