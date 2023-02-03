package main

import "testing"

type password_validationTest struct {
	init_password string
	expected      int
}

var password_validationTests = []password_validationTest{
	// check 4 case (<6, >=20, no uppercase, lowercase or digit, have 3 consecutive)
	{"aA1", 3},                    // example shorter than 6
	{"1445D1cd", 0},               // example
	{"aA1aA1aA1aA1aA1aA1aA1", 2},  // password longer than 19
	{"a1a1a1", 1},                 // no uppercase letter
	{"aaaA11", 1},                 // have 3 consecutive char
	{"aa.AA", 1},                  // shorter than 6 and no digit
	{"1aaaA", 1},                  // shorter than 6 and have 3 consecutive char
	{"11aa11aa..11aa!!11aa..", 4}, // password longer than 19 and no uppercase letter
	{".....aA1aA1aA1aA1aA1", 2},   // password longer than 19 and have 3 consecutive char
	{"aaaA", 2},                   // shorter than 6, no digit and have 3 consecutive char
	{"...a...A...a...A...a", 5},   // password longer than 19, no digit and have 3 consecutive char
}

func TestPassword_validation(t *testing.T) {

	for _, test := range password_validationTests {
		got := password_validation(test.init_password)
		if got != test.expected {
			t.Errorf("%v : got %v, wanted %v\n", test.init_password, got, test.expected)
		}
	}
}
