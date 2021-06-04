package controllers

import "github.com/superbet-group/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"

// BetValidator validates bet requests.
type BetValidator interface {
	BetInputIsValid(eventUpdateRequestDto models.BetRequestDto) bool
}
