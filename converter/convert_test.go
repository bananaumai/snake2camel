package converter

import "testing"

func TestConvert(t *testing.T) {
	testCases := []struct {
		test 		string
		expect 		string
		lowerExpect string
	} {
		{ "snake_case", "SnakeCase", "snakeCase"},
		{"long_snake_case_pattern", "LongSnakeCasePattern", "longSnakeCasePattern"},
	}

	for _, testCase := range testCases {
		res := Convert(testCase.test, false)
		if res != testCase.expect {
			t.Fatalf("Failed convert to CamelCase. Expect %s but got %s\n", testCase.expect, res)
		}

		res = Convert(testCase.test, true)
		if res != testCase.lowerExpect {
			t.Fatalf("Failed convert to lowerCamelCase. Expect %s but got %s\n", testCase.lowerExpect, res	)
		}
	}
}
