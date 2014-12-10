package generator

import "testing"

func TestRoman(t *testing.T) {
	for _, tt := range []struct {
		in  int
		out string
	}{
		{-1, ""},
		{0, ""},
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{10, "X"},
		{15, "XV"},
		{17, "XVII"},
		{20, "XX"},
		{42, "XLII"},
		{100, "C"},
		{102, "CII"},
		{412, "CDXII"},
		{688, "DCLXXXVIII"},
	} {
		got, err := roman(tt.in)
		if err != nil && err.Error() != "not a valid number" {
			t.Fatalf("unexpected error: %s", err.Error())
		}

		if got != tt.out {
			t.Fatalf("roman(%d) = %s, want %s", tt.in, got, tt.out)
		}
	}
}
