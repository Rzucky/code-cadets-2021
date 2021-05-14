package progressivetax

import (
	"github.com/pkg/errors"
)

type Bracket struct {
	TaxRate   float32
	Threshold float32
}

func ValidateIncomeAndTaxBrackets(income float32, taxBrackets []Bracket) error {
	if income < 0 {
		return errors.New("income is lower than 0.0")
	}
	if taxBrackets[0].Threshold != 0.0 {
		return errors.New("threshold of the first bracket is not 0.0")
	}

	if len(taxBrackets) < 2 {
		return errors.New("there is less than 2 tax brackets")
	}

	for _, taxBracket := range taxBrackets {
		if taxBracket.TaxRate < 0.0 {
			return errors.New("negative tax rate in one or more brackets")
		}
		if taxBracket.Threshold < 0.0 {
			return errors.New("negative threshold in one or more brackets")
		}
	}
	return nil
}

func CalculateProgressiveTax(income float32, taxBrackets []Bracket) (float32, error) {
	err := ValidateIncomeAndTaxBrackets(income, taxBrackets)
	if err != nil {
		return 0, err
	}
	var calculatedTax float32

	for i := 1; i < len(taxBrackets); i++ {
		//checking if we are over one of the thresholds
		if income > taxBrackets[i].Threshold {
			//taking tax percentage into 0.X form from the last bracket
			taxPercentage := taxBrackets[i-1].TaxRate / 100

			//difference between the threshold to multiply with the taxPercentage
			calculatedTax += (taxBrackets[i].Threshold - taxBrackets[i-1].Threshold) * taxPercentage

			//once we have taken care of every bracket, we need to calculate tax on the last one
			//last one doesn't have a top bracket over it so we take the difference from income
			if i == (len(taxBrackets) - 1) {
				//now we take info from the i-th bracket
				taxPercentage = taxBrackets[i].TaxRate / 100

				calculatedTax += (income - taxBrackets[i].Threshold) * taxPercentage

			}

			//similar to the top one, this checks the tax when we are over a threshold but not the max threshold
		} else if income > taxBrackets[i-1].Threshold {
			taxPercentage := taxBrackets[i-1].TaxRate / 100

			calculatedTax += (income - taxBrackets[i-1].Threshold) * taxPercentage
		}

	}

	return calculatedTax, nil
}
