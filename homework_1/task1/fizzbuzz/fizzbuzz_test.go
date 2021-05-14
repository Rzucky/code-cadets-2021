package fizzbuzz_test

import (
	"testing"

	"code-cadets-2021/homework1/task1/fizzbuzz"
)

func areSlicesEqual(first []string, second []string) bool {
	if len(first) != len(second) {
		return false
	}

	for idx, x := range first {
		if x != second[idx] {
			return false
		}
	}

	return true
}


func TestFizzBuzz(t *testing.T){
	for _, tc := range getTestCases() {
		actualOutput, actualErr := fizzbuzz.FizzBuzz(tc.inputStart, tc.inputEnd)

		if tc.expectingError {
			if actualErr == nil {
				t.Errorf("Expected an error but got `nil` error")
			}
		} else {
			if actualErr != nil {
				t.Errorf("Expected no error but got non-nil error: %v", actualErr)
			}

			if !areSlicesEqual(actualOutput, tc.expectedOutput) {
				t.Errorf(
					"Actual and expected output is not equal - actual: %v, expected: %v",
					actualOutput,
					tc.expectedOutput)
			}
		}
	}
}
