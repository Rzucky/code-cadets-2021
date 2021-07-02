package services

import (
	controllermodels "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure/sqlite/models"
)

// BetMapper maps bet to output format.
type BetMapper interface {
	MapStorageBetToBetDto(storageBet storagemodels.Bet) controllermodels.BetDto
}
