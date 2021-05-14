package fizzbuzz_test

type testCase struct {
	inputStart   int
	inputEnd     int

	expectedOutput []string
	expectingError bool
}


func getTestCases() []testCase {
	return []testCase{
		{
			inputStart:   1,
			inputEnd:     20,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
			expectingError: false,
		},
		{
			inputStart:   5,
			inputEnd:     8,

			expectedOutput: []string{"Buzz", "Fizz", "7", "8"},
			expectingError: false,
		},
		{
			inputStart:   10,
			inputEnd:     46,

			expectedOutput: []string{"Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz", "26", "Fizz", "28", "29", "FizzBuzz", "31", "32", "Fizz", "34", "Buzz", "Fizz", "37", "38", "Fizz", "Buzz", "41", "Fizz", "43", "44", "FizzBuzz", "46"},
			expectingError: false,
		},
		{
			inputStart:   10,
			inputEnd:     10,
			expectedOutput: []string{"Buzz"},
			expectingError: false,
		},
		{
			inputStart:   0,
			inputEnd:     12,

			expectingError: true,
		},
		{
			inputStart:   -1,
			inputEnd:     0,

			expectingError: true,
		},
		{
			inputStart:   1,
			inputEnd:     0,

			expectingError: true,
		},
		{
			inputStart:   -1,
			inputEnd:     -10,

			expectingError: true,
		},
		{
			inputStart:   10,
			inputEnd:     0,

			expectingError: true,
		},

	}
}
