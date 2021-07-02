package services

import "github.com/superbet-group/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/infrastructure/rabbitmq/models"

// BetPublisher handles bets received queue publishing.
type BetPublisher interface {
	Publish(dto models.BetDto) error
}
