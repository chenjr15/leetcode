package vaildparentheses

import "testing"

func TestIsVaild(t *testing.T) {
	tests := []struct {
		input   string
		correct bool
	}{
		{"{", false},
		{"{}", true},
		{"({})", true},
		{"({)}", false},
		{"{({[]})}", true},
		{"[[[]]]", true},
	}

	for _, tcase := range tests {

		r := isValid(tcase.input)
		if r == tcase.correct {
			t.Logf("Testcase %v  passed. ", tcase.input)
		} else {
			t.Errorf("Faillde %s , go %v should be %v", tcase.input, r, tcase.correct)
		}
	}
}
