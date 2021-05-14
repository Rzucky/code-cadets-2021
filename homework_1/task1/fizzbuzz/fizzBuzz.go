package fizzbuzz

import (
	"strconv"

	"github.com/pkg/errors"
)

func FizzBuzz(start, end int) ([]string, error){
	if start > end {
		return nil, errors.New("range start is greater than range end")
	}

	if start <= 0 {
		return nil, errors.New("start is 0 or negative")
	}

	if end <= 0 {
		return nil, errors.New("end is 0 or negative")
	}

	var changedOutput []string

	for i := start; i <= end; i++ {
		switch{
		//using mod 15 can save one check instead of (mod 3 && mod 5)
		case i % 15 == 0:
			changedOutput = append(changedOutput, "FizzBuzz")
		case i % 3 == 0:
			changedOutput = append(changedOutput, "Fizz")
		case i % 5 == 0:
			changedOutput = append(changedOutput, "Buzz")
		default:
			//strconv.Itoa(i) converts int to a decimal string
			changedOutput = append(changedOutput, strconv.Itoa(i))
		}
	}

	return changedOutput, nil
}
