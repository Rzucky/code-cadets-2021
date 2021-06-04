package services

import (
	"github.com/google/uuid"
	"github.com/superbet-group/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"
	rabbitmodel "github.com/superbet-group/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/infrastructure/rabbitmq/models"
)

// BetService implements bet related functions.
type BetService struct {
	betPublisher BetPublisher
}

// NewBetService creates a new instance of BetService.
func NewBetService(betePublisher BetPublisher) *BetService {
	return &BetService{
		betPublisher: betePublisher,
	}
}

// SendBetWithId sends bet received message to the queues with id.
func (e BetService) SendBetWithId(requestDto models.BetRequestDto) error {
	newId := (uuid.New()).String()

	outModelDto := rabbitmodel.BetDto{
		Id:                   newId,
		CustomerId:           requestDto.CustomerId,
		SelectionId:          requestDto.SelectionId,
		SelectionCoefficient: requestDto.SelectionCoefficient,
		Payment:              requestDto.Payment,
	}

	return e.betPublisher.Publish(outModelDto)
}
