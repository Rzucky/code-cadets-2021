package mappers

import (
	controllermodel "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
	storagemodel "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure/sqlite/models"
)

// BetMapper maps storage bets to domain bets and vice versa.
type BetMapper struct {
}

// NewBetStorageMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapStorageBetToBetDto maps the given storage bet into bet DTO. Floating point values will
// be converted from corresponding integer values of the storage bet by dividing them with 100.
// Removes data which should not be given over API calls.
func (m *BetMapper) MapStorageBetToBetDto(storageBet storagemodel.Bet) controllermodel.BetDto {
	return controllermodel.BetDto{
		Id:                   storageBet.Id,
		Status:               storageBet.Status,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
		Payout:               float64(storageBet.Payout) / 100,
	}
}
