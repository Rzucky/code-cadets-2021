package progressivetax_test

import (
	"testing"

	"code-cadets-2021/homework1/task2/progressivetax"
)

func TestCalculateProgressiveTax(t *testing.T) {
	for _, tc := range getTestCases() {
		actualOutput, actualErr := progressivetax.CalculateProgressiveTax(tc.income, tc.taxBrackets)

		if tc.expectingError {
			if actualErr == nil {
				t.Errorf("Expected an error but got `nil` error")
			}
		} else {
			if actualErr != nil {
				t.Errorf("Expected no error but got non-nil error: %v", actualErr)
			}

			if actualOutput != tc.expectedOutput{
				t.Errorf(
					"Actual and expected output is not equal - actual: %v, expected: %v",
					actualOutput,
					tc.expectedOutput)
			}
		}
	}
}
