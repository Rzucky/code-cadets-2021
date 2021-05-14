package main

import (
	"fmt"
	"log"

	"code-cadets-2021/homework1/task2/progressivetax"
	"github.com/pkg/errors"
)

func main() {

	fmt.Println("Enter amount earned: ")

	var enteredIncome float32
	_, err := fmt.Scanf("%f", &enteredIncome)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "while entering income"))
	}

	var taxBracket = []progressivetax.Bracket{
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
	}

	calculatedTax, err := progressivetax.CalculateProgressiveTax(enteredIncome, taxBracket)
	if err != nil {
		log.Fatal(
			err,
		)
	}
	fmt.Printf("Total tax for the amount of %.1f is: %.2f", enteredIncome, calculatedTax)
}
