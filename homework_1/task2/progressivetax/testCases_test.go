package progressivetax_test

import "code-cadets-2021/homework1/task2/progressivetax"

type testCase struct {
	income float32

	taxBrackets    []progressivetax.Bracket
	expectedOutput float32
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			income: 7000,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: 1000,
				},
				{
					TaxRate:   20,
					Threshold: 5000,
				},
				{
					TaxRate:   30,
					Threshold: 10000,
				},
			},
			expectedOutput: 800,
			expectingError: false,
		},

		{
			income: 3650,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: 1000,
				},
				{
					TaxRate:   20,
					Threshold: 5000,
				},
				{
					TaxRate:   30,
					Threshold: 10000,
				},
			},
			expectedOutput: 265,
			expectingError: false,
		},
		{
			income: 150,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: 1000,
				},
				{
					TaxRate:   20,
					Threshold: 5000,
				},
				{
					TaxRate:   30,
					Threshold: 10000,
				},
			},
			expectedOutput: 0,
			expectingError: false,
		},

		{
			income: -150,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: 1000,
				},
				{
					TaxRate:   20,
					Threshold: 5000,
				},
				{
					TaxRate:   30,
					Threshold: 10000,
				},
			},

			expectingError: true,
		},
		{
			income: 7000,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   10,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: 1000,
				},
				{
					TaxRate:   20,
					Threshold: 5000,
				},
				{
					TaxRate:   30,
					Threshold: 10000,
				},
			},
			expectedOutput: 900,
			expectingError: false,
		},
		{
			income: 7000,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   -10,
					Threshold: 1000,
				},
				{
					TaxRate:   20,
					Threshold: 5000,
				},
				{
					TaxRate:   30,
					Threshold: 10000,
				},
			},

			expectingError: true,
		},
		{
			income: 7000,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: -1000,
				},
				{
					TaxRate:   20,
					Threshold: 5000,
				},
				{
					TaxRate:   30,
					Threshold: 10000,
				},
			},

			expectingError: true,
		},
		{
			income: 7000,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
			},

			expectingError: false,
		},
		{
			income: 7000,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: 3000,
				},
				{
					TaxRate:   20,
					Threshold: 7000,
				},
				{
					TaxRate:   60,
					Threshold: 20000,
				},
			},
			expectedOutput: 400,
			expectingError: false,
		},
		{
			income: 27000,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: 3000,
				},
				{
					TaxRate:   20,
					Threshold: 7000,
				},
				{
					TaxRate:   60,
					Threshold: 20000,
				},
			},
			expectedOutput: 7200,
			expectingError: false,
		},
		{
			income: 0,

			taxBrackets: []progressivetax.Bracket{
				{
					TaxRate:   0,
					Threshold: 0,
				},
				{
					TaxRate:   10,
					Threshold: 3000,
				},
				{
					TaxRate:   20,
					Threshold: 7000,
				},
				{
					TaxRate:   60,
					Threshold: 20000,
				},
			},
			expectedOutput: 0,
			expectingError: false,
		},
	}
}
