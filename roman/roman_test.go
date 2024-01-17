package roman

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 converts to I", 1, "I"},
		{"2 converts to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to III", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"11 gets converted to XI", 11, "XI"},
		{"12 gets converted to XII", 12, "XII"},
		{"13 gets converted to XIII", 13, "XIII"},
		{"14 gets converted to XIII", 14, "XIV"},
		{"15 gets converted to XIII", 15, "XV"},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			assertString(got, test.Want, t)
		})
	}
}

func assertString(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
