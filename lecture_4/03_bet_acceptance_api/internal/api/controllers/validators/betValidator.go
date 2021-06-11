package validators

import (
	"github.com/superbet-group/code-cadets-2021/lecture_4/03_bet_acceptance_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"
)

// BetValidator validates bet input requests.
type BetValidator struct {
	Configuration config.Config
}

// NewBetValidator creates a new instance of BetValidator.
func NewBetValidator(configuration config.Config) *BetValidator {
	return &BetValidator{
		Configuration: configuration,
	}
}

// BetInputIsValid checks if bet input is valid.
// All data should be populated and non-default
// Some data should be according to config file
func (e *BetValidator) BetInputIsValid(betRequestDto models.BetRequestDto) bool {
	if betRequestDto.CustomerId != "" &&
		betRequestDto.SelectionId != "" &&
		betRequestDto.SelectionCoefficient <= e.Configuration.InputConfig.MaxCoefficient &&
		betRequestDto.Payment <= e.Configuration.InputConfig.MaxPayment &&
		betRequestDto.Payment >= e.Configuration.InputConfig.MinPayment {
		return true
	}
	return false
}
