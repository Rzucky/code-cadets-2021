package sqlite

import (
	controllermodel "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapStorageBetToBetDto(storageBet storagemodels.Bet) controllermodel.BetDto
}
